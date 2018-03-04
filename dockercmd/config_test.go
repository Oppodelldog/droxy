package dockercmd

import (
	"strings"
	"testing"

	"github.com/Oppodelldog/droxy/config"

	"os"

	"fmt"
	"github.com/pmezard/go-difflib/difflib"
	"github.com/stretchr/testify/assert"
)

func TestBuildCommandFromConfig(t *testing.T) {

	os.Setenv("VOLUME_ENV_VAR", "volEnvVarStub")
	os.Setenv("ENV_VAR", "envVarStub")

	commandName := "some-command"
	confgiuration := getFullFeatureConfig(commandName)

	cmd, err := BuildCommandFromConfig(commandName, confgiuration)
	if err != nil {
		t.Fatalf("Did not expect BuildCommandFromConfig to return an error, but got: %v", err)
	}

	expectedArgsFromTestCall := strings.Join(os.Args[1:], " ")
	commandString := strings.Join(cmd.Args, " ")

	expectedCommandStrings := []string{
		strings.TrimSpace(strings.Join([]string{"docker run -i --rm -p 8080:9080 -p 8081:9081 -p 8080:9080 -p 8081:9081 -v volEnvVarStub:volEnvVarStub -v /etc/passwd:/etc/passwd:ro -v /etc/group:/etc/group:ro -v /run/docker.sock:/run/docker.sock -e HOME:envVarStub -e SSH_AUTH_SOCK:/run/ssh.sock -e DOCKER_HOST=unix:///run/docker.sock -a STDIN -a STDOUT -a STDERR --network some-docker-network some-image:v1.02 some-entrypoint-cmd additionalArgument=123", expectedArgsFromTestCall}, " ")),
		strings.TrimSpace(strings.Join([]string{"docker run -t -i --rm -p 8080:9080 -p 8081:9081 -p 8080:9080 -p 8081:9081 -v volEnvVarStub:volEnvVarStub -v /etc/passwd:/etc/passwd:ro -v /etc/group:/etc/group:ro -v /run/docker.sock:/run/docker.sock -e HOME:envVarStub -e SSH_AUTH_SOCK:/run/ssh.sock -e DOCKER_HOST=unix:///run/docker.sock -a STDIN -a STDOUT -a STDERR --network some-docker-network some-image:v1.02 some-entrypoint-cmd additionalArgument=123", expectedArgsFromTestCall}, " ")),
	}

	for _, expected := range expectedCommandStrings {
		diff := difflib.UnifiedDiff{
			A:        difflib.SplitLines(expected),
			B:        difflib.SplitLines(commandString),
			FromFile: "Original",
			ToFile:   "Current",
			Context:  3,
		}
		text, _ := difflib.GetUnifiedDiffString(diff)
		fmt.Printf(text)
		fmt.Println("------------------------------------------------")
	}

	assert.Contains(t, expectedCommandStrings, commandString)

	os.Unsetenv("VOLUME_ENV_VAR")
	os.Unsetenv("ENV_VAR")
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
	addGroups := false   // disabled because of different values on build than on local...
	impersonate := false // disabled because of different values on build than on local...
	removeContainer := true
	workDir := "someDir/"
	volumes := []string{
		"${VOLUME_ENV_VAR}:${VOLUME_ENV_VAR}",
		"/etc/passwd:/etc/passwd:ro",
		"/etc/group:/etc/group:ro",
		"/run/docker.sock:/run/docker.sock",
	}
	envVars := []string{
		"HOME:${ENV_VAR}",
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
	addGroups := false   // disabled because of different values on build than on local...
	impersonate := false // disabled because of different values on build than on local...
	removeContainer := true
	workDir := "someDir/"
	volumes := []string{
		"${VOLUME_ENV_VAR}:${VOLUME_ENV_VAR}",
		"/etc/passwd:/etc/passwd:ro",
		"/etc/group:/etc/group:ro",
		"/run/docker.sock:/run/docker.sock",
	}
	envVars := []string{
		"HOME:${ENV_VAR}",
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
