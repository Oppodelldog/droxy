package config

import (
	"os"
	"path"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/assert"
)

const testFolder = "/tmp/droxy/test/config/load"

func TestNewLoader(t *testing.T) {
	assert.IsType(t, new(configLoader), NewLoader())
}

func TestConfigLoader_Load(t *testing.T) {
	testFilePath := createTestFile(t)

	os.Setenv("DROXY_CONFIG", testFilePath)

	loader := NewLoader()
	cfg := loader.Load()

	assert.Equal(t, "0815", cfg.Version)

	os.Unsetenv("DROXY_CONFIG")

	cleanupTestFile(t)
}

func cleanupTestFile(t *testing.T) {
	err := os.RemoveAll(testFolder)
	if err != nil {
		t.Fatalf("Did not expect os.RemoveAll to return an error, but got: %v", err)
	}
}

func createTestFile(t *testing.T) string {

	testFile := "droxy.toml"

	err := os.RemoveAll(testFolder)
	if err != nil {
		t.Fatalf("Did not expect os.RemoveAll to return an error, but got: %v", err)
	}
	err = os.MkdirAll(testFolder, 0777)
	if err != nil {
		t.Fatalf("Did not expect os.MkDirAll to return an error, but got: %v", err)
	}

	testFilePath := path.Join(testFolder, testFile)

	tempFile, err := os.OpenFile(testFilePath, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		t.Fatalf("Did not expect os.OpenFile to return an error, but got: %v", err)
	}
	defer tempFile.Close()

	cfg := Configuration{Version: "0815"}
	tomlEncoder := toml.NewEncoder(tempFile)
	err = tomlEncoder.Encode(cfg)
	if err != nil {
		t.Fatalf("Did not expect tomlEncoder.Encode to return an error, but got: %v", err)
	}

	return testFilePath
}
