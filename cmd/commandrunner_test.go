package cmd

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"os/exec"
)

func TestNewCommandRunner(t *testing.T) {
	assert.Implements(t, new(CommandRunner), NewCommandRunner())
}

func TestCommandRunner_RunCommand_smokeTest(t *testing.T) {

	commandRunner := NewCommandRunner()

	cmd := exec.Command("echo", "'1'")
	err := commandRunner.RunCommand(cmd)

	assert.NoError(t, err)
}
