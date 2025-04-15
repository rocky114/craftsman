package api

import (
	"github.com/labstack/echo/v4"
	"github.com/rocky114/craftman/internal/api/handlers"
	"github.com/rocky114/craftman/internal/database"
	"net/http"
)

func RegisterRoutes(e *echo.Echo, store *database.Store) {
	universityHandler := handlers.NewUniversityHandler(store.Queries)

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	universityGroup := e.Group("/api/university")
	universityGroup.POST("", universityHandler.CreateUniversity)
	universityGroup.GET("", universityHandler.ListUniversities)
	universityGroup.GET("/:id", universityHandler.GetUniversity)
	universityGroup.DELETE("/:id", universityHandler.DeleteUniversity)
}
