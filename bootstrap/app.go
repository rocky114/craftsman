package bootstrap

import (
	"craftsman/config"
	"flag"
	"fmt"
	"github.com/spf13/viper"
)

var (
	GlobalConfig config.Server
)

func Config() {
	var configFile string

	flag.StringVar(&configFile, "c", "config.toml", "choose config file(shorthand)")
	flag.StringVar(&configFile, "config", "config.toml", "chose config file")

	viper.SetConfigFile(configFile)

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if err := viper.Unmarshal(&GlobalConfig); err != nil {
		panic(fmt.Errorf("config parse error: %s \n", err))
	}
}
