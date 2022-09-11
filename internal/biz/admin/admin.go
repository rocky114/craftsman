package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/rocky114/craftsman/internal/storage"
)

func GetCaptcha(c *gin.Context) {
	storage.GetUsers()
}
