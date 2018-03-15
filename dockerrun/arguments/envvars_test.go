package arguments

import (
	"os"
	"testing"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockerrun/builder/mocks"
	"github.com/stretchr/testify/assert"
)

func TestBuildEnvVars_EnvVarsDefines(t *testing.T) {
	val1 := "VALUE1"
	val2 := "VALUE2"
	os.Setenv("ENV_VAR_1", val1)
	os.Setenv("ENV_VAR_2", val2)

	envVars := &[]string{
		"${ENV_VAR_1}",
		"${ENV_VAR_2}",
	}
	commandDef := &config.CommandDefinition{
		EnvVars: envVars,
	}
	builder := &mocks.Builder{}
	builder.On("AddEnvVar", val1).Return(builder)
	builder.On("AddEnvVar", val2).Return(builder)

	BuildEnvVars(commandDef, builder)

	os.Unsetenv("ENV_VAR_1")
	os.Unsetenv("ENV_VAR_2")
}

func TestBuildEnvVars_EnvVarDefinedButCannotResolve(t *testing.T) {

	assert.Panics(t, func() {
		val1 := "VALUE1"
		envVars := &[]string{
			"${ENV_VAR_1}",
		}
		commandDef := &config.CommandDefinition{
			EnvVars: envVars,
		}
		builder := &mocks.Builder{}
		builder.On("AddEnvVar", val1).Return(builder)

		BuildEnvVars(commandDef, builder)
	})
}

func TestBuildEnvVars_NoEnvVarsDefines(t *testing.T) {
	commandDef := &config.CommandDefinition{
		EnvVars: nil,
	}
	builder := &mocks.Builder{}
	BuildEnvVars(commandDef, builder)

	assert.Empty(t, builder.Calls)
}