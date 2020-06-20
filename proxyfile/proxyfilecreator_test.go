package proxyfile

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/Oppodelldog/droxy/logger"

	"github.com/Oppodelldog/droxy/config"

	"github.com/stretchr/testify/assert"
)

const testFolder = "/tmp/droxy/createProxyFilesTest"

func TestCreator_New(t *testing.T) {
	fileCreatorMock := &mockFileCreationStrategy{}
	configLoaderMock := &configLoaderMock{}
	creator := New(fileCreatorMock, configLoaderMock)

	assert.IsType(t, Creator{}, creator)
	assert.Exactly(t, fileCreatorMock, creator.creationStrategy)
	assert.Exactly(t, configLoaderMock, creator.configLoader)

	if reflect.ValueOf(creator.getExecutableFilePathFunc).Pointer() != reflect.ValueOf(getExecutableFilePath).Pointer() {
		t.Fatal("expected 'getExecutableFilePath' to be configured as getExecutableFilePathFunc, but was not")
	}
}

func getTestConfig() *config.Configuration {
	commandNameStub := "some-command-name"

	return &config.Configuration{
		Command: []config.CommandDefinition{
			{Name: &commandNameStub},
		},
	}
}

func getTestConfigWithEmptyCommand() *config.Configuration {
	return &config.Configuration{Command: []config.CommandDefinition{{}}}
}

func TestCreator_CreateProxyFiles(t *testing.T) {
	defer prepareTest(t)()

	commandBinaryFilePathStub := "/tmp/droxy"

	fileCreatorMock := &mockFileCreationStrategy{}
	configLoaderMock := &configLoaderMock{stubbedConfig: getTestConfig()}
	creator := &Creator{
		creationStrategy:          fileCreatorMock,
		configLoader:              configLoaderMock,
		getExecutableFilePathFunc: func() (string, error) { return commandBinaryFilePathStub, nil },
	}

	err := creator.CreateProxyFiles(false)
	if err != nil {
		t.Fatalf("Did not expect CreateProxyFiles to return an error, but got: %v", err)
	}

	expectedCommandFilename := ensureOsSpecificBinaryFilename(*configLoaderMock.stubbedConfig.Command[0].Name)

	assert.Equal(t, 1, fileCreatorMock.calls)
	assert.Equal(t, commandBinaryFilePathStub, fileCreatorMock.parmCommandBinaryFilePath)
	assert.Equal(t, expectedCommandFilename, fileCreatorMock.parmCommandNameFileName)
}

func TestCreator_CreateProxyFiles_commandHasNoName_noFileWillBeCreated(t *testing.T) {
	defer prepareTest(t)()

	fileCreatorMock := &mockFileCreationStrategy{}
	configLoaderMock := &configLoaderMock{stubbedConfig: getTestConfigWithEmptyCommand()}
	creator := &Creator{
		creationStrategy:          fileCreatorMock,
		configLoader:              configLoaderMock,
		getExecutableFilePathFunc: func() (string, error) { return "", nil },
	}

	err := creator.CreateProxyFiles(false)
	if err != nil {
		t.Fatalf("Did not expect CreateProxyFiles to return an error, but got: %v", err)
	}

	assert.Equal(t, 0, fileCreatorMock.calls)
}

func TestCreator_CreateProxyFiles_commandIsTemplate_noFileWillBeCreated(t *testing.T) {
	defer prepareTest(t)()

	fileCreatorMock := &mockFileCreationStrategy{}
	testConfig := getTestConfig()
	isTemplate := true
	testConfig.Command[0].IsTemplate = &isTemplate
	configLoaderMock := &configLoaderMock{stubbedConfig: testConfig}
	creator := &Creator{
		creationStrategy:          fileCreatorMock,
		configLoader:              configLoaderMock,
		getExecutableFilePathFunc: func() (string, error) { return "", nil },
	}

	err := creator.CreateProxyFiles(false)
	if err != nil {
		t.Fatalf("Did not expect CreateProxyFiles to return an error, but got: %v", err)
	}

	assert.Equal(t, 0, fileCreatorMock.calls)
}

