package admin

import (
	"craftsman/response"
	"craftsman/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	ginContent := response.GinContext{C: c}

	var json service.Login

	if err := c.ShouldBindJSON(&json); err != nil {
		ginContent.Response(http.StatusBadRequest, response.RequestParamError, nil)
		return
	}

	token, err := service.Authenticate(json)

	if err != nil {
		ginContent.Response(http.StatusBadRequest, response.RequestParamError, nil)
	}

	ginContent.Response(http.StatusOK, response.Success, token)
}
