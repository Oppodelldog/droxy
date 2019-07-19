package arguments

import (
	"testing"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder/mocks"
	"github.com/stretchr/testify/assert"
)

func TestBuildEnvFile_EnvFileIsSet(t *testing.T) {
	envFile := ".env"
	commandDef := &config.CommandDefinition{
		EnvFile: &envFile,
	}
	builder := &mocks.Builder{}

	builder.On("SetEnvFile", envFile).Return(builder)

	err := BuildEnvFile(commandDef, builder)
	if err != nil {
		t.Fatalf("Did not expect BuildEnvFile to return an error, but got: %v", err)
	}

	builder.AssertExpectations(t)
}

func TestBuildNetwork_EnvFileIsNotSet(t *testing.T) {
	commandDef := &config.CommandDefinition{}
	builder := &mocks.Builder{}

	err := BuildEnvFile(commandDef, builder)
	if err != nil {
		t.Fatalf("Did not expect BuildEnvFile to return an error, but got: %v", err)
	}

	assert.Empty(t, builder.Calls)
}
