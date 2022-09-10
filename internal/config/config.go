package config

import (
	"log"

	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var GlobalConfig server

type server struct {
	Mysql mysqlConf `json:"mysql" yaml:"mysql"`
	Log   logConf   `json:"log" yaml:"log"`
}

type mysqlConf struct {
	Host     string `json:"host" yaml:"host"`
	Port     string `json:"port" yaml:"port"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	Database string `json:"database" yaml:"database"`
	Options  string `json:"options" yaml:"options"`
}

type logConf struct {
	Path  string `json:"path" yaml:"path"`
	Level string `json:"level" yaml:"level"`
}

func init() {
	var configFilepath string
	flag.StringVarP(&configFilepath, "config", "c", "../../configs/config.yaml", "config file")
	flag.Parse()

	viper.SetConfigFile(configFilepath)
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("read config file err: %v\n", err)
	}

	if err = viper.Unmarshal(&GlobalConfig); err != nil {
		log.Fatalf("unmarshal config file err: %v\n", err)
	}
}
