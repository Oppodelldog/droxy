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

	err := BuildEntryPoint(commandDef, builder)
	if err != nil {
		t.Fatalf("Did not expect BuildEntryPoint to return an error, but got: %v", err)
	}

	builder.AssertExpectations(t)
}

func TestBuildEntryPoint_NoEntryPointDefined(t *testing.T) {
	commandDef := &config.CommandDefinition{Command: nil}
	builder := &mocks.Builder{}

	err := BuildEntryPoint(commandDef, builder)
	if err != nil {
		t.Fatalf("Did not expect BuildEntryPoint to return an error, but got: %v", err)
	}

	assert.Empty(t, builder.Calls)
}
