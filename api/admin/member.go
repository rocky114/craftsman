package admin

import (
	"craftsman/model"
	"craftsman/response"
	"craftsman/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	items := service.GetMembers()

	ginContent := response.GinContext{
		C: c,
	}

	ginContent.Response(http.StatusOK, response.Success, items)
}

func Create(c *gin.Context) {
	var json model.Member

	ginContext := response.GinContext{
		C: c,
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		ginContext.Response(http.StatusBadRequest, response.RequestParamError, nil)
		return
	}

	result := service.CreateMember(&json)

	if result != nil {
		ginContext.Response(http.StatusOK, response.MemberCreateFailed, nil)
		return
	}

	ginContext.Response(http.StatusOK, response.Success, []struct{}{})
}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {

}
