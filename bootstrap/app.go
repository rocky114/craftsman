package bootstrap

import (
	"craftsman/config"
	"fmt"
	"github.com/spf13/viper"
)

var (
	GlobalConfig config.Server
)

func Viper() {
	viper.SetConfigFile("config.toml")

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if err := viper.Unmarshal(&GlobalConfig); err != nil {
		panic(fmt.Errorf("config parse error: %s \n", err))
	}
}
