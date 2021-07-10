package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go_demo/log/config"
	"sync"
)

var mainOnce sync.Once
var configMgr map[string]*config.Config

func ConstructMgr(configPaths interface{}) {
	configDatas := configPaths.(map[string]interface{})
	for configKey, configVal := range configDatas {
		configData := new(config.Config)
		configData.Key = configKey
		configData.Value = configVal.(string)

		_, cancel := context.WithCancel(context.Background())
		configData.Cancel = cancel
		configMgr[configKey] = configData
	}
}

func main() {

	v := viper.New()
	configPaths, ok := config.ReadConfig(v)

	if configPaths == nil || !ok {
		fmt.Println("read config err")
		return
	}

	configMgr = make(map[string]*config.Config)
	ConstructMgr(configPaths)

	ctx, cancel := context.WithCancel(context.Background())
	pathChan := make(chan interface{})
	go config.WatchConfig(ctx, v, pathChan)

	defer func() {
		mainOnce.Do(func() {
			if err := recover(); err != nil {
				fmt.Println("main goroutine panic ", err)
			}
			cancel()
		})
	}()

	for {
		select {
		case pathData, ok := <-pathChan:
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
				if !ok {
					configData := new(config.Config)
					configData.Key = configKey
					configData.Value = configVal.(string)
					_, cancel := context.WithCancel(context.Background())
					configData.Cancel = cancel
					configMgr[configKey] = configData
					continue
				}
				if oldVal.Value != configVal.(string) {
					oldVal.Value = configVal.(string)
					oldVal.Cancel()
					_, cancel := context.WithCancel(context.Background())
					oldVal.Cancel = cancel
					continue
				}
			}

			for mgrKey, mgrVal := range configMgr {
				fmt.Println(mgrKey, mgrVal)
			}
		}
	}

}
