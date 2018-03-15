package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/spf13/cobra"
)

func TestNewRoot(t *testing.T) {
	assert.NotNil(t, NewRoot())
	assert.IsType(t, new(cobra.Command), NewRoot())
}

func TestNewRootCommands(t *testing.T) {
	rootCmd := NewRoot()
	assertCommand(t, "clones", rootCmd)
	assertCommand(t, "hardlinks", rootCmd)
	assertCommand(t, "symlinks", rootCmd)
}

func TestRoot_Use(t *testing.T) {
	assert.Equal(t, "droxy", NewRoot().Use)
}

func assertCommand(t *testing.T, commandName string, command *cobra.Command) {
	for _, subCommand := range command.Commands() {
		if subCommand.Name() == commandName {
			return
		}
	}

	t.Fatalf("command '%v' not found", commandName)
}
