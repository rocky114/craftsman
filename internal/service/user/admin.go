package user

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/rocky114/craftsman/internal/pkg/crypt"
	"github.com/rocky114/craftsman/internal/storage"
)

var (
	tokenSignatureKey = []byte("rocky114")
)

func CreateUser(req storage.CreateUserParams) error {
	req.Password = crypt.GetMd5Str(req.Password)
	_, err := storage.GetQueries().CreateUser(context.Background(), req)
	return err
}

func ListUser() ([]storage.User, error) {
	return storage.GetQueries().ListUser(context.Background())
}

func Login(req storage.GetUserParams) (string, error) {
	req.Password = crypt.GetMd5Str(req.Password)
	user, err := storage.GetQueries().GetUser(context.Background(), req)
	if err != nil {
		return "", err
	}

	return getToken(user.ID, user.Username)
}

type customerClaim struct {
	Id       int32  `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func getToken(id int32, username string) (string, error) {
	claim := customerClaim{
		Id:       id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
			Issuer:    "rocky",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	return token.SignedString(tokenSignatureKey)
}
