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
		Ip: &Ip,
	}
	builder := &mocks.Builder{}

	builder.On("SetIp", Ip).Return(builder)

	BuildIp(commandDef, builder)

	builder.AssertExpectations(t)
}

func TestBuildNetwork_IpIsNotSet(t *testing.T) {
	commandDef := &config.CommandDefinition{}
	builder := &mocks.Builder{}

	BuildIp(commandDef, builder)

	assert.Empty(t, builder.Calls)
}
