package logger

import (
	"io"

	"github.com/sirupsen/logrus"
)

func Info(args ...interface{}) {
	logrus.Info(args...)
}

func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

func Error(args ...interface{}) {
	logrus.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

func Warning(args ...interface{}) {
	logrus.Warning(args...)
}

func Warnf(format string, args ...interface{}) {
	logrus.Warnf(format, args...)
}

func SetOutput(out io.Writer) {
	logrus.SetOutput(out)
}

func StandardLogger() *logrus.Logger {
	return logrus.StandardLogger()
}
