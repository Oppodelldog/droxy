package proxyexecution

import (
	"io"
	"os/exec"
	"testing"

	"github.com/Oppodelldog/droxy/logger"

	"github.com/stretchr/testify/assert"
)

func TestNewCommandResultHandler(t *testing.T) {
	assert.IsType(t, commandResultHandler{}, newResultHandler())
}

func TestCommandResultHandler_HandleCommandResult_smokeTest(t *testing.T) {
	logger.SetOutput(io.Discard)

	commandResultHandler := newResultHandler()

	cmd := exec.Command("hostname")
	err := cmd.Run()
	exitCode := commandResultHandler.HandleCommandResult(cmd, err)

	assert.Equal(t, 0, exitCode)
}

func TestCommandResultHandler_HandleCommandResult_ExitCodeIsReturned(t *testing.T) {
	logger.SetOutput(io.Discard)

	commandResultHandler := newResultHandler()

	cmd := exec.Command("ping", "blackHole")
	err := cmd.Run()

	exitCode := commandResultHandler.HandleCommandResult(cmd, err)

	assert.NotEqual(t, 0, exitCode)
}

func TestCommandResultHandler_HandleCommandResult_ExtCodeError(t *testing.T) {
	logger.SetOutput(io.Discard)

	commandResultHandler := newResultHandler()

	cmd := exec.Command("horstName")
	err := cmd.Run()

	exitCode := commandResultHandler.HandleCommandResult(cmd, err)

	assert.Equal(t, ExtCodeError, exitCode)
}
