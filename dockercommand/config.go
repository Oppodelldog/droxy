package dockercommand

import (
	"os"
	"os/exec"

	"github.com/Masterminds/semver"
	"github.com/sirupsen/logrus"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/arguments"
	"github.com/Oppodelldog/droxy/dockercommand/builder"
)

//NewBuilder returns a new *Builder.
func NewBuilder() (*Builder, error) {
	clientAdapter, err := newDockerClientAdapter()
	if err != nil {
		return nil, err
	}

	return &Builder{
		dockerVersionProvider:     clientAdapter,
		containerExistenceChecker: clientAdapter,
	}, nil
}

type (
	dockerVersionProvider interface {
		getAPIVersion() (string, error)
	}

	containerExistenceChecker interface {
		exists(containerName string) bool
	}

	Builder struct {
		dockerVersionProvider     dockerVersionProvider
		containerExistenceChecker containerExistenceChecker
	}

	argumentBuilderDef func(commandDef config.CommandDefinition, builder builder.Builder) error
)

// BuildCommandFromConfig builds a docker-run command on base of the given CommandDefinition.
// If a container with the same name already exists a docker-exec command will be created.
func (cb *Builder) BuildCommandFromConfig(commandDef config.CommandDefinition) (*exec.Cmd, error) {
	cmd, err := cb.buildRunCommand(commandDef)
	if err != nil {
		return nil, err
	}

	if containerName, ok := commandDef.GetName(); ok {
		if cb.containerExistenceChecker.exists(containerName) {
			cmd, err = cb.buildExecCommand(commandDef)
			if err != nil {
				return nil, err
			}
		}
	}

	return cmd, nil
}

func (cb *Builder) buildRunCommand(commandDef config.CommandDefinition) (*exec.Cmd, error) {
	commandBuilder := builder.New()

	args := prepareCommandLineArguments(commandDef, os.Args[1:])
	args = prependAdditionalArguments(commandDef, args)

	commandBuilder.AddCmdArguments(args)

	err := buildRunArgumentsFromFunctions(commandDef, commandBuilder)
	if err != nil {
		return nil, err
	}

	err = buildRunArgumentsFromBuilders(commandDef, commandBuilder)
	if err != nil {
		return nil, err
	}

	return commandBuilder.Build(), nil
}

func (cb *Builder) buildExecCommand(commandDef config.CommandDefinition) (*exec.Cmd, error) {
	commandBuilder := builder.New()

	args := prepareCommandLineArguments(commandDef, os.Args[1:])
	args = prependAdditionalArguments(commandDef, args)

	commandBuilder.AddCmdArguments(args)

	err := cb.buildExecArgumentsFromFunctions(commandDef, commandBuilder)
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

func (cb *Builder) buildExecArgumentsFromFunctions(
	commandDef config.CommandDefinition,
	builder builder.Builder,
) error {
	argumentBuilderFunctions := []argumentBuilderDef{
		arguments.BuildInteractiveFlag,
		arguments.BuildTerminalContext,
		arguments.BuildDetachedFlag,
		cb.withVersionConstraint(arguments.BuildEnvVars, ">= 1.25"),
		arguments.BuildEnvFile,
		cb.withVersionConstraint(arguments.BuildWorkDir, ">= 1.35"),
		arguments.BuildImpersonation,
		arguments.BuildCommand,
	}

	for _, argumentBuilderFunc := range argumentBuilderFunctions {
		err := argumentBuilderFunc(commandDef, builder)
		if err != nil {
			return err
		}
	}

	return nil
}

func (cb *Builder) withVersionConstraint(
	argumentBuilderFunc argumentBuilderDef,
	versionConstraint string,
) argumentBuilderDef {
	return func(commandDef config.CommandDefinition, builder builder.Builder) error {
		if cb.isVersionSupported(versionConstraint) {
			return argumentBuilderFunc(commandDef, builder)
		}

		return nil
	}
}

func (cb *Builder) isVersionSupported(versionConstraint string) bool {
	constraints, err := semver.NewConstraint(versionConstraint)
	if err != nil {
		logrus.Errorf("unable to check version constraint '%s': %v", versionConstraint, err)

		return false
	}

	dockerVersion, err := cb.dockerVersionProvider.getAPIVersion()
	if err != nil {
		logrus.Errorf("unable to check version constraint '%s': %v", versionConstraint, err)

		return false
	}

	dockerSemVer, err := semver.NewVersion(dockerVersion)
	if err != nil {
		logrus.Errorf("unable to check version constraint '%s': %v", versionConstraint, err)

		return false
	}

	return constraints.Check(dockerSemVer)
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

func buildRunArgumentsFromFunctions(
	commandDef config.CommandDefinition,
	builder builder.Builder,
) error {
	argumentBuilderFunctions := []argumentBuilderDef{
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

	for _, argumentBuilderFunc := range argumentBuilderFunctions {
		err := argumentBuilderFunc(commandDef, builder)
		if err != nil {
			return err
		}
	}

	return nil
}
