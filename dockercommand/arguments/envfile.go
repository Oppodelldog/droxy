package arguments

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder"
)

//BuildEnvFile maps the given env file into the container
func BuildEnvFile(commandDef *config.CommandDefinition, builder builder.Builder) error {
	if envFile, ok := commandDef.GetEnvFile(); ok {
		builder.SetEnvFile(envFile)
	}

	return nil
}
