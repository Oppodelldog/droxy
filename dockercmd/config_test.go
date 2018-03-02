package dockercmd

import (
	"strings"
	"testing"

	"github.com/Oppodelldog/droxy/config"

	"os"

	"github.com/stretchr/testify/assert"
)

func TestBuildCommandFromConfig(t *testing.T) {
	commandName := "some-command"
	confgiuration := getFullFeatureConfig(commandName)

	cmd, err := BuildCommandFromConfig(commandName, confgiuration)
	if err != nil {
		t.Fatalf("Did not expect BuildCommandFromConfig to return an error, but got: %v", err)
	}

	expectedArgsFromTestCall := strings.Join(os.Args[1:], " ")
	commandString := strings.Join(cmd.Args, " ")
	expectedCommandString := "docker run -i --rm -p 8080:9080 -p 8081:9081 -p 8080:9080 -p 8081:9081 -v /home/nils:/home/nils -v /run/user/1000/keyring/ssh:/run/ssh.sock -v /etc/passwd:/etc/passwd:ro -v /etc/group:/etc/group:ro -v /run/docker.sock:/run/docker.sock -e HOME:/home/nils -e SSH_AUTH_SOCK:/run/ssh.sock -e DOCKER_HOST=unix:///run/docker.sock --group-add 1000 --group-add 4 --group-add 5 --group-add 20 --group-add 24 --group-add 27 --group-add 30 --group-add 46 --group-add 113 --group-add 128 --group-add 130 -u 1000:1000 -a STDIN -a STDOUT -a STDERR --network some-docker-network some-image:v1.02 some-entrypoint-cmd additionalArgument=123"
	expectedCommandString = strings.Join([]string{expectedCommandString, expectedArgsFromTestCall}, " ")

	assert.Equal(t, expectedCommandString, commandString)
}

func getFullFeatureConfig(commandName string) *config.Configuration {

	fullFeatureTemplate := getFullFeatureTemplateDef()
	fullFeatureCommand := getFullFeatureDef(commandName)
	return &config.Configuration{
		Command: []config.CommandDefinition{
			fullFeatureTemplate,
			fullFeatureCommand,
		},
	}
}

func getFullFeatureTemplateDef() config.CommandDefinition {
	isTemplate := true
	entryPoint := "some-entrypoint-cmd"
	name := "some template"
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

	replaceArgs := [][]string{
		{
			"-dxdebug.remote_host=127.0.0.1",
			"-dxdebug.remote_host=172.17.0.1",
		},
	}

	additionalArgs := []string{
		"additionalArgument=123",
	}

	return config.CommandDefinition{
		IsTemplate:      &isTemplate,
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
		ReplaceArgs:     &replaceArgs,
		AdditionalArgs:  &additionalArgs,
	}
}
func getFullFeatureDef(commandName string) config.CommandDefinition {
	isTemplate := true
	template := "some template"
	entryPoint := "some-entrypoint-cmd"
	name := commandName
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

	replaceArgs := [][]string{
		{
			"-dxdebug.remote_host=127.0.0.1",
			"-dxdebug.remote_host=172.17.0.1",
		},
	}

	additionalArgs := []string{
		"additionalArgument=123",
	}

	return config.CommandDefinition{
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
		ReplaceArgs:     &replaceArgs,
		AdditionalArgs:  &additionalArgs,
	}
}
