package bootstrap

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rocky114/craftsman/internal/config"
	"github.com/sirupsen/logrus"
)

func initLog() {
	writer, err := rotatelogs.New(
		config.GlobalConfig.Log.Path+".%Y%m%d",
		rotatelogs.WithLinkName(config.GlobalConfig.Log.Path),
		rotatelogs.WithMaxAge(time.Duration(180)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(60)*time.Second),
	)

	if err != nil {
		log.Fatalf("create rotate file err: %v\n", err)
	}

	logrus.SetOutput(writer)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			return f.Function, filepath.Base(f.File) + fmt.Sprintf(":%d", f.Line)
		},
		SortingFunc: func(s []string) {
			s[0], s[1], s[2], s[3], s[4] = "time", "level", "func", "file", "msg"
		},
	})
}
