package arguments

import (
	"testing"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockerrun/builder/mocks"
	"github.com/stretchr/testify/assert"
)

func TestBuildEnvFile_EnvFileIsSet(t *testing.T) {
	envFile := ".env"
	commandDef := &config.CommandDefinition{
		EnvFile: &envFile,
	}
	builder := &mocks.Builder{}

	builder.On("SetEnvFile", envFile).Return(builder)

	BuildEnvFile(commandDef, builder)

	builder.AssertExpectations(t)
}

func TestBuildNetwork_EnvFileIsNotSet(t *testing.T) {
	commandDef := &config.CommandDefinition{}
	builder := &mocks.Builder{}

	BuildEnvFile(commandDef, builder)

	assert.Empty(t, builder.Calls)
}
