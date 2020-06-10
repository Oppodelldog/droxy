package commands

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func Test_newRoot(t *testing.T) {
	assert.NotNil(t, newRoot())
	assert.IsType(t, new(cobra.Command), newRoot())
}

func Test_newRoot_Use(t *testing.T) {
	assert.Equal(t, "droxy", newRoot().Use)
}

func Test_newRootCommands(t *testing.T) {
	rootCmd := newRoot()
	assertCommand(t, "clones", rootCmd)
	assertCommand(t, "hardlinks", rootCmd)
	assertCommand(t, "symlinks", rootCmd)
}

func assertCommand(t *testing.T, commandName string, command *cobra.Command) {
	for _, subCommand := range command.Commands() {
		if subCommand.Name() == commandName {
			return
		}
	}

	t.Fatalf("command '%v' not found", commandName)
}
