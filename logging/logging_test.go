package logging

import (
	"io"
	"os"
	"path"
	"testing"

	"github.com/Oppodelldog/droxy/config"
	"github.com/stretchr/testify/assert"
)

const testFolder = "/tmp/droxy-tests"

func TestGetLogWriter_returnsAWriterCloser(t *testing.T) {

	prepareTestFolder(t)
	configFilePath := path.Join(testFolder, "droxy.toml")

	cfg := &config.Configuration{}
	cfg.SetConfigurationFilePath(configFilePath)

	writer, err := GetLogWriter(cfg)
	if err != nil {
		t.Fatalf("Did not expect GetLogWriter to return an error, but got: %v", err)
	}
	err = writer.Close()
	if err != nil {
		t.Fatalf("Did not expect writer.Close() to return an error, but got: %v", err)
	}

	assert.Implements(t, new(io.WriteCloser), writer)
}

func TestGetLogWriter_createsFileNextToConfig(t *testing.T) {

	prepareTestFolder(t)
	configFilePath := path.Join(testFolder, "droxy.toml")

	cfg := &config.Configuration{}
	cfg.SetConfigurationFilePath(configFilePath)

	writer, err := GetLogWriter(cfg)
	if err != nil {
		t.Fatalf("Did not expect GetLogWriter to return an error, but got: %v", err)
	}
	err = writer.Close()
	if err != nil {
		t.Fatalf("Did not expect writer.Close() to return an error, but got: %v", err)
	}

	assertLogfile(t, path.Join(testFolder, "droxy.log"))
}

func TestGetLogWriter_returnsErrorOnFileCreation(t *testing.T) {

	configFilePath := "/invalid-apth/this-will-create-an-error"

	cfg := &config.Configuration{}
	cfg.SetConfigurationFilePath(configFilePath)

	_, err := GetLogWriter(cfg)

	assert.Error(t, err)
}

func assertLogfile(t *testing.T, logfilePath string) {
	fileFound := true
	if _, err := os.Stat(logfilePath); os.IsNotExist(err) {
		fileFound = false
	}

	assert.True(t, fileFound)
}

func prepareTestFolder(t *testing.T) {

	err := os.RemoveAll(testFolder)
	if err != nil {
		t.Fatalf("Did not expect os.RemoveAll to return an error, but got: %v", err)
	}

	err = os.MkdirAll(testFolder, 0755)
	if err != nil {
		t.Fatalf("Did not expect os.MkdirAll to return an error, but got: %v", err)
	}
}
