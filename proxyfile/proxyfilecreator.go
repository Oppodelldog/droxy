package proxyfile

import (
	"docker-proxy-command/config"
	"os"

	"github.com/sirupsen/logrus"
)

// FileCreationStrategy defines the interface for creation of a docker-proxy commands in filesystem
type FileCreationStrategy interface {
	CreateProxyFile(string, string) error
}

// New creates a new proxy file creator
func New(creationStrategy FileCreationStrategy) *Creator {
	return &Creator{
		creationStrategy: creationStrategy,
	}
}

// Creator creates commands
type Creator struct {
	creationStrategy FileCreationStrategy
}

// CreateProxyFiles creates docker-proxy commands
func (pfc *Creator) CreateProxyFiles(commandBinaryFilePath string, configuration *config.Configuration, isForced bool) error {
	for _, command := range configuration.Command {

		if !command.HasName() {
			logrus.Warnf("skipped command because name is missing!")
			continue
		}

		if isTemplate, ok := command.GetIsTemplate(); isTemplate && ok {
			continue
		}

		if _, err := os.Stat(*command.Name); err == nil {
			if isForced {
				err := os.Remove(*command.Name)
				if err != nil {
					panic(err)
				}
			} else {
				logrus.Warnf("command symlink already exists for '%s'", *command.Name)
				continue
			}
		}

		err := pfc.creationStrategy.CreateProxyFile(commandBinaryFilePath, *command.Name)
		if err != nil {
			logrus.Errorf("error creating symlink '%s': %v", *command.Name, err)
			continue
		}

		logrus.Infof("created '%s'", *command.Name)
	}

	return nil
}
