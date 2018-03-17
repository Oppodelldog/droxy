package proxyexecution

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCommandRunner(t *testing.T) {
	assert.IsType(t, new(commandRunner), NewCommandRunner())
}

func TestCommandRunner_RunCommand_smokeTest(t *testing.T) {

	commandRunner := NewCommandRunner()

	cmd := exec.Command("echo", "'1'")
	err := commandRunner.RunCommand(cmd)

	assert.NoError(t, err)
}
