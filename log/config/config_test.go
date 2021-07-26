package config

import (
	"github.com/spf13/viper"
	"testing"
)

func TestReadConfig(t *testing.T) {

	v := viper.New()
	configPaths, success := ReadConfig(v)
	if !success {
		t.Fail()
		return
	}
	t.Log(configPaths)

}
