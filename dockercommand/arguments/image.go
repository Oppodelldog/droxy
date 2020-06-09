package arguments

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder"
)

//BuildImage sets the docker image.
func BuildImage(commandDef config.CommandDefinition, builder builder.Builder) error {
	if imageName, ok := commandDef.GetImage(); ok {
		builder.SetImageName(imageName)
	}

	return nil
}
