package proxyfile

import (
	"os"
	"github.com/sirupsen/logrus"
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/helper"
)

// FileCreationStrategy defines the interface for creation of a droxy commands in filesystem
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

// CreateProxyFiles creates droxy commands
func (pfc *Creator) CreateProxyFiles(commandBinaryFilePath string, configuration *config.Configuration, isForced bool) error {
	for _, command := range configuration.Command {

		if !command.HasName() {
			logrus.Warnf("skipped command because name is missing!")
			continue
		}

		if isTemplate, ok := command.GetIsTemplate(); isTemplate && ok {
			continue
		}

		if commandName, ok := command.GetName(); ok {

			commandNameFileName := helper.GetCommandNameFilename(commandName)
			if _, err := os.Stat(commandNameFileName); err == nil {
				if isForced {
					err := os.Remove(commandNameFileName)
					if err != nil {
						panic(err)
					}
				} else {
					logrus.Warnf("droxy command file (%s) already exists for command '%s'", commandNameFileName, commandName)
					continue
				}
			}

			err := pfc.creationStrategy.CreateProxyFile(commandBinaryFilePath, commandNameFileName)
			if err != nil {
				logrus.Errorf("error creating symlink '%s': %v", commandName, err)
				continue
			}

			logrus.Infof("created '%s'", commandName)
		}

	}

	return nil
}
