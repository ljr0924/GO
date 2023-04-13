package viper_demo

import (
	"fmt"
	"testing"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Env  string `mapstructure:"Env" yaml:"Env"`
	Port int64  `mapstructure:"Port" yaml:"Port"`
	Root string `mapstructure:"Root" yaml:"Root"`
}
type ConfigV1 struct {
	Env  string `yaml:"Env"`
	Port int64  `yaml:"Port"`
	Root string `yaml:"Root"`
}
type ConfigV2 struct {
	Env  string `mapstructure:"Env"`
	Port int64  `mapstructure:"Port"`
	Root string `mapstructure:"Root"`
}
type ConfigV3 struct {
	Env  string
	Port int64
	Root string
}

var v = viper.New()

var c = &ConfigV3{}

func init() {
	fmt.Println("init")
	v.SetConfigFile("config.yaml")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic("error")
	}
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		if err = v.Unmarshal(c); err != nil {
			fmt.Println("配置变更")
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(c); err != nil {
		fmt.Println(err)
	}
}

func TestOnChange(t *testing.T) {
	fmt.Println("run test")
	for true {
		fmt.Println(c)
		time.Sleep(time.Second * 2)
	}
}
