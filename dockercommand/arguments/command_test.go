package arguments

import (
	"testing"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder/mocks"
	"github.com/stretchr/testify/assert"
)

func TestBuildCommand_CommandDefined(t *testing.T) {
	command := "command"
	commandDef := config.CommandDefinition{
		Command: &command,
	}

	builder := &mocks.Builder{}
	builder.On("SetCommand", command).Return(builder)

	err := BuildCommand(commandDef, builder)
	if err != nil {
		t.Fatalf("Did not expect BuildCommand to return an error, but got: %v", err)
	}

	builder.AssertExpectations(t)
}

func TestBuildCommand_NoCommandDefined(t *testing.T) {
	commandDef := config.CommandDefinition{Command: nil}
	builder := &mocks.Builder{}

	err := BuildCommand(commandDef, builder)
	if err != nil {
		t.Fatalf("Did not expect BuildCommand to return an error, but got: %v", err)
	}

	assert.Empty(t, builder.Calls)
}
