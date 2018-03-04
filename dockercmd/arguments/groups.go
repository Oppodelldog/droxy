package arguments

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercmd/builder"
)

func BuildGroups(commandDef *config.CommandDefinition, builder *builder.Builder) error {
	return addGroups(commandDef, builder)
}
