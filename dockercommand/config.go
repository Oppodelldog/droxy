package dockercommand

import (
	"context"
	"os"
	"os/exec"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/sirupsen/logrus"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/arguments"
	"github.com/Oppodelldog/droxy/dockercommand/builder"
)

//NewCommandBuilder returns a new commandBuilder
func NewCommandBuilder() CommandBuilder {
	return &commandBuilder{}
}

type (
	// CommandBuilder builds a "docker run" command for the given command name and configuration
	CommandBuilder interface {
		BuildCommandFromConfig(commandName string, cfg *config.Configuration) (*exec.Cmd, error)
	}

	commandBuilder struct{}

	argumentBuilderDef func(commandDef *config.CommandDefinition, builder builder.Builder) error
)

// BuildCommandFromConfig builds a docker-run command on base of the given configuration
func (cb *commandBuilder) BuildCommandFromConfig(commandName string, cfg *config.Configuration) (*exec.Cmd, error) {
	commandDef, err := cfg.FindCommandByName(commandName)
	if err != nil {
		return nil, err
	}

	cmd, err := cb.buildRunCommand(commandDef)
	if err != nil {
		return nil, err
	}

	if containerName, ok := commandDef.GetName(); ok {
		if containerExists(containerName) {
			cmd, err = cb.buildExecCommand(commandDef)
			if err != nil {
				return nil, err
			}
		}
	}

	return cmd, nil
}

func (cb *commandBuilder) buildRunCommand(commandDef *config.CommandDefinition) (*exec.Cmd, error) {

	commandBuilder := builder.New()

	args := prepareCommandLineArguments(commandDef, os.Args[1:])
	args = prependAdditionalArguments(commandDef, args)

	commandBuilder.AddCmdArguments(args)

	err := cb.buildRunArgumentsFromFuncs(commandDef, commandBuilder)
	if err != nil {
		return nil, err
	}

	err = cb.buildRunArgumentsFromBuilders(commandDef, commandBuilder)
	if err != nil {
		return nil, err
	}

	return commandBuilder.Build(), nil
}

func (cb *commandBuilder) buildExecCommand(commandDef *config.CommandDefinition) (*exec.Cmd, error) {

	commandBuilder := builder.New()

	args := prepareCommandLineArguments(commandDef, os.Args[1:])
	args = prependAdditionalArguments(commandDef, args)

	commandBuilder.AddCmdArguments(args)

	err := cb.buildExecArgumentsFromFuncs(commandDef, commandBuilder)
	if err != nil {
		return nil, err
	}

	if containerName, ok := commandDef.GetName(); ok {
		commandBuilder.SetImageName(containerName)
	}
	if entryPoint, ok := commandDef.GetEntryPoint(); ok {
		commandBuilder.SetCommand(entryPoint)
	} else if command, ok := commandDef.GetCommand(); ok {
		commandBuilder.SetCommand(command)
	}

	commandBuilder.SetDockerSubCommand("exec")

	return commandBuilder.Build(), nil
}

func (cb *commandBuilder) buildRunArgumentsFromBuilders(commandDef *config.CommandDefinition, builder builder.Builder) error {
	argumentBuilders := []arguments.ArgumentBuilderInterface{
		arguments.NewUserGroupsArgumentBuilder(),
		arguments.NewNameArgumentBuilder(),
	}

	for _, argumentBuilder := range argumentBuilders {
		err := argumentBuilder.BuildArgument(commandDef, builder)
		if err != nil {
			return err
		}
	}

	return nil
}

func containerExists(containerName string) bool {
	dockerClient, err := client.NewEnvClient()
	if err != nil {
		logrus.Errorf("error building name argument, opening docker client failed: %v", err)

		return false
	}

	ctx := context.Background()
	options := types.ContainerListOptions{
		All: true,
	}

	containers, err := dockerClient.ContainerList(ctx, options)
	if err != nil {
		logrus.Errorf("error loading container list: %v", err)

		return false
	}

	for _, container := range containers {
		for _, name := range container.Names {
			if name == "/"+containerName {
				return true
			}
		}
	}

	return false
}

func (cb *commandBuilder) buildRunArgumentsFromFuncs(commandDef *config.CommandDefinition, builder builder.Builder) error {
	argumentBuilderFuncs := []argumentBuilderDef{
		arguments.AttachStreams,
		arguments.BuildTerminalContext,
		arguments.BuildEntryPoint,
		arguments.BuildCommand,
		arguments.BuildNetwork,
		arguments.BuildEnvFile,
		arguments.BuildIp,
		arguments.BuildInteractiveFlag,
		arguments.BuildDaemonFlag,
		arguments.BuildRemoveContainerFlag,
		arguments.BuildImpersonation,
		arguments.BuildImage,
		arguments.BuildEnvVars,
		arguments.LabelContainer,
		arguments.BuildPorts,
		arguments.BuildPortsFromParams,
		arguments.BuildVolumes,
		arguments.BuildLinks,
		arguments.BuildWorkDir,
	}

	for _, argumentBuilderFunc := range argumentBuilderFuncs {
		err := argumentBuilderFunc(commandDef, builder)
		if err != nil {
			return err
		}
	}

	return nil
}

func (cb *commandBuilder) buildExecArgumentsFromFuncs(commandDef *config.CommandDefinition, builder builder.Builder) error {
	argumentBuilderFuncs := []argumentBuilderDef{
		arguments.BuildInteractiveFlag,
		arguments.BuildTerminalContext,
		arguments.BuildDaemonFlag,
		arguments.BuildEnvVars,
		arguments.BuildEnvFile,
		arguments.BuildWorkDir,
		arguments.BuildImpersonation,
		arguments.BuildCommand,
	}

	for _, argumentBuilderFunc := range argumentBuilderFuncs {
		err := argumentBuilderFunc(commandDef, builder)
		if err != nil {
			return err
		}
	}

	return nil
}
