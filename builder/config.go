package builder

import (
	"docker-proxy-command/config"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/drone/envsubst"
	"os/user"
	"docker-proxy-command/helper"
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

	err = autoBuildInteractiveMode(builder)
	if err != nil {
		return nil, err
	}

	err = autoBuildAttachStreams(builder)
	if err != nil {
		return nil, err
	}

	err = autoBuildTerminalContext(builder)
	if err != nil {
		return nil, err
	}

	err = buildRemoveContainer(commandDef.RemoveContainer, builder)
	if err != nil {
		return nil, err
	}

	err = buildGroups(commandDef.AddGroups, builder)
	if err != nil {
		return nil, err
	}

	err = buildImpersonation(commandDef.Impersonate, builder)
	if err != nil {
		return nil, err
	}

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

func autoBuildInteractiveMode(builder *DockerCommandBuilder) error {
	builder.AddArgument("-i")

	return nil
}

func autoBuildAttachStreams(builder *DockerCommandBuilder) error {
	builder.
		AttachTo("STDIN").
		AttachTo("STDOUT").
		AttachTo("STDERR")

	return nil
}

func autoBuildTerminalContext(builder *DockerCommandBuilder) error {
	if helper.IsTerminalContext() {
		builder.AddArgument("-t")
	}

	return nil
}

func buildRemoveContainer(isContainerRemoved bool, builder *DockerCommandBuilder) error {
	if !isContainerRemoved {
		return nil
	}

	builder.AddArgument("--rm")

	return nil
}

func buildGroups(areGroupsAdded bool, builder *DockerCommandBuilder) error {
	if !areGroupsAdded {
		return nil
	}

	currentUser, err := user.Current()
	if err != nil {
		return err
	}

	groupIds, err := currentUser.GroupIds()
	if err != nil {
		return err
	}

	if len(groupIds) > 0 {
		for _, groupId := range groupIds {
			builder.AdduserGroup(groupId)
		}
	}

	return nil
}

func buildImpersonation(isImpersonated bool, builder *DockerCommandBuilder) error {
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
