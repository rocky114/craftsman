package user

import (
	"context"

	"github.com/rocky114/craftsman/internal/storage"
)

func AddUser(user storage.CreateUserParams) error {
	_, err := storage.GetQueries().CreateUser(context.Background(), user)
	return err
}
