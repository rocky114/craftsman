package service

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	MemberId uint `json:"member_id"`
	jwt.StandardClaims
}

func Check() (err error) {
	return nil
}

func GetToken() string {
	nowTime := time.Now()
	expireTime := nowTime.Add(30 * 24 * time.Hour)

	claims := Claims{
		1,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "rocky",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := tokenClaims.SignedString("rocky114")

	return token
}
