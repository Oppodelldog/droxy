package dockercmd

import (
	"os"
	"os/exec"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercmd/arguments"
	"github.com/Oppodelldog/droxy/dockercmd/builder"
)

// BuildCommandFromConfig builds a docker-run command on base of the given configuration
func BuildCommandFromConfig(commandName string, cfg *config.Configuration) (*exec.Cmd, error) {
	commandDef, err := cfg.FindCommandByName(commandName)
	if err != nil {
		return nil, err
	}

	commandBuilder := builder.New()
	cmd, err := buildCommandFromCommandDefinition(commandDef, commandBuilder)
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

type argumentBuilderDef func(commandDef *config.CommandDefinition, builder builder.Builder) error

func buildCommandFromCommandDefinition(commandDef *config.CommandDefinition, builder builder.Builder) (*exec.Cmd, error) {

	args := prepareCommandLineArguments(commandDef, os.Args[1:])
	args = prependAdditionalArguments(commandDef, args)

	builder.AddCmdArguments(args)

	err := buildArgumentsFromFuncs(commandDef, builder)
	if err != nil {
		return nil, err
	}

	err = buildArgumentsFromBuilders(commandDef, builder)
	if err != nil {
		return nil, err
	}

	return builder.Build(), nil
}

func buildArgumentsFromBuilders(commandDef *config.CommandDefinition, builder builder.Builder) error {
	argumentBuilders := []arguments.ArgumentBuilderInterface{
		arguments.NewUserGroupsArgumentBuilder(),
	}

	for _, argumentBuilder := range argumentBuilders {
		err := argumentBuilder.BuildArgument(commandDef, builder)
		if err != nil {
			return err
		}
	}

	return nil
}

func buildArgumentsFromFuncs(commandDef *config.CommandDefinition, builder builder.Builder) error {
	argumentBuilderFuncs := []argumentBuilderDef{
		arguments.AttachStreams,
		arguments.BuildTerminalContext,
		arguments.BuildEntryPoint,
		arguments.BuildNetwork,
		arguments.BuildInteractiveFlag,
		arguments.BuildRemoveContainerFlag,
		arguments.BuildImpersonation,
		arguments.BuildImage,
		arguments.BuildEnvVars,
		arguments.BuildPorts,
		arguments.BuildVolumes,
	}

	for _, argumentBuilderFunc := range argumentBuilderFuncs {
		err := argumentBuilderFunc(commandDef, builder)
		if err != nil {
			return err
		}
	}

	return nil
}
