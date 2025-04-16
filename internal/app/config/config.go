package config

import (
	"time"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	JWT      JWTConfig
	Logging  LoggingConfig
	Web      WebConfig
}

type AppConfig struct {
	Name  string `yaml:"name"`
	Env   string `yaml:"env"` // development/staging/production
	Port  string `yaml:"port"`
	Debug bool   `yaml:"debug"`
}

type DatabaseConfig struct {
	URL          string `yaml:"url"`
	MaxOpenConns int    `yaml:"max_open_conns"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
}

type JWTConfig struct {
	Secret     string        `yaml:"secret"`
	Expiration time.Duration `yaml:"expiration"`
}

type LoggingConfig struct {
	Level  string `yaml:"level"`  // debug/info/warn/error
	Format string `yaml:"format"` // json/text
}

type WebConfig struct {
	StaticDir   string `yaml:"static_dir"`
	TemplateDir string `yaml:"template_dir"`
}
