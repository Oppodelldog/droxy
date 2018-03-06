package arguments

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercmd/builder/mocks"
	"github.com/testify/assert"
	"os"
	"testing"
)

func TestBuildPorts_portsDefined(t *testing.T) {
	ports := []string{"PORT_HOST:PORT_CONTAINER"}
	commandDef := &config.CommandDefinition{
		Ports: &ports,
	}

	builder := &mocks.Builder{}

	builder.On("AddPortMapping", ports[0]).Return(builder)

	BuildPorts(commandDef, builder)

	builder.AssertExpectations(t)
}

func TestBuildPorts_portsNotDefined(t *testing.T) {
	commandDef := &config.CommandDefinition{
		Ports: nil,
	}

	builder := &mocks.Builder{}

	BuildPorts(commandDef, builder)

	assert.Empty(t, builder.Calls)
}

func TestBuildPorts_portsWithEnvVarsDefined_ExpectEnvVarsTobeResolved(t *testing.T) {
	os.Setenv("HOST_PORT", "0815")
	os.Setenv("CONTAINER_PORT", "4711")

	ports := []string{"${HOST_PORT}:${CONTAINER_PORT}"}
	commandDef := &config.CommandDefinition{
		Ports: &ports,
	}

	builder := &mocks.Builder{}
	builder.On("AddPortMapping", "0815:4711").Return(builder)

	BuildPorts(commandDef, builder)

	builder.AssertExpectations(t)

	os.Unsetenv("HOST_PORT")
	os.Unsetenv("CONTAINER_PORT")
}
