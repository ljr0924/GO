package config

import (
	"context"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"path"
	"runtime"
	"sync"
)

var onceLogConfig sync.Once

type Config struct {
	Key    string
	Value  string
	Cancel context.CancelFunc
}

const (
	KConfigPaths = "configpath"
	KKafkaAddr = "kafka.addr"
)


func ReadConfig(v *viper.Viper) (*viper.Viper, bool) {
	// 设置读取配置文件
	v.SetConfigName("config")
	// 添加读取的配置文件路径
	_, fileName, _, _ := runtime.Caller(0)
	fmt.Println(fileName)
	fmt.Println(path.Dir(fileName))
	// 设置配置文件目录
	v.AddConfigPath(path.Dir(fileName))
	// 设置配置文件类型
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("err: %s\n", err.Error())
		return nil, false
	}

	return v, true
}

func WatchConfig(ctx context.Context, v *viper.Viper, key2chan map[string]chan interface{}) {

	defer func() {
		onceLogConfig.Do(func() {
			fmt.Println("watch config goroutine exists")
			if err := recover(); err != nil {
				fmt.Println("watch config goroutine panic ", err)
			}
			for _, c := range key2chan {
				close(c)
			}
		})
	}()

	// 设置监听回调参数
	v.OnConfigChange(func(event fsnotify.Event) {
		// 更新所有key
		updateConfig := func(keys ...string) {
			for _, k := range keys {
				data := v.Get(k)
				if data != nil {
					if _, ok := key2chan[k]; !ok {
						key2chan[k] = make(chan interface{})
					}
					key2chan[k] <- data
				}
			}
		}

		updateConfig(KConfigPaths, KKafkaAddr)

	})

	// 开始监听
	v.WatchConfig()

	<-ctx.Done()
}