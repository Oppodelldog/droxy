package dockercmd

import "docker-proxy-command/config"

func addGroups(commandDef *config.CommandDefinition, builder *Builder) error {
	if addGroups, ok := commandDef.GetAddGroups(); ok {
		err := buildGroups(addGroups, builder)
		if err != nil {
			return err
		}
	}
	return nil
}
