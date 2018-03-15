package cmd

import (
	"os/exec"
	"syscall"

	"github.com/sirupsen/logrus"
)

type (
	commandResultHandler struct{}
)

//NewCommandResultHandler handles an executed command and it's error code to get the executed commands exit code.
func NewCommandResultHandler() *commandResultHandler {
	return &commandResultHandler{}
}

// HandleCommandResult tries to get to ExitCode of and already run cmd. Returns the exit code or a custom one if original exitcode could not be determined.
func (rh *commandResultHandler) HandleCommandResult(cmd *exec.Cmd, err error) int {
	if exitErr, ok := err.(*exec.ExitError); ok {

		if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
			logrus.Infof("docker finished with exit code '%v'", status.ExitStatus())
			return status.ExitStatus()
		}

		logrus.Warning("Could not get exit code")
		return 990
	}

	if status, ok := cmd.ProcessState.Sys().(syscall.WaitStatus); ok {
		logrus.Infof("docker finished with exit code '%v'", status.ExitStatus())
		return status.ExitStatus()
	}

	logrus.Warning("Could not get exit code")
	return 991
}
