package arguments

import (
	"testing"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder/mocks"
	"github.com/stretchr/testify/assert"
)

func TestBuildCommand_CommandDefined(t *testing.T) {

	command := "command"
	commandDef := &config.CommandDefinition{
		Command: &command,
	}

	builder := &mocks.Builder{}
	builder.On("SetCommand", command).Return(builder)

	BuildCommand(commandDef, builder)

	builder.AssertExpectations(t)
}

func TestBuildCommand_NoCommandDefined(t *testing.T) {
	commandDef := &config.CommandDefinition{Command: nil}
	builder := &mocks.Builder{}

	BuildCommand(commandDef, builder)

	assert.Empty(t, builder.Calls)
}
