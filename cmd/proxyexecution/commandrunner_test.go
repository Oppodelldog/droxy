package proxyexecution

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCommandRunner(t *testing.T) {
	assert.IsType(t, commandRunner{}, newCommandRunner())
}

func TestCommandRunner_RunCommand_smokeTest(t *testing.T) {
	commandRunner := newCommandRunner()

	cmd := exec.Command("echo", "'1'")
	err := commandRunner.RunCommand(cmd)

	assert.NoError(t, err)
}
