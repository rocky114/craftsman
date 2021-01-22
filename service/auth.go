package service

import (
	"craftsman/config"
	"craftsman/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	UserId   int    `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

type Login struct {
	Username string
	Password string
	Captcha  string
}

func Authenticate(login Login) (token string, err error) {
	var member struct {
		Id       int
		Username string
	}

	err = model.MysqlConn.
		Model(&model.Member{}).
		Where("username = ? and password = ?", login.Username, login.Password).
		First(&member).Error

	if err != nil {
		return "", err
	}

	token, err = GetToken(member.Id, member.Username)

	return
}

func GetToken(id int, username string) (token string, err error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(30 * 24 * time.Hour)

	claims := Claims{
		id,
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "rocky",
		},
	}

	secret := config.GlobalConfig.JWT.Secret

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tokenClaims.SignedString([]byte(secret))

	return
}
