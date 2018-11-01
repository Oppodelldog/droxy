package arguments

import (
	"testing"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder/mocks"
	"github.com/stretchr/testify/assert"
)

func TestBuildIpIp_IpIsSet(t *testing.T) {
	Ip := ".env"
	commandDef := &config.CommandDefinition{
		IP: &Ip,
	}
	builder := &mocks.Builder{}

	builder.On("SetIP", Ip).Return(builder)

	BuildIP(commandDef, builder)

	builder.AssertExpectations(t)
}

func TestBuildNetwork_IpIsNotSet(t *testing.T) {
	commandDef := &config.CommandDefinition{}
	builder := &mocks.Builder{}

	BuildIP(commandDef, builder)

	assert.Empty(t, builder.Calls)
}
