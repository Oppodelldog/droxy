package arguments

import (
	"os/user"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercmd/builder"
)

func addImpersonation(commandDef *config.CommandDefinition, builder builder.Builder) error {
	if impersonate, ok := commandDef.GetImpersonate(); ok {
		err := buildImpersonation(impersonate, builder)
		if err != nil {
			return err
		}
	}

	return nil
}

func buildImpersonation(isImpersonated bool, builder builder.Builder) error {
	if !isImpersonated {
		return nil
	}

	currentUser, err := user.Current()
	if err != nil {
		return err
	}

	builder.SetContainerUserAndGroup(currentUser.Uid, currentUser.Gid)

	return nil
}
