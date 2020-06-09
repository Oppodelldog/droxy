package arguments

import (
	"testing"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder/mocks"
	"github.com/stretchr/testify/assert"
)

func TestBuildNetwork_NetworkIsTrue(t *testing.T) {
	network := "my-docker-network"
	commandDef := config.CommandDefinition{
		Network: &network,
	}
	builder := &mocks.Builder{}

	builder.On("SetNetwork", network).Return(builder)

	err := BuildNetwork(commandDef, builder)
	if err != nil {
		t.Fatalf("Did not expect BuildNetwork to return an error, but got: %v", err)
	}

	builder.AssertExpectations(t)
}

func TestBuildNetwork_NetworkIsFalse(t *testing.T) {
	commandDef := config.CommandDefinition{}
	builder := &mocks.Builder{}

	err := BuildNetwork(commandDef, builder)
	if err != nil {
		t.Fatalf("Did not expect BuildNetwork to return an error, but got: %v", err)
	}

	assert.Empty(t, builder.Calls)
}
