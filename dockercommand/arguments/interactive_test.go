package arguments

import (
	"testing"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder/mocks"
	"github.com/stretchr/testify/assert"
)

func TestBuildInteractiveFlag_InteractiveIsTrue(t *testing.T) {

	isInteractive := true
	commandDef := &config.CommandDefinition{
		IsInteractive: &isInteractive,
	}
	builder := &mocks.Builder{}

	builder.On("AddArgument", "-i").Return(builder)

	BuildInteractiveFlag(commandDef, builder)

	builder.AssertExpectations(t)
}

func TestBuildInteractiveFlag_InteractiveIsFalse(t *testing.T) {
	isInteractive := false
	commandDef := &config.CommandDefinition{
		IsInteractive: &isInteractive,
	}
	builder := &mocks.Builder{}

	BuildInteractiveFlag(commandDef, builder)

	assert.Empty(t, builder.Calls)
}