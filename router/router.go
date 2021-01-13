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

	Router.GET("/members", admin.Index)
}
