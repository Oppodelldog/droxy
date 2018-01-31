package builder

import (
	"docker-proxy-command/config"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/drone/envsubst"
)

func BuildCommandFromConfig(commandName string, cfg *config.Configuration) (*exec.Cmd, error) {
	commandDef, err := cfg.FindCommandByName(commandName)
	if err != nil {
		return nil, err
	}

	commandBuilder := NewDockerCommandBuilder()
	cmd, err := buildCommandFromCommandDefinition(commandDef, commandBuilder)
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

func buildCommandFromCommandDefinition(commandDef *config.CommandDefinition, builder *DockerCommandBuilder) (*exec.Cmd, error) {

	var err error

	builder.SetEntryPoint(commandDef.EntryPoint)

	builder.AddCmdArguments(os.Args[1:])

	err = buildImage(commandDef.Image, builder)
	if err != nil {
		return nil, err
	}

	err = buildVolumes(commandDef.Volumes, builder)
	if err != nil {
		return nil, err
	}

	return builder.Build(), nil
}

func buildImage(imageName string, builder *DockerCommandBuilder) error {
	builder.SetImageName(imageName)

	return nil
}

func buildVolumes(volumes []string, builder *DockerCommandBuilder) error {
	for _, volume := range volumes {
		volumeParts := strings.Split(volume, ":")
		if len(volumeParts) < 2 || len(volumeParts) > 3 {
			return fmt.Errorf("invalid number of volume parts in '%s'", volume)
		}

		var hostPart, containerPart, options string
		var resolveErr error

		if len(volumeParts) > 0 {
			hostPart, resolveErr = resolve(volumeParts[0])
			if resolveErr != nil {
				return resolveErr
			}
		}
		if len(volumeParts) > 1 {
			containerPart, resolveErr = resolve(volumeParts[1])
			if resolveErr != nil {
				return resolveErr
			}
		}
		if len(volumeParts) > 2 {
			options, resolveErr = resolve(volumeParts[2])
			if resolveErr != nil {
				return resolveErr
			}
		}

		builder.AddVolumePlain(hostPart, containerPart, options)
	}

	return nil
}

func resolve(envVarName string) (string, error) {
	return envsubst.Eval(envVarName, func(normalizedEnvVarName string) string {
		if envVar, ok := os.LookupEnv(normalizedEnvVarName); !ok {
			panic(fmt.Sprintf("env var %v is not set!", normalizedEnvVarName))
		} else {
			return envVar
		}
	})
}
