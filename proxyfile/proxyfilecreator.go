package proxyfile

import (
	"os"

	"github.com/Oppodelldog/droxy/config"
	"github.com/sirupsen/logrus"
)

// New creates a new proxy file creator
func New(creationStrategy FileCreationStrategy, configLoader config.Loader) *Creator {
	return &Creator{
		creationStrategy:          creationStrategy,
		configLoader:              configLoader,
		getExecutableFilePathFunc: getExecutableFilePath,
	}
}

// Creator creates commands
type Creator struct {
	creationStrategy          FileCreationStrategy
	configLoader              config.Loader
	getExecutableFilePathFunc getExecutableFilePathFuncDef
}

type getExecutableFilePathFuncDef func() (string, error)

// CreateProxyFiles creates droxy commands
func (pfc *Creator) CreateProxyFiles(isForced bool) error {
	cfg := pfc.configLoader.Load()

	commandBinaryFilePath, err := pfc.getExecutableFilePathFunc()
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	for _, command := range cfg.Command {

		if !command.HasName() {
			logrus.Warnf("skipped command because name is missing!")
			continue
		}

		if isTemplate, ok := command.GetIsTemplate(); isTemplate && ok {
			continue
		}

		if commandName, ok := command.GetName(); ok {

			commandNameFileName := GetCommandNameFilename(commandName)
			if fileInfo, err := os.Stat(commandNameFileName); err == nil {
				if fileInfo.IsDir() {
					logrus.Warnf("droxy command file already exists as a directory '%s'", commandNameFileName)
					return nil
				}
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
