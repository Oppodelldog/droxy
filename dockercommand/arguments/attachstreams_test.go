package arguments

import (
	"testing"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder/mocks"
	"github.com/stretchr/testify/assert"
)

func TestAttachStreams(t *testing.T) {
	builder := &mocks.Builder{}
	commandDef := &config.CommandDefinition{}

	builder.
		On("AttachTo", "STDIN").Once().Return(builder).
		On("AttachTo", "STDOUT").Once().Return(builder).
		On("AttachTo", "STDERR").Once().Return(builder)

	err := AttachStreams(commandDef, builder)

	assert.Nil(t, err)
	builder.AssertExpectations(t)
	builder.AssertNumberOfCalls(t, "AttachTo", 3)
}

func TestAttachStreams_inDetachedMode_notAttached(t *testing.T) {
	isDetached := true
	builder := &mocks.Builder{}
	commandDef := &config.CommandDefinition{
		IsDetached: &isDetached,
	}

	builder.
		On("AttachTo", "STDIN").Return(builder).
		On("AttachTo", "STDOUT").Return(builder).
		On("AttachTo", "STDERR").Return(builder)

	err := AttachStreams(commandDef, builder)

	assert.Nil(t, err)
	builder.AssertNotCalled(t, "AttachTo")
}

func TestAttachStreams_isInteractiveIsSetToFalse_notAttached(t *testing.T) {
	isInteractive := true
	builder := &mocks.Builder{}
	commandDef := &config.CommandDefinition{
		IsInteractive: &isInteractive,
	}

	builder.
		On("AttachTo", "STDIN").Return(builder).
		On("AttachTo", "STDOUT").Return(builder).
		On("AttachTo", "STDERR").Return(builder)

	err := AttachStreams(commandDef, builder)

	assert.Nil(t, err)
	builder.AssertNotCalled(t, "AttachTo")
}
