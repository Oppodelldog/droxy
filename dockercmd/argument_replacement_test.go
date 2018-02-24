package dockercmd

import (
	"testing"

	"github.com/Oppodelldog/droxy/config"
	"github.com/stretchr/testify/assert"
)

func TestPrepareArguments(t *testing.T) {

	commandDef := &config.CommandDefinition{
		ReplaceArgs: &[][]string{
			{
				"arg2", "arg99",
			},
		},
	}
	arguments := []string{"arg1", "arg2", "arg3"}
	preparedArguments := prepareCommandLineArguments(commandDef, arguments)

	expectedArguments := []string{"arg1", "arg99", "arg3"}

	assert.Equal(t, expectedArguments, preparedArguments)
}
