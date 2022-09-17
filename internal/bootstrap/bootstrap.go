package bootstrap

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rocky114/craftsman/internal/biz/admin"
	"github.com/rocky114/craftsman/internal/biz/university"
	_ "github.com/rocky114/craftsman/internal/config"
	"github.com/rocky114/craftsman/internal/storage"
	"github.com/sirupsen/logrus"
)

var router = gin.Default()

func init() {
	InitLog()

	storage.InitDatabase()
	storage.InitMigrate()

	routes := []func(r *gin.Engine){admin.GetRoutes(), university.GetRoutes()}

	for _, fn := range routes {
		fn(router)
	}
}

func StartingHttpService() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	ctx, cannel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cannel()
	if err := server.Shutdown(ctx); err != nil {
		logrus.Fatal("server forced to shutdown: ", err)
	}

	logrus.Println("server exiting")
}
