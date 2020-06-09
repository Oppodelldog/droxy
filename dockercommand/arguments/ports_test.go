package arguments

import (
	"os"
	"testing"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder/mocks"
	"github.com/stretchr/testify/assert"
)

func TestBuildPorts_portsDefined(t *testing.T) {
	ports := []string{"PORT_HOST:PORT_CONTAINER"}
	commandDef := config.CommandDefinition{
		Ports: &ports,
	}

	builder := &mocks.Builder{}

	builder.On("AddPortMapping", ports[0]).Return(builder)

	err := BuildPorts(commandDef, builder)
	if err != nil {
		t.Fatalf("Did not expect BuildPorts to return an error, but got: %v", err)
	}

	builder.AssertExpectations(t)
}

func TestBuildPorts_portsNotDefined(t *testing.T) {
	commandDef := config.CommandDefinition{
		Ports: nil,
	}

	builder := &mocks.Builder{}

	err := BuildPorts(commandDef, builder)
	if err != nil {
		t.Fatalf("Did not expect BuildPorts to return an error, but got: %v", err)
	}

	assert.Empty(t, builder.Calls)
}

func TestBuildPorts_portsWithEnvVarsDefined_ExpectEnvVarsTobeResolved(t *testing.T) {
	err := os.Setenv("HOST_PORT", "0815")
	if err != nil {
		t.Fatalf("Did not expect os.Setenv to return an error, but got: %v", err)
	}

	err = os.Setenv("CONTAINER_PORT", "4711")
	if err != nil {
		t.Fatalf("Did not expect os.Setenv to return an error, but got: %v", err)
	}

	ports := []string{"${HOST_PORT}:${CONTAINER_PORT}"}
	commandDef := config.CommandDefinition{
		Ports: &ports,
	}

	builder := &mocks.Builder{}
	builder.On("AddPortMapping", "0815:4711").Return(builder)

	err = BuildPorts(commandDef, builder)
	if err != nil {
		t.Fatalf("Did not expect BuildPorts to return an error, but got: %v", err)
	}

	builder.AssertExpectations(t)

	err = os.Unsetenv("HOST_PORT")
	if err != nil {
		t.Fatalf("Did not expect os.Unsetenv to return an error, but got: %v", err)
	}

	err = os.Unsetenv("CONTAINER_PORT")
	if err != nil {
		t.Fatalf("Did not expect os.Unsetenv to return an error, but got: %v", err)
	}
}
