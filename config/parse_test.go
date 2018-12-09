package config

import (
	"testing"

	"os"
	"path"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/assert"
)

func TestParseFromBytes_fullFeatureConfig(t *testing.T) {
	configBytes := getFullFeatureConfigFixture()

	configuration, err := parseFromBytes(configBytes)
	if err != nil {
		t.Fatalf("Did expect parseFromBytes to return no error, but got: %v", err)
	}

	var expectedConfiguration Configuration
	expectedConfiguration.Version = "1"

	expectedConfiguration.Command = []CommandDefinition{getFullFeatureCommandDefinition()}

	assert.Equal(t, &expectedConfiguration, configuration)
}

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

func getFullFeatureConfigFixture() []byte {
	return []byte(`
	Version="1"	

    [[command]]
      requireEnvVars=true
      name = "some command"  # name of the command
      isTemplate = true       # this command can be used as a template, no command will be created
      addGroups = true        # add current systems groups
      impersonate = true      # use executing user and group for execution in the container
      workDir = "someDir/"        # define working directory
      removeContainer=true    # remove container after command has finished
      isInteractive=true      # enable interaction with the called command
	  isDetached=true
      isDaemon=true			  # deprecated
      uniqueNames=true
      network="some-docker-network"
      image="some-image:v1.02"
	  entryPoint="some-entryPoint"
      command="some-cmd"
      template="some template"
      envFile=".env"	
      ip="127.1.2.3"

      # volume mappings
      volumes = [
          "${HOME}:${HOME}",
          "${SSH_AUTH_SOCK}:/run/ssh.sock",
          "/etc/passwd:/etc/passwd:ro",
          "/etc/group:/etc/group:ro",
          "/run/docker.sock:/run/docker.sock"
      ]

      # environment variable mappings
      envvars = [
          "HOME:${HOME}",
          "SSH_AUTH_SOCK:/run/ssh.sock",
          "DOCKER_HOST=unix:///run/docker.sock"
      ]

      links = [
		"${LINK_ENV_VAR}:${LINK_ENV_VAR}",
        "containerXY:aliasXY"
      ]

      ports = [
          "8080:9080",
	      "8081:9081",
      ]

	  portsFromParams = [
	      "some regex where the group (\\d*) parses the port from",
	  ]

      replaceArgs = [
      	[
			"-dxdebug.remote_host=127.0.0.1",
			"-dxdebug.remote_host=172.17.0.1"		
	    ]
      ]

	  additionalArgs = ["additionalArgument=123"]
`)
}

func TestParseFromBytes_emptyConfig(t *testing.T) {
	configBytes := getEmptyConfig()

	configuration, err := parseFromBytes(configBytes)
	if err != nil {
		t.Fatalf("Did expect parseFromBytes to return no error, but got: %v", err)
	}

	var expectedConfiguration Configuration

	assert.Equal(t, &expectedConfiguration, configuration)
}

func getEmptyConfig() []byte {
	return []byte{}
}

func TestParse(t *testing.T) {
	const testFolder = "/tmp/droxy/test/config/parse"
	const testFile = "testFile.toml"
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
	cfg := Configuration{
		Version: "4711",
		Command: []CommandDefinition{
			command,
		},
	}

	testFilePath := path.Join(testFolder, testFile)
	tempFile, err := os.OpenFile(testFilePath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		t.Fatalf("Did not expect os.OpenFile to return an error, but got: %v", err)
	}

	tomlEncoder := toml.NewEncoder(tempFile)
	err = tomlEncoder.Encode(cfg)
	if err != nil {
		t.Fatalf("Did not expect tomlEncoder.Encode to return an error, but got: %v", err)
	}
	tempFile.Close()

	parsedCfg, err := Parse(testFilePath)
	if err != nil {
		t.Fatalf("Did not expect Parse to return an error, but got: %v", err)
	}

	assert.Equal(t, &cfg, parsedCfg)

	err = os.RemoveAll(testFolder)
	if err != nil {
		t.Fatalf("Did not expect os.RemoveAll to return an error, but got: %v", err)
	}
}

func TestParse_FileNotExists_Error(t *testing.T) {
	_, err := Parse("/tmp/tdroxy/his-does-not-exist.toml")
	assert.Error(t, err)
}

func TestParseBytes_InvalidInput_ExpectError(t *testing.T) {
	_, err := parseFromBytes([]byte("SBSGUOPGBSUOsg"))
	assert.Error(t, err)
}
