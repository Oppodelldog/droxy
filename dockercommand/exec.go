package dockercommand

import (
	"os"
	"os/exec"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/arguments"
	"github.com/Oppodelldog/droxy/dockercommand/builder"
)

func NewExecBuilder(version string) ExecBuilder {
	return ExecBuilder{
		vc: versionChecker{dockerVersion: version},
	}
}

type ExecBuilder struct {
	vc versionChecker
}

func (b ExecBuilder) BuildCommandFromConfig(commandDef config.CommandDefinition) (*exec.Cmd, error) {
	commandBuilder := builder.New()

	args := prepareCommandLineArguments(commandDef, os.Args[1:])
	args = prependAdditionalArguments(commandDef, args)

	commandBuilder.AddCmdArguments(args)

	err := buildArgumentsFromFunctions(commandDef, commandBuilder, b.getExecArgumentBuilderFuncs())
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

	commandBuilder.SetDockerSubCommand(builder.DockerExecSubCommand)

	return commandBuilder.Build(), nil
}

func (b ExecBuilder) getExecArgumentBuilderFuncs() []argumentBuilderFunc {
	return []argumentBuilderFunc{
		arguments.BuildInteractiveFlag,
		arguments.BuildTerminalContext,
		arguments.BuildDetachedFlag,
		withVersionConstraint(arguments.BuildEnvVars, ">= 1.25", b.vc),
		arguments.BuildEnvFile,
		withVersionConstraint(arguments.BuildWorkDir, ">= 1.35", b.vc),
		arguments.BuildImpersonation,
		arguments.BuildCommand,
	}
}

func withVersionConstraint(
	buildArgument argumentBuilderFunc,
	versionConstraint string,
	vc versionChecker,
) argumentBuilderFunc {
	return func(commandDef config.CommandDefinition, builder builder.Builder) error {
		if vc.isVersionSupported(versionConstraint) {
			return buildArgument(commandDef, builder)
		}

		return nil
	}
}
