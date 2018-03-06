package arguments

import (
	"testing"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercmd/builder/mocks"
	"github.com/stretchr/testify/assert"
)

func TestBuildRemoveContainerFlag_RemoveIsTrue(t *testing.T) {
	removeContainer := true
	commandDef := &config.CommandDefinition{
		RemoveContainer: &removeContainer,
	}

	builder := &mocks.Builder{}

	builder.On("AddArgument", "--rm").Return(builder)

	BuildRemoveContainerFlag(commandDef, builder)

	builder.AssertExpectations(t)
}

func TestBuildRemoveContainerFlag_RemoveIsFalse(t *testing.T) {
	removeContainer := false
	commandDef := &config.CommandDefinition{
		RemoveContainer: &removeContainer,
	}

	builder := &mocks.Builder{}

	BuildRemoveContainerFlag(commandDef, builder)

	assert.Empty(t, builder.Calls)
}
