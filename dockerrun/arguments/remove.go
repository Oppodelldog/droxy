package arguments

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockerrun/builder"
)

//BuildRemoveContainerFlag adds --rm flag to remove container after it terminated
func BuildRemoveContainerFlag(commandDef *config.CommandDefinition, builder builder.Builder) error {
	if removeContainer, ok := commandDef.GetRemoveContainer(); ok {
		if !removeContainer {
			return nil
		}

		builder.AddArgument("--rm")
	}

	return nil
}
