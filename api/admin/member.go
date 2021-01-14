package admin

import (
	"craftsman/model"
	"craftsman/service"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	var items []map[string]interface{}

	model.MysqlConn.Debug().Model(&model.Member{}).Find(&items)

	c.JSON(200, service.Success(items))
}
