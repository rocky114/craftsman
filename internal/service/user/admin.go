package user

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/rocky114/craftsman/internal/storage"
)

func AddUser(req storage.CreateUserParams) error {
	_, err := storage.GetQueries().CreateUser(context.Background(), req)
	return err
}

type customerClaim struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func Login(req storage.GetUserParams) (string, error) {
	user, err := storage.GetQueries().GetUser(context.Background(), req)
	if err != nil {
		return "", err
	}

	return getToken(user.ID, user.Username)
}

func getToken(id int32, username string) (string, error) {
	claim := customerClaim{
		Id:       1,
		Username: "",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
			Issuer:    "rocky",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claim)

	return token.SignedString("craftman")
}
