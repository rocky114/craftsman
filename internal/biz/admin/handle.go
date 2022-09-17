package admin

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/rocky114/craftsman/internal/storage"
)

func GetCaptcha(c *gin.Context) {
	storage.GetQueries().ListUsers(context.Background())
}
