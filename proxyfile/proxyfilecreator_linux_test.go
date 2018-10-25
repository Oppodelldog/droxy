package proxyfile

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/Oppodelldog/droxy/config"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

const testFolder = "/tmp/droxy/createProxyFilesTest"

func TestCreator_CreateProxyFiles(t *testing.T) {
	prepareTest(t)

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

	expectedCommandFilename := commandNameStub

	assert.Equal(t, 1, fileCreatorMock.calls)
	assert.Equal(t, commandBinaryFilePathStub, fileCreatorMock.parmCommandBinaryFilePath)
	assert.Equal(t, expectedCommandFilename, fileCreatorMock.parmCommandNameFileName)
}

func TestCreator_CreateProxyFiles_commandHasNoName_noFileWillBeCreated(t *testing.T) {
	prepareTest(t)

	fileCreatorMock := &mockFileCreationStrategy{}
	creator := New(fileCreatorMock)

	commandBinaryFilePathStub := "droxy-file-somewhere"
	cfg := &config.Configuration{
		Command: []config.CommandDefinition{
			{
				Name: nil,
			},
		},
	}
	creator.CreateProxyFiles(commandBinaryFilePathStub, cfg, false)

	assert.Equal(t, 0, fileCreatorMock.calls)
}

func TestCreator_CreateProxyFiles_commandIsTemplate_noFileWillBeCreated(t *testing.T) {
	prepareTest(t)

	fileCreatorMock := &mockFileCreationStrategy{}
	creator := New(fileCreatorMock)

	commandBinaryFilePathStub := "droxy-file-somewhere"
	commandStubName := "template"
	isTempalte := true
	cfg := &config.Configuration{
		Command: []config.CommandDefinition{
			{
				Name:       &commandStubName,
				IsTemplate: &isTempalte,
			},
		},
	}
	creator.CreateProxyFiles(commandBinaryFilePathStub, cfg, false)

	assert.Equal(t, 0, fileCreatorMock.calls)
}

func TestCreator_CreateProxyFiles_fileAlreadyExistsAndCreationIsNotForced_existingFileWillNotBeReplaced(t *testing.T) {
	prepareTest(t)

	logrus.SetOutput(ioutil.Discard)

	fileCreatorMock := &mockFileCreationStrategy{}
	creator := New(fileCreatorMock)

	commandNameStub := "some-command-name"
	fileThatShouldBeDeleted := commandNameStub
	err := ioutil.WriteFile(fileThatShouldBeDeleted, []byte("TEST"), 0666)
	if err != nil {
		t.Fatalf("Did not expect ioutil.WriteFile to return an error, but got: %v", err)
	}

	cfg := &config.Configuration{
		Command: []config.CommandDefinition{
			{
				Name: &commandNameStub,
			},
		},
	}
	creator.CreateProxyFiles("", cfg, false)

	_, err = os.Stat(fileThatShouldBeDeleted)
	assert.Nil(t, err, "Expect no error, since file should not have been deleted")
}

func TestCreator_CreateProxyFiles_fileAlreadyExistsAndCreationIsForced_existingFileWillBeReplaced(t *testing.T) {
	prepareTest(t)

	logrus.SetOutput(ioutil.Discard)

	fileCreatorMock := &mockFileCreationStrategy{}
	creator := New(fileCreatorMock)

	commandNameStub := "some-command-name"
	fileThatShouldBeDeleted := commandNameStub
	err := ioutil.WriteFile(fileThatShouldBeDeleted, []byte("TEST"), 0666)
	if err != nil {
		t.Fatalf("Did not expect ioutil.WriteFile to return an error, but got: %v", err)
	}

	cfg := &config.Configuration{
		Command: []config.CommandDefinition{
			{
				Name: &commandNameStub,
			},
		},
	}
	creator.CreateProxyFiles("", cfg, true)

	_, err = os.Stat(fileThatShouldBeDeleted)
	assert.Error(t, err, "Expect error, since file should be deleted")
}

func prepareTest(t *testing.T) {
	logrus.SetOutput(ioutil.Discard)

	err := os.RemoveAll(testFolder)
	if err != nil {
		t.Fatalf("Did not expect os.RemoveAll to return an error, but got: %v", err)
	}

	err = os.MkdirAll(testFolder, 0776)
	if err != nil {
		t.Fatalf("Did not expect os.MkdirAll to return an error, but got: %v", err)
	}
}
