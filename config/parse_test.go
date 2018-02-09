package config

import (
	"testing"

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
	template := "some template"
	entryPoint := "some-entrypoint-cmd"
	name := "some command"
	image := "some-image:v1.02"
	network := "some-docker-network"
	isInteractive := true
	addGroups := true
	impersonate := true
	removeContainer := true
	workDir := "someDir/"
	volumes := []string{
		"${HOME}:${HOME}",
		"${SSH_AUTH_SOCK}:/run/ssh.sock",
		"/etc/passwd:/etc/passwd:ro",
		"/etc/group:/etc/group:ro",
		"/run/docker.sock:/run/docker.sock",
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

	return CommandDefinition{
		IsTemplate:      &isTemplate,
		Template:        &template,
		EntryPoint:      &entryPoint,
		Name:            &name,
		Image:           &image,
		Network:         &network,
		IsInteractive:   &isInteractive,
		AddGroups:       &addGroups,
		Impersonate:     &impersonate,
		RemoveContainer: &removeContainer,
		WorkDir:         &workDir,
		Volumes:         &volumes,
		EnvVars:         &envVars,
		Ports:           &ports,
	}

}

func getFullFeatureConfigFixture() []byte {
	return []byte(`
	Version="1"	

    [[command]]
      name = "some command"  # name of the command
      isTemplate = true       # this command can be used as a template, no command will be created
      addGroups = true        # add current systems groups
      impersonate = true      # use executing user and group for execution in the container
      workDir = "someDir/"        # define working directory
      removeContainer=true    # remove container after command has finished
      isInteractive=true      # enable interaction with the called command
      network="some-docker-network"
      image="some-image:v1.02"
      entryPoint="some-entrypoint-cmd"
      template="some template"

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

      ports = [
          "8080:9080",
	      "8081:9081",
      ]
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
