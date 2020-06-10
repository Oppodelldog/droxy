package proxyexecution

import (
	"io/ioutil"
	"os/exec"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestNewCommandResultHandler(t *testing.T) {
	assert.IsType(t, ResultHandler{}, newResultHandler())
}

func TestCommandResultHandler_HandleCommandResult_smokeTest(t *testing.T) {
	logrus.SetOutput(ioutil.Discard)

	commandResultHandler := newResultHandler()

	cmd := exec.Command("hostname")
	err := cmd.Run()
	exitCode := commandResultHandler.HandleCommandResult(cmd, err)

	assert.Equal(t, 0, exitCode)
}

func TestCommandResultHandler_HandleCommandResult_ExitCodeIsReturned(t *testing.T) {
	logrus.SetOutput(ioutil.Discard)

	commandResultHandler := newResultHandler()

	cmd := exec.Command("ping", "blackHole")
	err := cmd.Run()

	exitCode := commandResultHandler.HandleCommandResult(cmd, err)

	assert.NotEqual(t, 0, exitCode)
}

func TestCommandResultHandler_HandleCommandResult_ExtCodeError(t *testing.T) {
	logrus.SetOutput(ioutil.Discard)

	commandResultHandler := newResultHandler()

	cmd := exec.Command("horstName")
	err := cmd.Run()

	exitCode := commandResultHandler.HandleCommandResult(cmd, err)

	assert.Equal(t, ExtCodeError, exitCode)
}
