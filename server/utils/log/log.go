package log

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var (
	Logger *logrus.Logger
)

func init() {
	Logger = logrus.New()
	DebugMode()
}

func DebugMode() {
	Logger.SetLevel(logrus.DebugLevel)
	Logger.SetReportCaller(true)
	Logger.SetOutput(os.Stdout)
}

func InfoMode() {
	Logger.SetLevel(logrus.InfoLevel)
	Logger.SetReportCaller(false)
	fw, _ := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	Logger.SetOutput(io.MultiWriter(os.Stdout, fw))
}
