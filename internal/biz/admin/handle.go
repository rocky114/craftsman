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
		logrus.Errorf("GetCaptcha err: %v", err)
		c.JSON(http.StatusOK, response.NewFailed(response.ErrUnknown))
		return
	}

	ret = captchaResponse{CaptchaId: captchaId, CaptchaVal: captchaVal}
	c.JSON(http.StatusOK, response.NewSuccessed(ret))
}

func CreateUser(c *gin.Context) {
	var req storage.CreateUserParams
	if err := c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("createUser err: %v", err)
		c.JSON(http.StatusBadRequest, response.NewFailed(response.ErrInvalidParam))
		return
	}

	if err := user.CreateUser(req); err != nil {
		logrus.Errorf("createUser err: %v", err)
		c.JSON(http.StatusOK, response.NewFailed(response.ErrUnknown))
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessed([]struct{}{}))
}

func ListUser(c *gin.Context) {
	users, err := user.ListUser()
	if err != nil {
		logrus.Errorf("ListUsers err: %v", err)
		c.JSON(http.StatusOK, response.NewFailed(response.ErrUnknown))
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessed(users))
}

type loginResponse struct {
	Token string `json:"token"`
}

func LoginIn(c *gin.Context) {
	var req storage.GetUserParams

	var token string
	var err error

	if err := c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("createUser err: %v", err)
		c.JSON(http.StatusBadRequest, response.NewFailed(response.ErrInvalidParam))
		return
	}

	if token, err = user.Login(req); err != nil {
		logrus.Errorf("login err: %v", err)
		c.JSON(http.StatusOK, response.NewFailed(response.ErrUnknown))
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessed(loginResponse{Token: token}))
}
