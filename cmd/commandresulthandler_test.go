package cmd

import (
	"io/ioutil"
	"os/exec"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestNewCommandResultHandler(t *testing.T) {
	assert.IsType(t, new(commandResultHandler), NewCommandResultHandler())
}

func TestCommandResultHandler_HandleCommandResult_smokeTest(t *testing.T) {
	logrus.SetOutput(ioutil.Discard)
	commandResultHandler := NewCommandResultHandler()

	cmd := exec.Command("hostname")
	err := cmd.Run()
	exitCode := commandResultHandler.HandleCommandResult(cmd, err)

	assert.Equal(t, 0, exitCode)
}

func TestCommandResultHandler_HandleCommandResult_ExitCodeIsReturned(t *testing.T) {
	logrus.SetOutput(ioutil.Discard)
	commandResultHandler := NewCommandResultHandler()

	cmd := exec.Command("ping", "blackhole")
	err := cmd.Run()

	exitCode := commandResultHandler.HandleCommandResult(cmd, err)

	assert.Equal(t, 2, exitCode)
}

func TestCommandResultHandler_HandleCommandResult_ExtCodeError(t *testing.T) {
	logrus.SetOutput(ioutil.Discard)
	commandResultHandler := NewCommandResultHandler()

	cmd := exec.Command("horstname")
	err := cmd.Run()

	exitCode := commandResultHandler.HandleCommandResult(cmd, err)

	assert.Equal(t, ExtCodeError, exitCode)
}
