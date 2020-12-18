package arguments

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder"
)

// BuildLinks maps Links from host to container.
func BuildLinks(commandDef config.CommandDefinition, builder builder.Builder) error {
	if Links, ok := commandDef.GetLinks(); ok {
		for _, volume := range Links {
			resolvedLinkMapping, err := newEnvVarResolver(commandDef).substitute(volume)
			if err != nil {
				return err
			}

			builder.AddLinkMapping(resolvedLinkMapping)
		}
	}

	return nil
}
