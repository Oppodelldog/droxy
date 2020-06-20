package dockercommand

import (
	"testing"

	"github.com/Oppodelldog/droxy/config"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testDataSet := map[string]struct {
		dockerAPIVersion         string
		expectedCommandArgString string
		expectCommandContainsArg bool
	}{
		"workdir is not added for versions below 1.35": {
			dockerAPIVersion:         "1.34",
			expectedCommandArgString: "-w",
			expectCommandContainsArg: false,
		},
		"workdir is added for versions from 1.35": {
			dockerAPIVersion:         "1.35",
			expectedCommandArgString: "-w",
			expectCommandContainsArg: true,
		},
		"envvars are not added for versions below 1.25": {
			dockerAPIVersion:         "1.24",
			expectedCommandArgString: "-e",
			expectCommandContainsArg: false,
		},
		"envvars are added for versions from 1.25": {
			dockerAPIVersion:         "1.25",
			expectedCommandArgString: "-e",
			expectCommandContainsArg: true,
		},
	}

	for testCaseName, testData := range testDataSet {
		t.Run(testCaseName, func(t *testing.T) {
			cb := NewExecBuilder(testData.dockerAPIVersion)

			workDir := "someWorkDir"
			envVars := []string{"SOME_ENV_VAR"}
			commandDef := config.CommandDefinition{
				WorkDir: &workDir,
				EnvVars: &envVars,
			}
			cmd, err := cb.BuildCommandFromConfig(commandDef)
			if err != nil {
				t.Fatalf("did not expect buildExecCommand to return an error, but got: %v", err)
			}

			if testData.expectCommandContainsArg {
				assert.Contains(t, cmd.Args, testData.expectedCommandArgString)
			} else {
				assert.NotContains(t, cmd.Args, testData.expectedCommandArgString)
			}
		})
	}
}
