package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go_demo/log/config"
	logtail "go_demo/log/log_tail"
	"sync"
	"time"
)

var mainOnce sync.Once
var configMgr map[string]*config.Config

func ConstructMgr(configPaths interface{}, keyChan chan string) {
	configDatas := configPaths.(map[string]interface{})
	for configKey, configVal := range configDatas {
		Start(configKey, configVal.(string), keyChan)
	}
}

func Start(key, path string, keyChan chan string) {
	configData := new(config.Config)
	configData.Key = key
	configData.Value = path
	ctx, cancel := context.WithCancel(context.Background())
	configData.Cancel = cancel
	configMgr[key] = configData
	// 启动协程监听日志文件
	go logtail.WatchLogFile(key, path, ctx, keyChan)
}

func Restart(key string, keyChan chan string) {
	c, ok := configMgr[key]
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
	go logtail.WatchLogFile(key, c.Value, ctx, keyChan)
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

	keyChan := make(chan string, 4)

	configMgr = make(map[string]*config.Config)
	ConstructMgr(configPaths, keyChan)

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
			// 销毁监听文件的协程
			cancel()
			for _, oldVal := range configMgr {
				oldVal.Cancel()
			}
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
			for oldKey, oldVal := range configMgr {
				_, ok := pathDataNew[oldKey]
				if ok {
					continue
				}
				oldVal.Cancel()
				delete(configMgr, oldKey)
			}

			// 更新配置
			for configKey, configVal := range pathDataNew {
				oldVal, ok := configMgr[configKey]
				// 没找到 创建新的协程监听日志文件
				if !ok {
					Start(configKey, configVal.(string), keyChan)
					continue
				}
				// 同一个key，不同的路径，先停掉原来的，再重新开一个新的
				if oldVal.Value != configVal.(string) {
					Restart(configKey, keyChan)
					continue
				}
			}

			for mgrKey, mgrVal := range configMgr {
				fmt.Println(mgrKey, mgrVal)
			}
		case addrList, ok := <- kafkaAddrChan:
			if !ok {
				continue
			}
			for _, a := range addrList.([]interface{}) {
				fmt.Println("new addr: ", a)
			}

		case key := <- keyChan:
			fmt.Printf("restart %s after 5 second\n", key)
			time.Sleep(5 * time.Second)
			Restart(key, keyChan)
		}
	}

}
