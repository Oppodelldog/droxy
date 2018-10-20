package arguments

import (
	"testing"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder/mocks"
	"github.com/stretchr/testify/assert"
)

func TestBuildEntryPoint_EntryPointDefined(t *testing.T) {

	entryPoint := "entryPoint"
	commandDef := &config.CommandDefinition{
		EntryPoint: &entryPoint,
	}

	builder := &mocks.Builder{}
	builder.On("SetEntryPoint", entryPoint).Return(builder)

	BuildEntryPoint(commandDef, builder)

	builder.AssertExpectations(t)
}

func TestBuildEntryPoint_NoEntryPointDefined(t *testing.T) {
	commandDef := &config.CommandDefinition{Command: nil}
	builder := &mocks.Builder{}

	BuildEntryPoint(commandDef, builder)

	assert.Empty(t, builder.Calls)
}
