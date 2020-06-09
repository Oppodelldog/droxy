package arguments

import (
	"testing"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder/mocks"
	"github.com/stretchr/testify/assert"
)

func TestBuildInteractiveFlag_InteractiveIsTrue(t *testing.T) {
	isInteractive := true
	commandDef := config.CommandDefinition{
		IsInteractive: &isInteractive,
	}
	builder := &mocks.Builder{}

	builder.On("AddArgument", "-i").Return(builder)

	err := BuildInteractiveFlag(commandDef, builder)
	if err != nil {
		t.Fatalf("Did not expect BuildInteractiveFlag to return an error, but got: %v", err)
	}

	builder.AssertExpectations(t)
}

func TestBuildInteractiveFlag_InteractiveIsFalse(t *testing.T) {
	isInteractive := false
	commandDef := config.CommandDefinition{
		IsInteractive: &isInteractive,
	}
	builder := &mocks.Builder{}

	err := BuildInteractiveFlag(commandDef, builder)
	if err != nil {
		t.Fatalf("Did not expect BuildInteractiveFlag to return an error, but got: %v", err)
	}

	assert.Empty(t, builder.Calls)
}
