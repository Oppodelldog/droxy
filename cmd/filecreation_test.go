package cmd

import (
	"github.com/Oppodelldog/droxy/config"
	"testing"
)

type fileCreationStrategyStub struct {
	result error
}

func (fc *fileCreationStrategyStub) CreateProxyFile(string, string) error {
	return fc.result
}

type configLoaderStub struct {
	result error
}

func (cl *configLoaderStub) Load() *config.Configuration {
	return &config.Configuration{}
}

func Test_fileCreationSubCommandWrapper_createCommand(t *testing.T) {
	fileCreationStrategyStub := &fileCreationStrategyStub{}
	configLoaderStub := &configLoaderStub{}

	filecreator := newFileCreationSubCommand("symlinks", fileCreationStrategyStub, configLoaderStub)

	filecreator.cobraCommand.Run(nil, nil)
}
