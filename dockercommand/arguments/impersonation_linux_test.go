package arguments

import (
	"os/user"
	"testing"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder/mocks"
)

func TestBuildImpersonation(t *testing.T) {
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
