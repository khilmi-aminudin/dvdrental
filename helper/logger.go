package helper

import (
	"os"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()
	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	logger.SetLevel(logrus.TraceLevel)
	logger.SetOutput(file)
}

func Logger() *logrus.Logger {
	return logger
}

func LogTrace(msg interface{}) {
	logger.Trace(msg)
}
