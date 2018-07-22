package arguments

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockerrun/builder"
)

//BuildDaemonFlag sets the daemon flag, which starts the container in background
func BuildDaemonFlag(commandDef *config.CommandDefinition, builder builder.Builder) error {
	if isDaemon, ok := commandDef.GetIsDaemon(); isDaemon && ok {
		builder.AddArgument("-d")
	}

	return nil
}
