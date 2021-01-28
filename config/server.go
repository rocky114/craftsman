package config

type Server struct {
	Addr string `json:"addr" yaml:"addr"`
	Port string `json:"port" yaml:"port"`
}
