package admin

import (
	"craftsman/bootstrap"
	"craftsman/model"
	"craftsman/service"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var items model.Member
	bootstrap.MysqlConn.Select("name", "nickname", "email").Find(&items)

	c.JSON(200, service.Success(items))
}
