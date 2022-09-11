package admin

import (
	"github.com/gin-gonic/gin"
)

func InitRoute(r *gin.Engine) {
	r.GET("/captcha", GetCaptcha)
}
