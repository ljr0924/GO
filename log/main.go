package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go_demo/log/config"
	"go_demo/log/kafka"
	logtail "go_demo/log/log_tail"
	"sync"
	"time"
)

var mainOnce sync.Once
var configMgr *ConfigMgr

type ConfigMgr struct {
	Mgr map[string]*config.Config
	Producer *kafka.KafProducer
}

func NewConfigMgr(p *kafka.KafProducer) *ConfigMgr {
	return &ConfigMgr{
		Mgr:      make(map[string]*config.Config),
		Producer: p,
	}
}

// Construct 初始化
func (m *ConfigMgr) Construct(configPaths interface{}, keyChan chan string) {
	configDatas := configPaths.(map[string]interface{})
	for configKey, configVal := range configDatas {
		m.Start(configKey, configVal.(string), keyChan)
	}
}

// Destroy 销毁所有监听log的协程
func (m *ConfigMgr) Destroy() {
	for _, v := range m.Mgr {
		v.Cancel()
	}
}

// Start 开启监听日志文件协程
func (m *ConfigMgr) Start(key, path string, keyChan chan string) {
	configData := new(config.Config)
	configData.Key = key
	configData.Value = path
	ctx, cancel := context.WithCancel(context.Background())
	configData.Cancel = cancel
	m.Mgr[key] = configData
	// 启动协程监听日志文件
	go logtail.WatchLogFile(ctx, key, path, m.Producer, keyChan)
}

// Restart 开启监听日志文件协程
func (m *ConfigMgr) Restart(key string, keyChan chan string) {
	c, ok := m.Mgr[key]
	if !ok {
		fmt.Printf("%s not exists", key)
		return
	}
	// 先取消旧的
	c.Cancel()

	// 重启新的
	ctx, cancel := context.WithCancel(context.Background())
	c.Cancel = cancel
	// 启动协程监听日志文件
	go logtail.WatchLogFile(ctx, key, c.Value, m.Producer, keyChan)
}

func main() {

	v := viper.New()
	v, ok := config.ReadConfig(v)
	if !ok {
		fmt.Println("read config err")
		return
	}

	configPaths := v.Get("configpath")
	if configPaths == nil  {
		fmt.Println("configpath is empty")
		return
	}
	kafkaAddr := v.Get("kafka.addr")
	if kafkaAddr == nil {
		fmt.Println("kafka.addr is empty")
		return
	}
	var addrList []string
	for _, a := range kafkaAddr.([]interface{}) {
		addrList = append(addrList, a.(string))
	}

	keyChan := make(chan string, 4)

	p, err := kafka.NewKafProducer(addrList)
	if err != nil {
		fmt.Println("create kafka err")
		return
	}
	configMgr = NewConfigMgr(p)
	configMgr.Construct(configPaths, keyChan)

	ctx, cancel := context.WithCancel(context.Background())
	key2chan := make(map[string]chan interface{})
	configPathsChan := make(chan interface{})
	kafkaAddrChan := make(chan interface{})
	key2chan[config.KConfigPaths] = configPathsChan
	key2chan[config.KKafkaAddr] = kafkaAddrChan
	go config.WatchConfig(ctx, v, key2chan)

	defer func() {
		mainOnce.Do(func() {
			if err := recover(); err != nil {
				fmt.Println("main goroutine panic ", err)
			}
			// 销毁监听配置文件的协程
			cancel()
			// 销毁监听日志文件携程
			configMgr.Destroy()
			configMgr = nil
		})
	}()

	for {
		select {
		case pathData, ok := <- configPathsChan:
			if !ok {
				return
			}
			fmt.Println("main goroutine receive path")
			fmt.Println(pathData)

			pathDataNew := pathData.(map[string]interface{})

			// 先将已清除的停掉
			for oldKey, oldVal := range configMgr.Mgr {
				_, ok := pathDataNew[oldKey]
				if ok {
					continue
				}
				oldVal.Cancel()
				delete(configMgr.Mgr, oldKey)
			}

			// 更新配置
			for configKey, configVal := range pathDataNew {
				oldVal, ok := configMgr.Mgr[configKey]
				// 没找到 创建新的协程监听日志文件
				if !ok {
					configMgr.Start(configKey, configVal.(string), keyChan)
					continue
				}
				// 同一个key，不同的路径，先停掉原来的，再重新开一个新的
				if oldVal.Value != configVal.(string) {
					oldVal.Cancel()
					configMgr.Start(configKey, configVal.(string), keyChan)
					continue
				}
			}

			for mgrKey, mgrVal := range configMgr.Mgr {
				fmt.Println(mgrKey, mgrVal)
			}
		case addrList, ok := <- kafkaAddrChan:
			// 监听kafka的地址是否发生变化
			if !ok {
				continue
			}
			for _, a := range addrList.([]interface{}) {
				fmt.Println("new addr: ", a)
			}
		case key := <- keyChan:
			fmt.Printf("restart %s after 5 second\n", key)
			time.Sleep(5 * time.Second)
			configMgr.Restart(key, keyChan)
		}
	}

}
