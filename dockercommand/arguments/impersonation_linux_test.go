package arguments

import (
	"os/user"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder/mocks"
)

func TestBuildImpersonation_enabledByConfig_builderIsCalled(t *testing.T) {
	impersonate := true
	commandDef := &config.CommandDefinition{Impersonate: &impersonate}
	builder := &mocks.Builder{}

	usr, err := user.Current()
	if err != nil {
		t.Fatalf("Did not expect to return an error, but got: %v", err)
	}

	builder.On("SetContainerUserAndGroup", usr.Uid, usr.Gid).Return(builder)

	BuildImpersonation(commandDef, builder)

	builder.AssertExpectations(t)
}

func TestBuildImpersonation_disabledByConfig_builderIsNotCalled(t *testing.T) {
	impersonate := false
	commandDef := &config.CommandDefinition{Impersonate: &impersonate}
	builder := &mocks.Builder{}

	_, err := user.Current()
	if err != nil {
		t.Fatalf("Did not expect to return an error, but got: %v", err)
	}

	BuildImpersonation(commandDef, builder)

	builder.AssertNotCalled(t, "SetContainerUserAndGroup", mock.Anything, mock.Anything)
}
