package arguments

import (
	"testing"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder/mocks"
	"github.com/stretchr/testify/assert"
)

func TestBuildIsDaemonFlag_InteractiveIsTrue(t *testing.T) {
	isDetached := true
	commandDef := &config.CommandDefinition{
		IsDetached: &isDetached,
	}
	builder := &mocks.Builder{}

	builder.On("AddArgument", "-d").Return(builder)

	err := BuildDetachedFlag(commandDef, builder)
	if err != nil {
		t.Fatalf("Did not expect BuildDetachedFlag to return an error, but got: %v", err)
	}

	builder.AssertExpectations(t)
}

func TestBuildIsDaemonFlag_InteractiveIsFalse(t *testing.T) {
	isDetached := false
	commandDef := &config.CommandDefinition{
		IsDetached: &isDetached,
	}
	builder := &mocks.Builder{}

	err := BuildDetachedFlag(commandDef, builder)
	if err != nil {
		t.Fatalf("Did not expect BuildDetachedFlag to return an error, but got: %v", err)
	}

	assert.Empty(t, builder.Calls)
}

// deprecated.
func TestBuildIsDaemonFlag_InteractiveIsTrue_deprecatedIsDaemon(t *testing.T) {
	isDaemon := true
	commandDef := &config.CommandDefinition{
		IsDaemon: &isDaemon,
	}
	builder := &mocks.Builder{}

	builder.On("AddArgument", "-d").Return(builder)

	err := BuildDetachedFlag(commandDef, builder)
	if err != nil {
		t.Fatalf("Did not expect BuildDetachedFlag to return an error, but got: %v", err)
	}

	builder.AssertExpectations(t)
}

// deprecated.
func TestBuildIsDaemonFlag_InteractiveIsFalse_deprecatedIsDaemon(t *testing.T) {
	isDaemon := false
	commandDef := &config.CommandDefinition{
		IsDaemon: &isDaemon,
	}
	builder := &mocks.Builder{}

	err := BuildDetachedFlag(commandDef, builder)
	if err != nil {
		t.Fatalf("Did not expect BuildDetachedFlag to return an error, but got: %v", err)
	}

	assert.Empty(t, builder.Calls)
}
