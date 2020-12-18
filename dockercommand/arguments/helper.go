package arguments

import (
	"fmt"
	"github.com/Oppodelldog/droxy/config"
	"os"
	"strings"

	"github.com/drone/envsubst"
)

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

	return envVarResolver{
		overwrites: overwrites,
	}
}

type envVarResolver struct {
	overwrites map[string]string
}

func (r envVarResolver) resolve(envVarName, defaultValue string) string {
	if overwriteValue, ok := r.overwrites[envVarName]; ok {
		return overwriteValue
	}

	return defaultValue
}

func (r envVarResolver) resolveEnvVarStrict(envVarName string) (string, error) {
	return envsubst.Eval(envVarName, func(normalizedEnvVarName string) string {
		if envValue, ok := os.LookupEnv(normalizedEnvVarName); !ok {
			panic(fmt.Sprintf("env var %v is not set!", normalizedEnvVarName))
		} else {
			return r.resolve(normalizedEnvVarName, envValue)
		}
	})
}

func (r envVarResolver) resolveEnvVar(envVarName string) (string, error) {
	return envsubst.Eval(envVarName, func(normalizedEnvVarName string) string {
		if envValue, ok := os.LookupEnv(normalizedEnvVarName); ok {
			return r.resolve(normalizedEnvVarName, envValue)
		}

		return ""
	})
}
