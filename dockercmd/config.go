package dockercmd

import (
	"docker-proxy-command/config"
	"docker-proxy-command/helper"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"

	"github.com/drone/envsubst"
)

// BuildCommandFromConfig builds a docker-run command on base of the given configuration
func BuildCommandFromConfig(commandName string, cfg *config.Configuration) (*exec.Cmd, error) {
	commandDef, err := cfg.FindCommandByName(commandName)
	if err != nil {
		return nil, err
	}

	commandBuilder := NewBuilder()
	cmd, err := buildCommandFromCommandDefinition(commandDef, commandBuilder)
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

func buildCommandFromCommandDefinition(commandDef *config.CommandDefinition, builder *Builder) (*exec.Cmd, error) {

	var err error

	builder.AddCmdArguments(os.Args[1:])

	err = autoBuildAttachStreams(builder)
	if err != nil {
		return nil, err
	}

	err = autoBuildTerminalContext(builder)
	if err != nil {
		return nil, err
	}

	if entryPoint, ok := commandDef.GetEntryPoint(); ok {
		builder.SetEntryPoint(entryPoint)
	}

	if network, ok := commandDef.GetNetwork(); ok {
		builder.SetNetwork(network)
	}

	if isInteractive, ok := commandDef.GetIsInteractive(); isInteractive && ok {
		builder.AddArgument("-i")
	}
	if removeContainer, ok := commandDef.GetRemoveContainer(); ok {
		err = buildRemoveContainer(removeContainer, builder)
		if err != nil {
			return nil, err
		}
	}

	if addGroups, ok := commandDef.GetAddGroups(); ok {
		err = buildGroups(addGroups, builder)
		if err != nil {
			return nil, err
		}
	}

	if impersonate, ok := commandDef.GetImpersonate(); ok {
		err = buildImpersonation(impersonate, builder)
		if err != nil {
			return nil, err
		}
	}

	if image, ok := commandDef.GetImage(); ok {
		err = buildImage(image, builder)
		if err != nil {
			return nil, err
		}
	}

	if volumes, ok := commandDef.GetVolumes(); ok {
		err = buildVolumes(volumes, builder)
		if err != nil {
			return nil, err
		}
	}

	if envVars, ok := commandDef.GetEnvVars(); ok {
		err = buildEnvVars(envVars, builder)
		if err != nil {
			return nil, err
		}
	}

	if ports, ok := commandDef.GetPorts(); ok {
		err = buildPorts(ports, builder)
		if err != nil {
			return nil, err
		}
	}

	return builder.Build(), nil
}

func buildEnvVars(envVars []string, builder *Builder) error {
	for _, envVar := range envVars {
		envVarValue, err := resolveEnvVar(envVar)
		if err != nil {
			return err
		}
		builder.AddEnvVar(envVarValue)
	}

	return nil
}

func autoBuildAttachStreams(builder *Builder) error {
	builder.
		AttachTo("STDIN").
		AttachTo("STDOUT").
		AttachTo("STDERR")

	return nil
}

func autoBuildTerminalContext(builder *Builder) error {
	if helper.IsTerminalContext() {
		builder.AddArgument("-t")
	}

	return nil
}

func buildRemoveContainer(isContainerRemoved bool, builder *Builder) error {
	if !isContainerRemoved {
		return nil
	}

	builder.AddArgument("--rm")

	return nil
}

func buildGroups(areGroupsAdded bool, builder *Builder) error {
	if !areGroupsAdded {
		return nil
	}

	currentUser, err := user.Current()
	if err != nil {
		return err
	}

	groupIDs, err := currentUser.GroupIds()
	if err != nil {
		return err
	}

	if len(groupIDs) > 0 {
		for _, groupID := range groupIDs {
			builder.AddGroup(groupID)
		}
	}

	return nil
}

func buildImpersonation(isImpersonated bool, builder *Builder) error {
	if !isImpersonated {
		return nil
	}

	currentUser, err := user.Current()
	if err != nil {
		return err
	}

	builder.SetContainerUserAndGroup(currentUser.Uid, currentUser.Gid)

	return nil
}

func buildImage(imageName string, builder *Builder) error {
	builder.SetImageName(imageName)

	return nil
}

func buildVolumes(volumes []string, builder *Builder) error {
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

func buildPorts(portMappings []string, builder *Builder) error {
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
