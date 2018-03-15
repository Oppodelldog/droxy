package dockerrun

import (
	"os"
	"os/exec"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockerrun/arguments"
	"github.com/Oppodelldog/droxy/dockerrun/builder"
)

//NewCommandBuilder returns a new commandBuilder
func NewCommandBuilder() *commandBuilder {
	return &commandBuilder{}
}

type (
	commandBuilder struct{}

	argumentBuilderDef func(commandDef *config.CommandDefinition, builder builder.Builder) error
)

// BuildCommandFromConfig builds a docker-run command on base of the given configuration
func (cb *commandBuilder) BuildCommandFromConfig(commandName string, cfg *config.Configuration) (*exec.Cmd, error) {
	commandDef, err := cfg.FindCommandByName(commandName)
	if err != nil {
		return nil, err
	}

	commandBuilder := builder.New()
	cmd, err := cb.buildCommandFromCommandDefinition(commandDef, commandBuilder)
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

func (cb *commandBuilder) buildCommandFromCommandDefinition(commandDef *config.CommandDefinition, builder builder.Builder) (*exec.Cmd, error) {

	args := prepareCommandLineArguments(commandDef, os.Args[1:])
	args = prependAdditionalArguments(commandDef, args)

	builder.AddCmdArguments(args)

	err := cb.buildArgumentsFromFuncs(commandDef, builder)
	if err != nil {
		return nil, err
	}

	err = cb.buildArgumentsFromBuilders(commandDef, builder)
	if err != nil {
		return nil, err
	}

	return builder.Build(), nil
}

func (cb *commandBuilder) buildArgumentsFromBuilders(commandDef *config.CommandDefinition, builder builder.Builder) error {
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

func (cb *commandBuilder) buildArgumentsFromFuncs(commandDef *config.CommandDefinition, builder builder.Builder) error {
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
