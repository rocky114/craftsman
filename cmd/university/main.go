package main

import (
	"github.com/rocky114/craftsman/internal/bootstrap"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.WithField("name", "rocky").Info("this is test")
	bootstrap.StartingHttpService()
}
