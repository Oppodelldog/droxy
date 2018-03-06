package dockerrun

import (
	"testing"

	"bytes"
	"io/ioutil"

	"github.com/Oppodelldog/droxy/config"
	"github.com/sirupsen/logrus"
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

func TestPrepareArguments_WithInvalidArgumentLength_ExpectWarning(t *testing.T) {

	logRecorderBuffer := bytes.NewBufferString("")
	logrus.SetOutput(logRecorderBuffer)
	invalidReplacementArgs := &[][]string{
		{
			"arg2",
		},
	}

	commandDef := &config.CommandDefinition{
		ReplaceArgs: invalidReplacementArgs,
	}

	arguments := []string{"arg2"}
	prepareCommandLineArguments(commandDef, arguments)

	recordedLogEntires, err := ioutil.ReadAll(logRecorderBuffer)
	if err != nil {
		t.Fatalf("Did not expect ioutil.ReadAll to return an error, but got: %v", err)
	}
	assert.Contains(t, string(recordedLogEntires), "invalid argument replacement mapping '[arg2]'. Replacement mapping must consist of 2 array entries.")
}
