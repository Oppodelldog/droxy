package helper

import (
	"io"
	"github.com/sirupsen/logrus"
)

func NewLoggingWriter(targetWriter io.Writer, logger *logrus.Logger, logginPrefix string) io.Writer {
	return &LoggingWriter{targetWriter, logger, logginPrefix}
}

type LoggingWriter struct {
	targetWriter io.Writer
	logger       *logrus.Logger
	logginPrefix string
}

func (w *LoggingWriter) Write(p []byte) (n int, err error) {
	w.logger.Infof("%s:%s", w.logginPrefix, string(p))
	return w.targetWriter.Write(p)
}
