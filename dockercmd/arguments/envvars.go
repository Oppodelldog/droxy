package arguments

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercmd/builder"
)

func BuildEnvVars(commandDef *config.CommandDefinition, builder builder.Builder) error {
	if envVars, ok := commandDef.GetEnvVars(); ok {
		for _, envVar := range envVars {
			envVarValue, err := resolveEnvVar(envVar)
			if err != nil {
				return err
			}
			builder.AddEnvVar(envVarValue)
		}
	}

	return nil
}
