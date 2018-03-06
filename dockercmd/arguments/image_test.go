package arguments

import (
	"testing"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercmd/builder/mocks"
	"github.com/stretchr/testify/assert"
)

func TestBuildImage_ImageDefined(t *testing.T) {
	imageName := "imageName"
	commandDef := &config.CommandDefinition{
		Image: &imageName,
	}

	builder := &mocks.Builder{}
	builder.On("SetImageName", imageName).Return(builder)

	BuildImage(commandDef, builder)

	builder.AssertExpectations(t)
}

func TestBuildImage_NoImageNameDefined(t *testing.T) {
	commandDef := &config.CommandDefinition{Image: nil}
	builder := &mocks.Builder{}

	BuildImage(commandDef, builder)

	assert.Empty(t, builder.Calls)
}
