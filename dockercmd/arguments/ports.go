package arguments

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercmd/builder"
)

func BuildPorts(commandDef *config.CommandDefinition, builder *builder.Builder) error {
	if ports, ok := commandDef.GetPorts(); ok {
		return buildPorts(ports, builder)
	}

	return nil
}

func buildPorts(portMappings []string, builder *builder.Builder) error {
	for _, portMapping := range portMappings {

		portMappingWithValues, resolveErr := resolveEnvVar(portMapping)
		if resolveErr != nil {
			return resolveErr
		}
		builder.AddPortMapping(portMappingWithValues)
	}

	return nil
}
