package arguments

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder"
)

//BuildIP maps the given Ip file into the container
func BuildIP(commandDef *config.CommandDefinition, builder builder.Builder) error {
	if ip, ok := commandDef.GetIP(); ok {
		builder.SetIP(ip)
	}

	return nil
}
