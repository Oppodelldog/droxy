package arguments

import (
	"testing"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockerrun/builder/mocks"
	"github.com/stretchr/testify/assert"
)

func TestAttachStreams(t *testing.T) {
	builder := &mocks.Builder{}
	commandDef := &config.CommandDefinition{}

	builder.
		On("AttachTo", "STDIN").Return(builder).
		On("AttachTo", "STDOUT").Return(builder).
		On("AttachTo", "STDERR").Return(builder)

	err := AttachStreams(commandDef, builder)

	assert.Nil(t, err)
	builder.AssertExpectations(t)
}
