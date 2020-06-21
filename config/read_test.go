package config

import (
	"io/ioutil"
	"testing"

	"os"
	"path"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/assert"
)

func TestParseFromBytes_fullFeatureConfig(t *testing.T) {
	configBytes := getFullFeatureConfigFixture(t)

	configuration, err := parseFromBytes(configBytes)
	if err != nil {
		t.Fatalf("Did expect parseFromBytes to return no error, but got: %v", err)
	}

	var expectedConfiguration Configuration
	expectedConfiguration.Version = "1"

	expectedConfiguration.Command = []CommandDefinition{getFullFeatureCommandDefinition()}

	assert.Equal(t, expectedConfiguration, configuration)
}

//nolint:funlen
func getFullFeatureCommandDefinition() CommandDefinition {
	isTemplate := true
	isDetached := true
	requireEnvVars := true
	template := "some template"
	entryPoint := "some-entryPoint"
	command := "some-cmd"
	name := "some command"
	image := "some-image:v1.02"
	network := "some-docker-network"
	isInteractive := true
	addGroups := true
	impersonate := true
	removeContainer := true
	uniqueNames := true
	workDir := "someDir/"
	envFile := ".env"
	ip := "127.1.2.3"
	volumes := []string{
		"${HOME}:${HOME}",
		"${SSH_AUTH_SOCK}:/run/ssh.sock",
		"/etc/passwd:/etc/passwd:ro",
		"/etc/group:/etc/group:ro",
		"/run/docker.sock:/run/docker.sock",
	}
	links := []string{
		"${LINK_ENV_VAR}:${LINK_ENV_VAR}",
		"containerXY:aliasXY",
	}
	envVars := []string{
		"HOME:${HOME}",
		"SSH_AUTH_SOCK:/run/ssh.sock",
		"DOCKER_HOST=unix:///run/docker.sock",
	}
	ports := []string{
		"8080:9080",
		"8081:9081",
	}
	portsFromParams := []string{
		"some regex where the group (\\d*) parses the port from",
	}

	replaceArgs := [][]string{
		{
			"-dxdebug.remote_host=127.0.0.1",
			"-dxdebug.remote_host=172.17.0.1",
		},
	}

	additionalArgs := []string{
		"additionalArgument=123",
	}

	return CommandDefinition{
		RequireEnvVars:  &requireEnvVars,
		IsTemplate:      &isTemplate,
		IsDetached:      &isDetached,
		IsDaemon:        &isDetached,
		Template:        &template,
		EntryPoint:      &entryPoint,
		Command:         &command,
		Name:            &name,
		UniqueNames:     &uniqueNames,
		Image:           &image,
		Network:         &network,
		IsInteractive:   &isInteractive,
		EnvFile:         &envFile,
		IP:              &ip,
		AddGroups:       &addGroups,
		Impersonate:     &impersonate,
		RemoveContainer: &removeContainer,
		WorkDir:         &workDir,
		Volumes:         &volumes,
		Links:           &links,
		EnvVars:         &envVars,
		Ports:           &ports,
		PortsFromParams: &portsFromParams,
		ReplaceArgs:     &replaceArgs,
		AdditionalArgs:  &additionalArgs,
	}
}

func getFullFeatureConfigFixture(t *testing.T) []byte {
	fullFeatureConfig := path.Join(getProjectDir(), "config/test/fullFeature.toml")

	b, err := ioutil.ReadFile(fullFeatureConfig)
	if err != nil {
		t.Fatalf("did not expect ReadFile to return an error, but got: %v", err)
	}

	return b
}

func TestParseFromBytes_emptyConfig(t *testing.T) {
	configBytes := getEmptyConfig()

	configuration, err := parseFromBytes(configBytes)
	if err != nil {
		t.Fatalf("Did expect parseFromBytes to return no error, but got: %v", err)
	}

	var expectedConfiguration Configuration

	assert.Equal(t, expectedConfiguration, configuration)
}

func getEmptyConfig() []byte {
	return []byte{}
}

func TestParse(t *testing.T) {
	const (
		testFolder = "/tmp/droxy/test/config/readFromFile"
		testFile   = "testFile.toml"
	)

	err := os.RemoveAll(testFolder)
	if err != nil {
		t.Fatalf("Did not expect os.RemoveAll to return an error, but got: %v", err)
	}

	err = os.MkdirAll(testFolder, 0777)
	if err != nil {
		t.Fatalf("Did not expect os.MkDirAll to return an error, but got: %v", err)
	}

	commandName := "test-command"
	command := CommandDefinition{
		Name: &commandName,
	}
	testFilePath := path.Join(testFolder, testFile)

	cfg := Configuration{
		Version: "4711",
		Command: []CommandDefinition{
			command,
		},
		ConfigFilePath: testFilePath,
	}

	tempFile, err := os.OpenFile(testFilePath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		t.Fatalf("Did not expect os.OpenFile to return an error, but got: %v", err)
	}

	tomlEncoder := toml.NewEncoder(tempFile)

	err = tomlEncoder.Encode(cfg)
	if err != nil {
		t.Fatalf("Did not expect tomlEncoder.Encode to return an error, but got: %v", err)
	}

	err = tempFile.Close()
	if err != nil {
		t.Fatalf("Did not expect tempFile.Close to return an error, but got: %v", err)
	}

	parsedCfg, err := readFromFile(testFilePath)
	if err != nil {
		t.Fatalf("Did not expect readFromFile to return an error, but got: %v", err)
	}

	assert.Equal(t, cfg, parsedCfg)

	err = os.RemoveAll(testFolder)
	if err != nil {
		t.Fatalf("Did not expect os.RemoveAll to return an error, but got: %v", err)
	}
}

func TestParse_FileNotExists_Error(t *testing.T) {
	_, err := readFromFile("/tmp/droxy/this-does-not-exist.toml")
	assert.Error(t, err)
}

func TestParseBytes_InvalidInput_ExpectError(t *testing.T) {
	_, err := parseFromBytes([]byte("SBSGUOPGBSUOsg"))
	assert.Error(t, err)
}

func getProjectDir() string {
	gp := os.Getenv("GOPATH")
	return path.Join(gp, "src/github.com/Oppodelldog/droxy")
}
