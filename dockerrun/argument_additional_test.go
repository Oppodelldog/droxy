package dockerrun

import (
	"testing"

	"github.com/Oppodelldog/droxy/config"
	"github.com/stretchr/testify/assert"
)

func TestAddAdditionalArguments(t *testing.T) {
	commandDef := &config.CommandDefinition{
		AdditionalArgs: &[]string{"--some-additional-argument"},
	}
	args := []string{"--some-arg"}

	extendedArgs := prependAdditionalArguments(commandDef, args)

	expectedArgs := []string{
		"--some-additional-argument",
		"--some-arg",
	}

	assert.Equal(t, expectedArgs, extendedArgs)
}
