package arguments

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercmd/builder"
)

func BuildRemoveContainerFlag(commandDef *config.CommandDefinition, builder builder.Builder) error {
	if removeContainer, ok := commandDef.GetRemoveContainer(); ok {
		if !removeContainer {
			return nil
		}

		builder.AddArgument("--rm")
	}

	return nil
}
