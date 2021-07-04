package config

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"path"
	"runtime"
	"sync"
)

var onceLogConfig sync.Once

type Config struct {
	Key    string
	Value  string
	cancel context.CancelFunc
}


func ReadConfig(v *viper.Viper) (interface{}, bool) {
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
	configPaths := v.Get("configpath")
	if configPaths == nil {
		return nil, false
	}

	return configPaths, true
}