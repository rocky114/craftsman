package admin

import (
	"craftsman/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Captcha(c *gin.Context)  {
	ginContent := response.GinContext{C: c}
	ginContent.Response(http.StatusOK, response.Success, []string{})
}
