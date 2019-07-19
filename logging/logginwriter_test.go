package logging

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sirupsen/logrus"
)

func TestLoggingWriter_Write_WritesToWriter(t *testing.T) {
	loggerStub := getLoggetStub()

	var bytesBuffer []byte
	target := bytes.NewBuffer(bytesBuffer)
	loggingWriter := NewLoggingWriter(target, loggerStub, "logprefix")

	_, err := loggingWriter.Write([]byte("HELLO WORLD"))
	if err != nil {
		t.Fatalf("Did not expect loggingWriter.Write to return an error, but got: %v", err)
	}

	bufferContent, err := ioutil.ReadAll(target)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "HELLO WORLD", string(bufferContent))
}

func TestLoggingWriter_Write_LogsInfo(t *testing.T) {
	logger := logrus.StandardLogger()
	logger.Formatter = &testFormatter{}
	var bytesBuffer []byte
	outpoutBuffer := bytes.NewBuffer(bytesBuffer)
	logger.Out = outpoutBuffer

	targetStub := ioutil.Discard
	loggingWriter := NewLoggingWriter(targetStub, logger, "logPrefix")

	_, err := loggingWriter.Write([]byte("HELLO WORLD"))
	if err != nil {
		t.Fatalf("Did not expect loggingWriter.Write to return an error, but got: %v", err)
	}

	bufferContent, err := ioutil.ReadAll(outpoutBuffer)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "logPrefix:HELLO WORLD", string(bufferContent))
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
