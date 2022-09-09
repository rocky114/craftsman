package admin

import "github.com/rocky114/craftsman/internal/bootstrap"

func init() {
	bootstrap.Router.GET("/captcha", GetCaptcha)
}
