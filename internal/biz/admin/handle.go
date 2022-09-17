package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rocky114/craftsman/internal/response"
	"github.com/rocky114/craftsman/pkg/captcha"
	"github.com/sirupsen/logrus"
)

type code struct {
	CaptchaId  string `json:"captcha_id"`
	CaptchaVal string `json:"captcha_val"`
}

func GetCaptcha(c *gin.Context) {
	var ret code
	captchaId, captchaVal, err := captcha.Get()
	if err != nil {
		logrus.Errorf("captcha err: %v", err)
		c.JSON(http.StatusOK, response.Result{Code: response.Invalid, Message: "生成验证码失败", Data: ret})
		return
	}

	ret = code{CaptchaId: captchaId, CaptchaVal: captchaVal}
	c.JSON(http.StatusOK, response.Result{Code: response.OK, Message: "", Data: ret})
}
