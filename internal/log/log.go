package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/rocky114/craftsman/internal/pkg/path"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

func InitLog() {
	writer, err := rotatelogs.New(
		fmt.Sprintf("%s%s", path.GetLogPath(), "log.%Y%m%d"),
		rotatelogs.WithLinkName(fmt.Sprintf("%s%s", path.GetLogPath(), "craftsman.log")),
		rotatelogs.WithMaxAge(time.Duration(24)*time.Hour*7),
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
	)

	if err != nil {
		log.Fatalf("create rotate file err: %v\n", err)
	}

	logrus.SetOutput(io.MultiWriter(writer, os.Stderr))
	logrus.SetFormatter(new(logrus.JSONFormatter))
	//logrus.SetReportCaller(true)
	/*logrus.SetFormatter(&logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			return f.Function, filepath.Base(f.File) + fmt.Sprintf(":%d", f.Line)
		},
		SortingFunc: func(s []string) {
			s[0], s[1], s[2], s[3], s[4] = "time", "level", "func", "file", "msg"
		},
	})*/
}
