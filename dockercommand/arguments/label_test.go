package arguments

import (
	"testing"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder/mocks"
	"github.com/stretchr/testify/assert"
)

func TestLabelContainer(t *testing.T) {
	builder := &mocks.Builder{}
	commandDef := &config.CommandDefinition{}

	builder.On("AddLabel", containerLabel).Return(builder)

	err := LabelContainer(commandDef, builder)

	assert.Nil(t, err)
	builder.AssertExpectations(t)
}
