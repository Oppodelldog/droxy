package arguments

import (
	"testing"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder/mocks"
	"github.com/stretchr/testify/assert"
)

func TestBuildIpIp_IpIsSet(t *testing.T) {
	IP := ".env"
	commandDef := &config.CommandDefinition{
		IP: &IP,
	}
	builder := &mocks.Builder{}

	builder.On("SetIP", IP).Return(builder)

	err := BuildIP(commandDef, builder)
	if err != nil {
		t.Fatalf("Did not expect BuildIP to return an error, but got: %v", err)
	}

	builder.AssertExpectations(t)
}

func TestBuildNetwork_IpIsNotSet(t *testing.T) {
	commandDef := &config.CommandDefinition{}
	builder := &mocks.Builder{}

	err := BuildIP(commandDef, builder)
	if err != nil {
		t.Fatalf("Did not expect BuildIP to return an error, but got: %v", err)
	}

	assert.Empty(t, builder.Calls)
}
