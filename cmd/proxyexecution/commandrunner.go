package proxyexecution

import (
	"os"
	"os/exec"

	"github.com/Oppodelldog/droxy/logger"

	"github.com/Oppodelldog/droxy/logging"
)

type commandRunner struct {
}

func newCommandRunner() commandRunner {
	return commandRunner{}
}

// RunCommand executes the given command, but connecting a bypass logger to log std-stream communication.
func (r commandRunner) RunCommand(cmd *exec.Cmd) error {
	cmd.Stdout = logging.NewWriter(os.Stdout, logger.StandardLogger(), "StdOut")
	cmd.Stderr = logging.NewWriter(os.Stderr, logger.StandardLogger(), "StdErr")
	err := cmd.Start()

	if err != nil {
		return err
	}

	return cmd.Wait()
}
