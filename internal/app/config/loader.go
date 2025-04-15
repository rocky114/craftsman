package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Load(configPath string) (*Config, error) {
	v := viper.New()

	// 设置配置文件路径
	if configPath != "" {
		v.SetConfigFile(configPath)
	} else {
		// 默认查找 configs/ 目录
		v.AddConfigPath("./configs")
		v.SetConfigName("config")
	}

	v.SetConfigType("yaml")

	// 读取环境变量覆盖配置
	// v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// 监听配置文件变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		if err := v.Unmarshal(&cfg); err != nil {
			fmt.Printf("failed to reload config: %v\n", err)
		}
	})

	return &cfg, nil
}
