package arguments

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercmd/builder"
)

func BuildImpersonation(commandDef *config.CommandDefinition, builder *builder.Builder) error {
	return addImpersonation(commandDef, builder)
}
