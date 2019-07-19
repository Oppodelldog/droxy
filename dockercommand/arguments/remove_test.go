package arguments

import (
	"testing"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder/mocks"
	"github.com/stretchr/testify/assert"
)

func TestBuildRemoveContainerFlag_RemoveIsTrue(t *testing.T) {
	removeContainer := true
	commandDef := &config.CommandDefinition{
		RemoveContainer: &removeContainer,
	}

	builder := &mocks.Builder{}

	builder.On("AddArgument", "--rm").Return(builder)

	err := BuildRemoveContainerFlag(commandDef, builder)
	if err != nil {
		t.Fatalf("Did not expect BuildRemoveContainerFlag to return an error, but got: %v", err)
	}

	builder.AssertExpectations(t)
}

func TestBuildRemoveContainerFlag_RemoveIsFalse(t *testing.T) {
	removeContainer := false
	commandDef := &config.CommandDefinition{
		RemoveContainer: &removeContainer,
	}

	builder := &mocks.Builder{}

	err := BuildRemoveContainerFlag(commandDef, builder)
	if err != nil {
		t.Fatalf("Did not expect BuildRemoveContainerFlag to return an error, but got: %v", err)
	}

	assert.Empty(t, builder.Calls)
}
