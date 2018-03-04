package arguments

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercmd/builder"
	"os/user"
)

func addGroups(commandDef *config.CommandDefinition, builder *builder.Builder) error {
	if addGroups, ok := commandDef.GetAddGroups(); ok {
		err := buildGroups(addGroups, builder)
		if err != nil {
			return err
		}
	}
	return nil
}

func buildGroups(areGroupsAdded bool, builder *builder.Builder) error {
	if !areGroupsAdded {
		return nil
	}

	currentUser, err := user.Current()
	if err != nil {
		return err
	}

	groupIDs, err := currentUser.GroupIds()
	if err != nil {
		return err
	}

	if len(groupIDs) > 0 {
		for _, groupID := range groupIDs {
			builder.AddGroup(groupID)
		}
	}

	return nil
}

func addImpersonation(commandDef *config.CommandDefinition, builder *builder.Builder) error {
	if impersonate, ok := commandDef.GetImpersonate(); ok {
		err := buildImpersonation(impersonate, builder)
		if err != nil {
			return err
		}
	}

	return nil
}

func buildImpersonation(isImpersonated bool, builder *builder.Builder) error {
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
