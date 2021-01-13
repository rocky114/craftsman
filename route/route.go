package route

import (
	"craftsman/api/admin"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) *gin.Engine {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/members", admin.Index)

	return router
}
