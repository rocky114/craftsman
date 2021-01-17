package router

import (
	"craftsman/api/admin"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func Bootstrap() {
	Router = gin.Default()
	Router.Use(gin.Logger())
	Router.Use(gin.Recovery())

	Router.POST("/auth/login", admin.Login)
	Router.GET("/members", admin.Index)
	Router.POST("/members", admin.Create)
	Router.PUT("/members/:id", admin.Update)
	Router.DELETE("/members/:id", admin.Delete)
}
