package admin

import (
	"github.com/gin-gonic/gin"
)

func GetRoutes() func(r *gin.Engine) {
	return func(r *gin.Engine) {
		r.GET("/captcha", GetCaptcha)
		r.POST("/login", LoginIn)
		r.POST("/users", CreateUser)
		r.GET("/users", ListUser)
	}
}
