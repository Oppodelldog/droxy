package arguments

import (
	"testing"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder/mocks"
	"github.com/stretchr/testify/assert"
)

func TestBuildImage_ImageDefined(t *testing.T) {
	imageName := "imageName"
	commandDef := &config.CommandDefinition{
		Image: &imageName,
	}

	builder := &mocks.Builder{}
	builder.On("SetImageName", imageName).Return(builder)

	err := BuildImage(commandDef, builder)
	if err != nil {
		t.Fatalf("Did not expect BuildImage to return an error, but got: %v", err)
	}

	builder.AssertExpectations(t)
}

func TestBuildImage_NoImageNameDefined(t *testing.T) {
	commandDef := &config.CommandDefinition{Image: nil}
	builder := &mocks.Builder{}

	err := BuildImage(commandDef, builder)
	if err != nil {
		t.Fatalf("Did not expect BuildImage to return an error, but got: %v", err)
	}

	assert.Empty(t, builder.Calls)
}
