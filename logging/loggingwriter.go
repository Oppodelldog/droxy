package logging

import (
	"io"

	"github.com/sirupsen/logrus"
)

// NewLoggingWriter returns a new writer that also logs the output it writes
func NewLoggingWriter(targetWriter io.Writer, logger *logrus.Logger, logginPrefix string) io.Writer {
	return &loggingWriter{targetWriter, logger, logginPrefix}
}

type loggingWriter struct {
	targetWriter io.Writer
	logger       *logrus.Logger
	logginPrefix string
}

func (w *loggingWriter) Write(p []byte) (n int, err error) {
	w.logger.Infof("%s:%s", w.logginPrefix, string(p))
	return w.targetWriter.Write(p)
}
