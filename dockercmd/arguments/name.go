package arguments

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercmd/builder"
)

func BuildName(commandDef *config.CommandDefinition, builder builder.Builder) error {
	if containerName, ok := commandDef.GetName(); ok {
		builder.SetContainerName(containerName)
	}

	return nil
}
