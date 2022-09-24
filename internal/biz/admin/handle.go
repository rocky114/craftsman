package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rocky114/craftsman/internal/response"
	"github.com/rocky114/craftsman/internal/service/user"
	"github.com/rocky114/craftsman/internal/storage"
	"github.com/rocky114/craftsman/pkg/captcha"
	"github.com/sirupsen/logrus"
)

type captchaResponse struct {
	CaptchaId  string `json:"captcha_id"`
	CaptchaVal string `json:"captcha_val"`
}

func GetCaptcha(c *gin.Context) {
	var ret captchaResponse
	captchaId, captchaVal, err := captcha.Get()
	if err != nil {
		logrus.Errorf("captcha err: %v", err)
		c.JSON(http.StatusOK, response.Result{Code: response.Invalid, Message: "生成验证码失败", Data: ret})
		return
	}

	ret = captchaResponse{CaptchaId: captchaId, CaptchaVal: captchaVal}
	c.JSON(http.StatusOK, response.Result{Code: response.OK, Message: "", Data: ret})
}

func CreateUser(c *gin.Context) {
	var req storage.CreateUserParams
	if err := c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("createUser err: %v", err)
		c.JSON(http.StatusOK, response.Result{Code: response.Invalid, Message: response.ParameterErr})
		return
	}

	if err := user.AddUser(req); err != nil {
		logrus.Errorf("createUser err: %v", err)
		c.JSON(http.StatusOK, response.Result{Code: response.Invalid, Message: response.UnknownErr})
		return
	}

	c.JSON(http.StatusOK, response.Result{Code: response.OK, Message: ""})
}

func LoginIn(c *gin.Context) {

}