func TestCreator_CreateProxyFiles_fileAlreadyExistsAndCreationIsNotForced_existingFileWillNotBeReplaced(t *testing.T) {
	defer prepareTest(t)()

	logger.SetOutput(ioutil.Discard)

	fileCreatorMock := &mockFileCreationStrategy{}
	configLoaderMock := &configLoaderMock{stubbedConfig: getTestConfig()}
	creator := &Creator{
		creationStrategy:          fileCreatorMock,
		configLoader:              configLoaderMock,
		getExecutableFilePathFunc: func() (string, error) { return "", nil },
	}

	commandNameStub := *configLoaderMock.stubbedConfig.Command[0].Name
	fileThatShouldNotBeDeleted := commandNameStub

	err := ioutil.WriteFile(fileThatShouldNotBeDeleted, []byte("TEST"), writePerm)
	if err != nil {
		t.Fatalf("Did not expect ioutil.WriteFile to return an error, but got: %v", err)
	}

	err = creator.CreateProxyFiles(false)
	if err != nil {
		t.Fatalf("Did not expect CreateProxyFiles to return an error, but got: %v", err)
	}

	_, errStat := os.Stat(fileThatShouldNotBeDeleted)
	assert.Nil(t, errStat, "Expect no error, since file should not have been deleted")

	err = os.Remove(fileThatShouldNotBeDeleted)
	if err != nil {
		t.Fatalf("Did not expect os.Remove to return an error, but got: %v", err)
	}
}

func TestCreator_CreateProxyFiles_fileAlreadyExistsAsDirectoryAndCreationIsForced_folderWillNotBeDeleted(t *testing.T) {
	defer prepareTest(t)()

	logger.SetOutput(ioutil.Discard)

	fileCreatorMock := &mockFileCreationStrategy{}
	configLoaderMock := &configLoaderMock{stubbedConfig: getTestConfig()}
	creator := &Creator{
		creationStrategy:          fileCreatorMock,
		configLoader:              configLoaderMock,
		getExecutableFilePathFunc: func() (string, error) { return "", nil },
	}

	commandNameStub := *configLoaderMock.stubbedConfig.Command[0].Name
	folderThatShallNotBeDeleted := commandNameStub

	err := os.MkdirAll(folderThatShallNotBeDeleted, 0666)
	if err != nil {
		t.Fatalf("Did not expect os.MkdirAll to return an error, but got: %v", err)
	}

	err = creator.CreateProxyFiles(true)
	if err != nil {
		t.Fatalf("Did not expect CreateProxyFiles to return an error, but got: %v", err)
	}

	_, err = os.Stat(folderThatShallNotBeDeleted)
	assert.NoError(t, err, "Expect no error, since folder should not be deleted, but got: %v", err)
}

func TestCreator_CreateProxyFiles_fileAlreadyExistsAndCreationIsForced_existingFileWillBeReplaced(t *testing.T) {
	defer prepareTest(t)()

	logger.SetOutput(ioutil.Discard)

	fileCreatorMock := &mockFileCreationStrategy{}
	configLoaderMock := &configLoaderMock{stubbedConfig: getTestConfig()}
	creator := &Creator{
		creationStrategy:          fileCreatorMock,
		configLoader:              configLoaderMock,
		getExecutableFilePathFunc: func() (string, error) { return "", nil },
	}

	commandNameStub := *configLoaderMock.stubbedConfig.Command[0].Name
	fileThatShouldBeDeleted := ensureOsSpecificBinaryFilename(commandNameStub)

	err := ioutil.WriteFile(fileThatShouldBeDeleted, []byte("TEST"), writePerm)
	if err != nil {
		t.Fatalf("Did not expect ioutil.WriteFile to return an error, but got: %v", err)
	}

	err = creator.CreateProxyFiles(true)
	if err != nil {
		t.Fatalf("Did not expect CreateProxyFiles to return an error, but got: %v", err)
	}

	_, err = os.Stat(fileThatShouldBeDeleted)
	assert.Error(t, err, "Expect error, since file should be deleted")
}

func ensureOsSpecificBinaryFilename(filePath string) string {
	if runtime.GOOS == "windows" {
		filePath += ".exe"
	}

	return filePath
}

func prepareTest(t *testing.T) func() {
	logger.SetOutput(ioutil.Discard)

	err := os.RemoveAll(testFolder)
	if err != nil {
		t.Fatalf("Did not expect os.RemoveAll to return an error, but got: %v", err)
	}

	err = os.MkdirAll(testFolder, 0776)
	if err != nil {
		t.Fatalf("Did not expect os.MkdirAll to return an error, but got: %v", err)
	}

	originalDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	err = os.Chdir(testFolder)
	if err != nil {
		t.Fatalf("Did not expect os.Chdir to return an error, but got: %v", err)
	}

	return func() {
		err = os.Chdir(originalDir)
		if err != nil {
			t.Fatalf("Did not expect os.Chdir to return an error when switching back to original dir, but got: %v", err)
		}

		err := os.RemoveAll(testFolder)
		if err != nil {
			t.Fatalf("Did not expect os.RemoveAll to return an error, but got: %v", err)
		}
	}
}
