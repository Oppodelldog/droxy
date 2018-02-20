package dockercmd

import "github.com/Oppodelldog/droxy/config"

func addGroups(commandDef *config.CommandDefinition, builder *Builder) error {
	if addGroups, ok := commandDef.GetAddGroups(); ok {
		err := buildGroups(addGroups, builder)
		if err != nil {
			return err
		}
	}
	return nil
}

func addImpersonation(commandDef *config.CommandDefinition, builder *Builder) error {
	if impersonate, ok := commandDef.GetImpersonate(); ok {
		err := buildImpersonation(impersonate, builder)
		if err != nil {
			return err
		}
	}

	return nil
}
