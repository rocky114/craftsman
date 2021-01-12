package config

type Mysql struct {
	Host     string `json:"host" yaml:"host"`
	Port     uint   `json:"port" yaml:"port"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	Database string `json:"database" yaml:"yaml"`
}
