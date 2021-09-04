package router

import (
	"craftsman/api/admin"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
	Router.Use(gin.Logger())
	Router.Use(gin.Recovery())

	Router.GET("/captcha", admin.Captcha)
	Router.POST("/auth/login", admin.Login)
	Router.GET("/members", admin.GetMembers)
	Router.POST("/members", admin.AddMember)
	Router.PUT("/members/:id", admin.UpdateMember)
	Router.DELETE("/members/:id", admin.DeleteMember)
}
