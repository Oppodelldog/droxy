package arguments

import (
	"testing"
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockerrun/builder/mocks"
	"os/user"
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
