package dockercommand

import (
	"io"
	"testing"

	"github.com/Oppodelldog/droxy/logger"

	"bytes"
	"github.com/Oppodelldog/droxy/config"
	"github.com/stretchr/testify/assert"
)

func TestAddAdditionalArguments(t *testing.T) {
	commandDef := config.CommandDefinition{
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

func TestPrepareArguments(t *testing.T) {
	commandDef := config.CommandDefinition{
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

func TestPrepareArguments_WithInvalidArgumentLength_ExpectWarning(t *testing.T) {
	logRecorderBuffer := bytes.NewBufferString("")
	logger.SetOutput(logRecorderBuffer)

	invalidReplacementArgs := &[][]string{
		{
			"arg2",
		},
	}

	commandDef := config.CommandDefinition{
		ReplaceArgs: invalidReplacementArgs,
	}

	arguments := []string{"arg2"}
	prepareCommandLineArguments(commandDef, arguments)

	recordedLogEntries, err := io.ReadAll(logRecorderBuffer)
	if err != nil {
		t.Fatalf("Did not expect io.ReadAll to return an error, but got: %v", err)
	}

	assert.Contains(
		t,
		string(recordedLogEntries),
		"invalid argument replacement mapping '[arg2]'. Replacement mapping must consist of 2 array entries.",
	)
}
