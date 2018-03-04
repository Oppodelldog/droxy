package arguments

import (
	"fmt"
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercmd/builder"
	"github.com/Oppodelldog/droxy/helper"
	"github.com/drone/envsubst"
	"os"
	"strings"
)

func AttachStreams(commandDef *config.CommandDefinition, builder *builder.Builder) error {
	_ = commandDef
	builder.
		AttachTo("STDIN").
		AttachTo("STDOUT").
		AttachTo("STDERR")

	return nil
}

func BuildTerminalContext(commandDef *config.CommandDefinition, builder *builder.Builder) error {
	_ = commandDef
	if helper.IsTerminalContext() {
		builder.AddArgument("-t")
	}

	return nil
}

func BuildEntryPoint(commandDef *config.CommandDefinition, builder *builder.Builder) error {
	if entryPoint, ok := commandDef.GetEntryPoint(); ok {
		builder.SetEntryPoint(entryPoint)
	}

	return nil
}

func BuildNetwork(commandDef *config.CommandDefinition, builder *builder.Builder) error {
	if network, ok := commandDef.GetNetwork(); ok {
		builder.SetNetwork(network)
	}

	return nil
}

func BuildInteractiveFlag(commandDef *config.CommandDefinition, builder *builder.Builder) error {
	if isInteractive, ok := commandDef.GetIsInteractive(); isInteractive && ok {
		builder.AddArgument("-i")
	}

	return nil
}

func BuildRemoveContainerFlag(commandDef *config.CommandDefinition, builder *builder.Builder) error {
	if removeContainer, ok := commandDef.GetRemoveContainer(); ok {
		if !removeContainer {
			return nil
		}

		builder.AddArgument("--rm")
	}

	return nil
}

func BuildGroups(commandDef *config.CommandDefinition, builder *builder.Builder) error {
	return addGroups(commandDef, builder)
}

func BuildImpersonation(commandDef *config.CommandDefinition, builder *builder.Builder) error {
	return addImpersonation(commandDef, builder)
}

func BuildImage(commandDef *config.CommandDefinition, builder *builder.Builder) error {
	if imageName, ok := commandDef.GetImage(); ok {
		builder.SetImageName(imageName)
	}

	return nil
}

func BuildPorts(commandDef *config.CommandDefinition, builder *builder.Builder) error {
	if ports, ok := commandDef.GetPorts(); ok {
		return buildPorts(ports, builder)
	}

	return nil
}

func BuildEnvVars(commandDef *config.CommandDefinition, builder *builder.Builder) error {
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

func BuildVolumes(commandDef *config.CommandDefinition, builder *builder.Builder) error {
	if volumes, ok := commandDef.GetVolumes(); ok {
		return buildVolumes(volumes, builder)
	}
	return nil
}

func buildVolumes(volumes []string, builder *builder.Builder) error {
	for _, volume := range volumes {
		volumeParts := strings.Split(volume, ":")
		if len(volumeParts) < 2 || len(volumeParts) > 3 {
			return fmt.Errorf("invalid number of volume parts in '%s'", volume)
		}

		var hostPart, containerPart, options string
		var resolveErr error

		if len(volumeParts) > 0 {
			hostPart, resolveErr = resolveEnvVar(volumeParts[0])
			if resolveErr != nil {
				return resolveErr
			}
		}
		if len(volumeParts) > 1 {
			containerPart, resolveErr = resolveEnvVar(volumeParts[1])
			if resolveErr != nil {
				return resolveErr
			}
		}
		if len(volumeParts) > 2 {
			options, resolveErr = resolveEnvVar(volumeParts[2])
			if resolveErr != nil {
				return resolveErr
			}
		}

		builder.AddVolumeMapping(hostPart, containerPart, options)
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

func resolveEnvVar(envVarName string) (string, error) {
	return envsubst.Eval(envVarName, func(normalizedEnvVarName string) string {
		if envVar, ok := os.LookupEnv(normalizedEnvVarName); !ok {
			panic(fmt.Sprintf("env var %v is not set!", normalizedEnvVarName))
		} else {
			return envVar
		}
	})
}
