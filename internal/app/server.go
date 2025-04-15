package app

import (
	"context"
	"errors"
	"github.com/rocky114/craftman/internal/api"
	"github.com/rocky114/craftman/internal/app/config"
	"github.com/rocky114/craftman/internal/database"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Application struct {
	echo  *echo.Echo
	cfg   *config.Config
	store *database.Store
}

func NewApplication(cfg *config.Config) (*Application, error) {
	store, err := database.NewStore(cfg.Database)
	if err != nil {
		return nil, err
	}

	e := echo.New()

	// 注册中间件
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// 注册路由
	api.RegisterRoutes(e, store)

	return &Application{
		echo:  e,
		store: store,
		cfg:   cfg,
	}, nil
}

func (a *Application) Start() {
	// Start server
	go func() {
		if err := a.echo.Start(":" + a.cfg.App.Port); err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.echo.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := a.echo.Shutdown(ctx); err != nil {
		a.echo.Logger.Fatal(err)
	}
}
