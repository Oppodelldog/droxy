package arguments

import (
	"testing"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockerrun/builder/mocks"
	"github.com/stretchr/testify/assert"
)

func TestBuildIsDaemonFlag_InteractiveIsTrue(t *testing.T) {

	isDaemon := true
	commandDef := &config.CommandDefinition{
		IsDaemon: &isDaemon,
	}
	builder := &mocks.Builder{}

	builder.On("AddArgument", "-d").Return(builder)

	BuildDaemonFlag(commandDef, builder)

	builder.AssertExpectations(t)
}

func TestBuildIsDaemonFlag_InteractiveIsFalse(t *testing.T) {
	isDaemon := false
	commandDef := &config.CommandDefinition{
		IsDaemon: &isDaemon,
	}
	builder := &mocks.Builder{}

	BuildDaemonFlag(commandDef, builder)

	assert.Empty(t, builder.Calls)
}
