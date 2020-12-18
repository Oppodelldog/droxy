package arguments

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder"
)

// BuildPorts sets mappings of host ports to container ports.
func BuildPorts(commandDef config.CommandDefinition, builder builder.Builder) error {
	if ports, ok := commandDef.GetPorts(); ok {
		return buildPorts(ports, builder, commandDef)
	}

	return nil
}

func buildPorts(portMappings []string, builder builder.Builder, commandDef config.CommandDefinition) error {
	for _, portMapping := range portMappings {
		portMappingWithValues, resolveErr := newEnvVarResolver(commandDef).substitute(portMapping)
		if resolveErr != nil {
			return resolveErr
		}

		builder.AddPortMapping(portMappingWithValues)
	}

	return nil
}
