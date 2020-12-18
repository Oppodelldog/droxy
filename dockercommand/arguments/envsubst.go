package arguments

import (
	"fmt"
	"github.com/Oppodelldog/droxy/config"
	"os"
	"strings"

	"github.com/drone/envsubst"
)

type envVarResolver struct {
	overwrites      map[string]string
	requiresEnvVars bool
}

func newEnvVarResolver(definition config.CommandDefinition) envVarResolver {
	var overwrites = map[string]string{}

	if envVarDefs, ok := definition.GetEnvVarOverwrites(); ok {
		for _, def := range envVarDefs {
			parts := strings.Split(def, "=")
			if len(parts) != 2 {
				panic(fmt.Sprintf("invalid env var notation provided as '%s', notation must be a key value pair delimited by = 'VAR_NAME=VAR_VALUE'", def)) //nolint:lll
			}

			overwrites[parts[0]] = parts[1]
		}
	}

	requiresEnvVars, _ := definition.GetRequireEnvVars()

	return envVarResolver{
		overwrites:      overwrites,
		requiresEnvVars: requiresEnvVars,
	}
}

func (r envVarResolver) substitute(text string) (string, error) {
	return envsubst.Eval(text, func(normalizedEnvVarName string) string {
		if envValue, ok := os.LookupEnv(normalizedEnvVarName); ok {
			return r.resolve(normalizedEnvVarName, envValue)
		}

		envValue := r.resolve(normalizedEnvVarName, "")
		if envValue == "" && r.requiresEnvVars {
			panic(fmt.Sprintf("env var %v is not set!", normalizedEnvVarName))
		}

		return envValue
	})
}

func (r envVarResolver) resolve(envVarName, defaultValue string) string {
	if overwriteValue, ok := r.overwrites[envVarName]; ok {
		return overwriteValue
	}

	return defaultValue
}
