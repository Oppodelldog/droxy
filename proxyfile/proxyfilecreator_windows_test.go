package proxyfile

import (
	"github.com/sirupsen/logrus"
	"testing"
	"io/ioutil"
	"github.com/Oppodelldog/droxy/config"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestCreator_CreateProxyFiles(t *testing.T) {
	logrus.SetOutput(ioutil.Discard)

	fileCreatorMock := &mockFileCreationStrategy{}
	creator := New(fileCreatorMock)

	commandBinaryFilePathStub := "droxy-file-somewhere"
	commandNameStub := "some-command-name"

	cfg := &config.Configuration{
		Command: []config.CommandDefinition{
			{
				Name: &commandNameStub,
			},
		},
	}
	creator.CreateProxyFiles(commandBinaryFilePathStub, cfg, false)

	expectedCommandFilename := fmt.Sprintf("%s.exe", commandNameStub)

	assert.Equal(t, 1, fileCreatorMock.calls)
	assert.Equal(t, commandBinaryFilePathStub, fileCreatorMock.parmCommandBinaryFilePath)
	assert.Equal(t, expectedCommandFilename, fileCreatorMock.parmCommandNameFileName)
}
