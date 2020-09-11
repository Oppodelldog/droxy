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
		newExec                   newCommandBuilderFunc
		newRun                    newCommandBuilderFunc
	}

	argumentBuilderFunc func(commandDef config.CommandDefinition, builder builder.Builder) error
)

type (
	CommandBuilder interface {
		BuildCommandFromConfig(commandDef config.CommandDefinition) (*exec.Cmd, error)
	}
	newCommandBuilderFunc func(string) CommandBuilder
)

// NewBuilder returns a new *Builder.
func NewBuilder() (*Builder, error) {
	clientAdapter, err := newDockerClientAdapter()
	if err != nil {
		return nil, err
	}

	return &Builder{
		containerExistenceChecker: clientAdapter,
		versionProvider:           clientAdapter,
		newExec:                   func(dockerVersion string) CommandBuilder { return NewExecBuilder(dockerVersion) },
		newRun:                    func(dockerVersion string) CommandBuilder { return NewRunBuilder(dockerVersion) },
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
		return b.newExec(dockerVersion).BuildCommandFromConfig(commandDef)
	}

	return b.newRun(dockerVersion).BuildCommandFromConfig(commandDef)
}

func (b *Builder) containerExists(commandDef config.CommandDefinition) bool {
	if containerName, ok := commandDef.GetName(); ok {
		if b.containerExistenceChecker.exists(containerName) {
			return true
		}
	}

	return false
}
