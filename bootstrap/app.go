package bootstrap

import (
	"craftsman/config"
	"fmt"
	"github.com/spf13/viper"
)

var (
	MysqlConfig *config.Mysql
)

func Viper() {
	v := viper.New()
	v.SetConfigFile("config.yaml")

	err := v.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if err := v.Unmarshal(MysqlConfig); err != nil {
		panic(fmt.Errorf("config parse error: %s \n", err))
	}
}
