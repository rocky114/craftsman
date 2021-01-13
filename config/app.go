package config

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Application struct {
	Mysql Mysql `json:"mysql" yaml:"mysql"`
}

var GlobalConfig Application

func Bootstrap() {
	var configFile string
	flag.StringVarP(&configFile, "config", "c", "config.toml", "choose config file(shorthand)")

	flag.Parse()

	viper.SetConfigFile(configFile)

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if err := viper.Unmarshal(&GlobalConfig); err != nil {
		panic(fmt.Errorf("config parse error: %s \n", err))
	}
}
