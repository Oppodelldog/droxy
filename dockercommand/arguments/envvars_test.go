package arguments

import (
	"os"
	"testing"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder/mocks"
	"github.com/drone/envsubst/parse"
	"github.com/stretchr/testify/assert"
)

func TestBuildEnvVars_EnvVarsResolved_NoMatterIfTheyAreRequiredOrNot(t *testing.T) {
	val1 := "VALUE1"
	val2 := "VALUE2"
	err := os.Setenv("ENV_VAR_1", val1)
	if err != nil {
		t.Fatalf("Did not expect os.Setenv to return an error, but got: %v", err)
	}
	err = os.Setenv("ENV_VAR_2", val2)
	if err != nil {
		t.Fatalf("Did not expect os.Setenv to return an error, but got: %v", err)
	}

	defer func() {
		err := os.Unsetenv("ENV_VAR_1")
		if err != nil {
			t.Fatalf("Did not expect os.Unsetenv to return an error, but got: %v", err)
		}
		err = os.Unsetenv("ENV_VAR_2")
		if err != nil {
			t.Fatalf("Did not expect os.Unsetenv to return an error, but got: %v", err)
		}
	}()

	envVars := &[]string{
		"${ENV_VAR_1}",
		"${ENV_VAR_2}",
	}

	builder := &mocks.Builder{}
	builder.On("AddEnvVar", val1).Return(builder)
	builder.On("AddEnvVar", val2).Return(builder)

	testDataSet := map[string]struct {
		requireEnvVars bool
	}{
		"env vars are defined, but are not required": {true},
		"env vars are defined, and are required":     {false},
	}
	for testCaseName, testData := range testDataSet {
		t.Run(testCaseName, func(t *testing.T) {

			requireEnvVars := testData.requireEnvVars

			commandDef := &config.CommandDefinition{
				RequireEnvVars: &requireEnvVars,
				EnvVars:        envVars,
			}

			err := BuildEnvVars(commandDef, builder)
			if err != nil {
				t.Fatalf("Did not expect BuildEnvVars to return an error, but got: %v", err)
			}
		})
	}
}

func TestBuildEnvVars_EnvVarsNotRequired_EnvVarDefinedButCannotResolve_ResolvesEmptyString(t *testing.T) {

	envVars := &[]string{
		"${ENV_VAR_1}",
	}
	envVarsRequired := false
	commandDef := &config.CommandDefinition{
		RequireEnvVars: &envVarsRequired,
		EnvVars:        envVars,
	}
	builder := &mocks.Builder{}
	emptyString := ""
	builder.On("AddEnvVar", emptyString).Times(1).Return(builder)

	err := BuildEnvVars(commandDef, builder)
	if err != nil {
		t.Fatalf("Did not expect BuildEnvVars to return an error, but got: %v", err)
	}
}

func TestBuildEnvVars_EnvVarsRequired_EnvVarDefinedButCannotResolve_Panic(t *testing.T) {

	assert.Panics(t, func() {
		val1 := "VALUE1"
		envVars := &[]string{
			"${ENV_VAR_1}",
		}
		envVarsRequired := true
		commandDef := &config.CommandDefinition{
			RequireEnvVars: &envVarsRequired,
			EnvVars:        envVars,
		}
		builder := &mocks.Builder{}
		builder.On("AddEnvVar", val1).Return(builder)

		err := BuildEnvVars(commandDef, builder)
		if err != nil {
			t.Fatalf("Did not expect BuildEnvVars to return an error, but got: %v", err)
		}
	})
}

func TestBuildEnvVars_NoEnvVarsDefines(t *testing.T) {
	commandDef := &config.CommandDefinition{
		EnvVars: nil,
	}
	builder := &mocks.Builder{}
	err := BuildEnvVars(commandDef, builder)
	if err != nil {
		t.Fatalf("Did not expect BuildEnvVars to return an error, but got: %v", err)
	}

	assert.Empty(t, builder.Calls)
}

func TestBuildEnvVars_InvalidBashSubstitution_ExpectBadSubstitutionError(t *testing.T) {
	val1 := "VALUE1"
	err := os.Setenv("ENV_VAR_1", val1)
	if err != nil {
		t.Fatalf("Did not expect os.Setenv to return an error, but got: %v", err)
	}

	envVarsConfig := &[]string{
		"${ENV_VAR_1",
	}
	commandDef := &config.CommandDefinition{
		EnvVars: envVarsConfig,
	}
	builder := &mocks.Builder{}

	assert.IsType(t, parse.ErrBadSubstitution, BuildEnvVars(commandDef, builder))

	assert.Empty(t, builder.Calls)
}
