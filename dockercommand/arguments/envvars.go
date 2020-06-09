package arguments

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder"
)

// BuildEnvVars adds environment variable mappings.
func BuildEnvVars(commandDef config.CommandDefinition, builder builder.Builder) error {
	if envVars, ok := commandDef.GetEnvVars(); ok {
		for _, envVar := range envVars {
			var (
				envVarValue string
				err         error
			)

			if envVarsRequired, ok := commandDef.GetRequireEnvVars(); ok && envVarsRequired {
				envVarValue, err = resolveEnvVarStrict(envVar)
			} else {
				envVarValue, err = resolveEnvVar(envVar)
			}

			if err != nil {
				return err
			}

			builder.AddEnvVar(envVarValue)
		}
	}

	return nil
}
