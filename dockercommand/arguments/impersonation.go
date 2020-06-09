package arguments

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder"
)

// BuildImpersonation uses the current user and its group inside the container (linux only).
func BuildImpersonation(commandDef config.CommandDefinition, builder builder.Builder) error {
	return addImpersonation(commandDef, builder)
}
