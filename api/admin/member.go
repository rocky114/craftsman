package admin

import (
	"craftsman/response"
	"craftsman/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	items := service.GetMembers()

	result := response.Gin{
		C: c,
	}

	result.Response(http.StatusOK, response.Success, items)
}

func Create(c *gin.Context) {

}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {

}
