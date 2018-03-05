package arguments

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercmd/builder"
)

func BuildImage(commandDef *config.CommandDefinition, builder builder.Builder) error {
	if imageName, ok := commandDef.GetImage(); ok {
		builder.SetImageName(imageName)
	}

	return nil
}
