package dockercommand

import (
	"os/exec"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder"
)

type (
	dockerVersionProvider interface {
		getAPIVersion() (string, error)
	}

	containerExistenceChecker interface {
		exists(containerName string) bool
	}

	Builder struct {
		containerExistenceChecker containerExistenceChecker
		versionProvider           dockerVersionProvider
	}

	argumentBuilderFunc func(commandDef config.CommandDefinition, builder builder.Builder) error
)

//NewBuilder returns a new *Builder.
func NewBuilder() (*Builder, error) {
	clientAdapter, err := newDockerClientAdapter()
	if err != nil {
		return nil, err
	}

	return &Builder{
		containerExistenceChecker: clientAdapter,
		versionProvider:           clientAdapter,
	}, nil
}

// BuildCommandFromConfig builds a docker-run command on base of the given CommandDefinition.
// If a container with the same name already exists a docker-exec command will be created.
func (b *Builder) BuildCommandFromConfig(commandDef config.CommandDefinition) (*exec.Cmd, error) {
	dockerVersion, err := b.versionProvider.getAPIVersion()
	if err != nil {
		return nil, err
	}

	if b.containerExists(commandDef) {
		return NewExecBuilder(dockerVersion).BuildCommandFromConfig(commandDef)
	}

	return NewRunBuilder().BuildCommandFromConfig(commandDef)
}

func (b *Builder) containerExists(commandDef config.CommandDefinition) bool {
	if containerName, ok := commandDef.GetName(); ok {
		if b.containerExistenceChecker.exists(containerName) {
			return true
		}
	}

	return false
}

func buildArgumentsFromFunctions(
	commandDef config.CommandDefinition,
	builder builder.Builder,
	builders []argumentBuilderFunc,
) error {
	for _, argumentBuilderFunc := range builders {
		err := argumentBuilderFunc(commandDef, builder)
		if err != nil {
			return err
		}
	}

	return nil
}
