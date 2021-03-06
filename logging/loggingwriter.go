package logging

import (
	"io"

	"github.com/sirupsen/logrus"
)

// NewWriter returns a new Writer that also logs the output it writes.
func NewWriter(targetWriter io.Writer, logger *logrus.Logger, loggingPrefix string) Writer {
	return Writer{targetWriter, logger, loggingPrefix}
}

type Writer struct {
	targetWriter  io.Writer
	logger        *logrus.Logger
	loggingPrefix string
}

func (w Writer) Write(p []byte) (n int, err error) {
	w.logger.Infof("%s:%s", w.loggingPrefix, string(p))

	return w.targetWriter.Write(p)
}
