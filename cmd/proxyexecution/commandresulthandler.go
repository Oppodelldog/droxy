package proxyexecution

import (
	"os/exec"
	"syscall"

	"github.com/Oppodelldog/droxy/logger"
)

// ExtCodeError general error occurred when executing cmd.
const ExtCodeError = 993

// ExitCodeExitError of executed cmd could not be determined.
const ExitCodeExitError = 990

// ExitSuccessError ExitCode of successfully executed cmd could not be determined.
const ExitSuccessError = 991

func newResultHandler() commandResultHandler {
	return commandResultHandler{}
}

type (
	commandResultHandler struct{}
)

// HandleCommandResult tries to get to ExitCode of and already run cmd.
// Returns the exit code or a custom one if original exitCode could not be determined.
func (rh commandResultHandler) HandleCommandResult(cmd *exec.Cmd, err error) int {
	switch exitErr := err.(type) {
	case *exec.Error:
		logger.Warning("Could execute command")

		return ExtCodeError
	case *exec.ExitError:
		if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
			logger.Infof("docker finished with exit code '%v'", status.ExitStatus())

			return status.ExitStatus()
		}

		logger.Warning("Could not get exit code")

		return ExitCodeExitError
	}

	if status, ok := cmd.ProcessState.Sys().(syscall.WaitStatus); ok {
		logger.Infof("docker finished with exit code '%v'", status.ExitStatus())

		return status.ExitStatus()
	}

	logger.Warning("Could not get exit code")

	return ExitSuccessError
}
