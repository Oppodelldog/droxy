package arguments

import (
	"fmt"
	"os"

	"github.com/drone/envsubst"
)

func resolveEnvVarStrict(envVarName string) (string, error) {
	return envsubst.Eval(envVarName, func(normalizedEnvVarName string) string {
		if envVar, ok := os.LookupEnv(normalizedEnvVarName); !ok {
			panic(fmt.Sprintf("env var %v is not set!", normalizedEnvVarName))
		} else {
			return envVar
		}
	})
}

func resolveEnvVar(envVarName string) (string, error) {
	return envsubst.Eval(envVarName, func(normalizedEnvVarName string) string {
		if envVar, ok := os.LookupEnv(normalizedEnvVarName); ok {
			return envVar
		}

		return ""
	})
}
