package admin

import (
	"craftsman/config"
	"craftsman/response"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"net/http"
)

func Captcha(c *gin.Context) {
	captchaConfig := config.GlobalConfig.Captcha
	driver := base64Captcha.NewDriverDigit(captchaConfig.Height, captchaConfig.Width, captchaConfig.Length, 0.7, 80)
	captcha := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
	_, b64s, err := captcha.Generate()

	ginContext := response.GinContext{C: c}

	if err != nil {
		ginContext.Response(http.StatusOK, response.MemberCreateFailed, nil)
		return
	}

	result := map[string]string{"captcha": b64s}

	ginContext.Response(http.StatusOK, response.Success, result)
}
