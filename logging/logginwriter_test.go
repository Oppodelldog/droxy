package logging

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/sirupsen/logrus"
)

func TestLoggingWriter_Write_WritesToWriter(t *testing.T) {
	loggerStub := getLoggetStub()

	bytesBuffer := []byte{}
	target := bytes.NewBuffer(bytesBuffer)
	loggingWriter := NewLoggingWriter(target, loggerStub, "logprefix")

	loggingWriter.Write([]byte("HELLO WORLD"))

	bufferContent, err := ioutil.ReadAll(target)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, string(bufferContent), "HELLO WORLD")
}

func TestLoggingWriter_Write_LogsInfo(t *testing.T) {
	logger := logrus.StandardLogger()
	logger.Formatter = &testFormatter{}
	bytesBuffer := []byte{}
	outpoutBuffer := bytes.NewBuffer(bytesBuffer)
	logger.Out = outpoutBuffer

	targetStub := ioutil.Discard
	loggingWriter := NewLoggingWriter(targetStub, logger, "logPrefix")

	loggingWriter.Write([]byte("HELLO WORLD"))

	bufferContent, err := ioutil.ReadAll(outpoutBuffer)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, string(bufferContent), "logPrefix:HELLO WORLD")
}

func getLoggetStub() *logrus.Logger {
	logger := logrus.StandardLogger()
	logger.Out = ioutil.Discard
	return logger
}

type testFormatter struct {
}

func (tf *testFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(entry.Message), nil
}
