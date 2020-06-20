package logging_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/Oppodelldog/droxy/logger"

	"github.com/Oppodelldog/droxy/logging"

	"github.com/stretchr/testify/assert"

	"github.com/sirupsen/logrus"
)

func TestLoggingWriter_Write_WritesToWriter(t *testing.T) {
	loggerStub := getLoggerStub()

	var bytesBuffer []byte
	target := bytes.NewBuffer(bytesBuffer)
	loggingWriter := logging.NewWriter(target, loggerStub, "logPrefix")

	_, err := loggingWriter.Write([]byte("HELLO WORLD"))
	if err != nil {
		t.Fatalf("Did not expect Writer.Write to return an error, but got: %v", err)
	}

	bufferContent, err := ioutil.ReadAll(target)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "HELLO WORLD", string(bufferContent))
}

func TestLoggingWriter_Write_LogsInfo(t *testing.T) {
	log := logger.StandardLogger()
	log.Formatter = &testFormatter{}

	var bytesBuffer []byte
	outputBuffer := bytes.NewBuffer(bytesBuffer)
	log.Out = outputBuffer

	targetStub := ioutil.Discard
	loggingWriter := logging.NewWriter(targetStub, log, "logPrefix")

	_, err := loggingWriter.Write([]byte("HELLO WORLD"))
	if err != nil {
		t.Fatalf("Did not expect Writer.Write to return an error, but got: %v", err)
	}

	bufferContent, err := ioutil.ReadAll(outputBuffer)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "logPrefix:HELLO WORLD", string(bufferContent))
}

func getLoggerStub() *logrus.Logger {
	log := logger.StandardLogger()
	log.Out = ioutil.Discard

	return log
}

type testFormatter struct {
}

func (tf *testFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(entry.Message), nil
}
