package admin

import (
	"craftsman/response"
	"craftsman/service"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	items := service.GetMembers()
	c.JSON(200, response.Success(items))
}
