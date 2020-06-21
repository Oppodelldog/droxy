package dockercommand

import (
	"os"
	"os/exec"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/arguments"
	"github.com/Oppodelldog/droxy/dockercommand/builder"
)

func NewRunBuilder(_ string) RunBuilder {
	return RunBuilder{}
}

type RunBuilder struct {
}

func (b RunBuilder) BuildCommandFromConfig(commandDef config.CommandDefinition) (*exec.Cmd, error) {
	commandBuilder := builder.New()

	args := prepareCommandLineArguments(commandDef, os.Args[1:])
	args = prependAdditionalArguments(commandDef, args)

	commandBuilder.AddCmdArguments(args)

	err := buildArgumentsFromFunctions(commandDef, commandBuilder, getRunArgumentBuilders())
	if err != nil {
		return nil, err
	}

	err = buildRunArgumentsFromBuilders(commandDef, commandBuilder)
	if err != nil {
		return nil, err
	}

	return commandBuilder.Build(), nil
}

func buildRunArgumentsFromBuilders(
	commandDef config.CommandDefinition,
	builder builder.Builder,
) error {
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

func getRunArgumentBuilders() []argumentBuilderFunc {
	return []argumentBuilderFunc{
		arguments.AttachStreams,
		arguments.BuildTerminalContext,
		arguments.BuildEntryPoint,
		arguments.BuildCommand,
		arguments.BuildNetwork,
		arguments.BuildEnvFile,
		arguments.BuildIP,
		arguments.BuildInteractiveFlag,
		arguments.BuildDetachedFlag,
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
}
