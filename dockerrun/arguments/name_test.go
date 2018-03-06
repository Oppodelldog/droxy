package arguments

import (
	"testing"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockerrun/builder/mocks"
	"github.com/stretchr/testify/assert"
)

func TestBuildName_NameIsSet(t *testing.T) {
	containerName := "my-container"
	commandDef := &config.CommandDefinition{
		Name: &containerName,
	}
	builder := &mocks.Builder{}

	builder.On("SetContainerName", containerName).Return(builder)

	BuildName(commandDef, builder)

	builder.AssertExpectations(t)
}

func TestBuildName_NameIsNotSet(t *testing.T) {
	commandDef := &config.CommandDefinition{}
	builder := &mocks.Builder{}

	BuildName(commandDef, builder)

	assert.Empty(t, builder.Calls)
}
