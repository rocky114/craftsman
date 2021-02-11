package config

type Captcha struct {
	Width  int `json:"width" yaml:"width"`
	Height int `json:"height" yaml:"height"`
	Length int `json:"length" yaml:"length"`
}
