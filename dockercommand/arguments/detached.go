package arguments

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder"
)

//BuildDetachedFlag sets the daemon flag, which starts the container in background.
func BuildDetachedFlag(commandDef config.CommandDefinition, builder builder.Builder) error {
	if isDetached, ok := commandDef.GetIsDetached(); isDetached && ok {
		builder.AddArgument("-d")
	}

	return nil
}
