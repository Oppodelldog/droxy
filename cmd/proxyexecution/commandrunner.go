package proxyexecution

import (
	"os"
	"os/exec"

	"github.com/Oppodelldog/droxy/logging"

	"github.com/sirupsen/logrus"
)

type commandRunner struct {
}

//NewCommandRunner returns a new commandRunner which can run a exec.Cmd
func NewCommandRunner() CommandRunner {
	return &commandRunner{}
}

//RunCommand executes the given command, but connecting a bypass logger to log std-stream communication.
func (r *commandRunner) RunCommand(cmd *exec.Cmd) error {

	cmd.Stdout = logging.NewLoggingWriter(os.Stdout, logrus.StandardLogger(), "StdOut")
	cmd.Stderr = logging.NewLoggingWriter(os.Stderr, logrus.StandardLogger(), "StdErr")
	err := cmd.Start()
	if err != nil {
		return err
	}

	return cmd.Wait()
}
