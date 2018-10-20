package dockercommand

import (
	"strings"
	"testing"

	"github.com/Oppodelldog/droxy/config"

	"os"

	"github.com/stretchr/testify/assert"
)

func TestBuildCommandFromConfig(t *testing.T) {
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	os.Args = append(os.Args, "--inspect-brk=78129")

	os.Setenv("VOLUME_ENV_VAR", "volEnvVarStub")
	os.Setenv("LINK_ENV_VAR", "linkEnvVarStub")
	os.Setenv("ENV_VAR", "envVarStub")

	commandName := "some-command"
	configuration := getFullFeatureConfig(commandName)

	commandBuilder := NewCommandBuilder()
	cmd, err := commandBuilder.BuildCommandFromConfig(commandName, configuration)
	if err != nil {
		t.Fatalf("Did not expect BuildCommandFromConfig to return an error, but got: %v", err)
	}

	expectedArgsFromTestCall := strings.Join(os.Args[1:], " ")
	commandString := strings.Join(cmd.Args, " ")

	expectedCommandStrings := []string{
		strings.TrimSpace(strings.Join([]string{"docker run -i -d --rm --name some-command -w someDir/ -p 8080:9080 -p 8081:9081 -p 78129:78129 -v volEnvVarStub:volEnvVarStub -v /etc/passwd:/etc/passwd:ro -v /etc/group:/etc/group:ro -v /run/docker.sock:/run/docker.sock --link linkEnvVarStub:linkEnvVarStub --link containerXY:aliasXY -e HOME:envVarStub -e SSH_AUTH_SOCK:/run/ssh.sock -e DOCKER_HOST=unix:///run/docker.sock -l droxy -a STDIN -a STDOUT -a STDERR --network some-docker-network --env-file .env --ip 127.1.2.3 --entrypoint some-entrypoint some-image:v1.02 some-cmd additionalArgument=123", expectedArgsFromTestCall}, " ")),
		strings.TrimSpace(strings.Join([]string{"docker run -t -i -d --rm --name some-command -w someDir/ -p 8080:9080 -p 8081:9081 -p 78129:78129 -v volEnvVarStub:volEnvVarStub -v /etc/passwd:/etc/passwd:ro -v /etc/group:/etc/group:ro -v /run/docker.sock:/run/docker.sock --link linkEnvVarStub:linkEnvVarStub --link containerXY:aliasXY -e HOME:envVarStub -e SSH_AUTH_SOCK:/run/ssh.sock -e DOCKER_HOST=unix:///run/docker.sock -l droxy -a STDIN -a STDOUT -a STDERR --network some-docker-network --env-file .env --ip 127.1.2.3 --entrypoint some-entrypoint some-image:v1.02 some-cmd additionalArgument=123", expectedArgsFromTestCall}, " ")),
	}

	assert.Contains(t, expectedCommandStrings, commandString)

	os.Unsetenv("VOLUME_ENV_VAR")
	os.Unsetenv("LINK_ENV_VAR")
	os.Unsetenv("ENV_VAR")
}

func TestBuildCommandFromConfig_EmptyCommandDoesNotProduceSpaceInCommand(t *testing.T) {

	commandName := "some-command"

	configuration := &config.Configuration{
		Command: []config.CommandDefinition{
			{
				Name: &commandName,
			},
		},
	}

	commandBuilder := NewCommandBuilder()
	cmd, err := commandBuilder.BuildCommandFromConfig(commandName, configuration)
	if err != nil {
		t.Fatalf("Did not expect BuildCommandFromConfig to return an error, but got: %v", err)
	}

	expectedArgsFromTestCall := strings.Join(os.Args[1:], " ")
	commandString := strings.Join(cmd.Args, " ")

	expectedCommandStrings := []string{
		strings.TrimSpace(strings.Join([]string{"docker run --name some-command -l droxy -a STDIN -a STDOUT -a STDERR", expectedArgsFromTestCall}, " ")),
		strings.TrimSpace(strings.Join([]string{"docker run -t --name some-command -l droxy -a STDIN -a STDOUT -a STDERR", expectedArgsFromTestCall}, " ")),
	}

	assert.Contains(t, expectedCommandStrings, commandString)

	os.Unsetenv("VOLUME_ENV_VAR")
	os.Unsetenv("LINK_ENV_VAR")
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
	isDaemon := true
	entryPoint := "some-entrypoint"
	command := "some-cmd"
	name := "some template"
	image := "some-image:v1.02"
	network := "some-docker-network"
	envFile := ".env"
	ip := "127.1.2.3"
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
	links := []string{
		"${LINK_ENV_VAR}:${LINK_ENV_VAR}",
		"containerXY:aliasXY",
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
	portsFromParams := []string{
		"--inspect-brk=(\\d*)",
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
		IsDaemon:        &isDaemon,
		EntryPoint:      &entryPoint,
		Command:         &command,
		Name:            &name,
		Image:           &image,
		Network:         &network,
		EnvFile:         &envFile,
		Ip:              &ip,
		IsInteractive:   &isInteractive,
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

func getFullFeatureDef(commandName string) config.CommandDefinition {
	isTemplate := true
	isDaemon := true
	template := "some template"
	entryPoint := "some-entrypoint"
	command := "some-cmd"
	name := commandName
	image := "some-image:v1.02"
	network := "some-docker-network"
	envFile := ".env"
	ip := "127.1.2.3"
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
	links := []string{
		"${LINK_ENV_VAR}:${LINK_ENV_VAR}",
		"containerXY:aliasXY",
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
	portsFromParams := []string{
		"--inspect-brk=(\\d*)",
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
		IsDaemon:        &isDaemon,
		Template:        &template,
		EntryPoint:      &entryPoint,
		Command:         &command,
		Name:            &name,
		Image:           &image,
		Network:         &network,
		EnvFile:         &envFile,
		Ip:              &ip,
		IsInteractive:   &isInteractive,
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
