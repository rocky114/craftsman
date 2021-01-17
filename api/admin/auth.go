package admin

import (
	"craftsman/response"
	"craftsman/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	ginContent := response.GinContext{
		C: c,
	}

	token := service.GetToken()
	ginContent.Response(http.StatusOK, response.Success, token)
}
