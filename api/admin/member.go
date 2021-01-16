package admin

import (
	"craftsman/response"
	"craftsman/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Member struct {
	Name     string `json:"name" form:"name"`
	Nickname string `json:"nickname" form:"nickname"`
	Email    string `json:"email" form:"email"`
}

func Index(c *gin.Context) {
	items := service.GetMembers()

	result := response.GinContext{
		C: c,
	}

	result.Response(http.StatusOK, response.Success, items)
}

func Create(c *gin.Context) {
	var json Member

	result := response.GinContext{
		C: c,
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		result.Response(http.StatusBadRequest, response.MemberRequestParamError, nil)
		return
	}

	result.Response(http.StatusOK, response.Success, []struct{}{})
}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {

}
