package arguments

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder"
)

//BuildIp maps the given Ip file into the container
func BuildIp(commandDef *config.CommandDefinition, builder builder.Builder) error {
	if ip, ok := commandDef.GetIp(); ok {
		builder.SetIp(ip)
	}

	return nil
}
