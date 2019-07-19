package dockercommand

import (
	"fmt"
	"path/filepath"
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

	err := os.Setenv("VOLUME_ENV_VAR", "volEnvVarStub")
	if err != nil {
		t.Fatalf("Did not expect os.Setenv to return an error, but got: %v", err)
	}
	err = os.Setenv("LINK_ENV_VAR", "linkEnvVarStub")
	if err != nil {
		t.Fatalf("Did not expect os.Setenv to return an error, but got: %v", err)
	}
	err = os.Setenv("ENV_VAR", "envVarStub")
	if err != nil {
		t.Fatalf("Did not expect os.Setenv to return an error, but got: %v", err)
	}

	commandName := "some-command"
	configuration := getFullFeatureConfig(commandName)

	commandBuilder, err := NewCommandBuilder()
	if err != nil {
		t.Fatalf("Did not expect NewCommandBuilder to return an error, but got: %v", err)
	}
	cmd, err := commandBuilder.BuildCommandFromConfig(commandName, configuration)
	if err != nil {
		t.Fatalf("Did not expect BuildCommandFromConfig to return an error, but got: %v", err)
	}

	expectedArgsFromTestCall := strings.Join(os.Args[1:], " ")
	commandString := strings.Join(cmd.Args, " ")

	expectedHostDirMount := fmt.Sprintf("%s:%s", getTestHostDir(), getTestHostDir())

	expectedCommandStrings := []string{
		strings.TrimSpace(strings.Join([]string{"docker run -i --rm --name some-command -w " + getTestHostDir() + " -p 8080:9080 -p 8081:9081 -p 78129:78129 -v volEnvVarStub:volEnvVarStub -v /etc/passwd:/etc/passwd:ro -v /etc/group:/etc/group:ro -v /run/docker.sock:/run/docker.sock -v " + expectedHostDirMount + " --link linkEnvVarStub:linkEnvVarStub --link containerXY:aliasXY -e HOME:envVarStub -e SSH_AUTH_SOCK:/run/ssh.sock -e DOCKER_HOST=unix:///run/docker.sock -l droxy -a STDIN -a STDOUT -a STDERR --network some-docker-network --env-file .env --ip 127.1.2.3 --entrypoint some-entrypoint some-image:v1.02 some-cmd additionalArgument=123", expectedArgsFromTestCall}, " ")),
		strings.TrimSpace(strings.Join([]string{"docker run -t -i --rm --name some-command -w " + getTestHostDir() + " -p 8080:9080 -p 8081:9081 -p 78129:78129 -v volEnvVarStub:volEnvVarStub -v /etc/passwd:/etc/passwd:ro -v /etc/group:/etc/group:ro -v /run/docker.sock:/run/docker.sock -v " + expectedHostDirMount + " --link linkEnvVarStub:linkEnvVarStub --link containerXY:aliasXY -e HOME:envVarStub -e SSH_AUTH_SOCK:/run/ssh.sock -e DOCKER_HOST=unix:///run/docker.sock -l droxy -a STDIN -a STDOUT -a STDERR --network some-docker-network --env-file .env --ip 127.1.2.3 --entrypoint some-entrypoint some-image:v1.02 some-cmd additionalArgument=123", expectedArgsFromTestCall}, " ")),
	}

	assert.Contains(t, expectedCommandStrings, commandString)

	err = os.Unsetenv("VOLUME_ENV_VAR")
	if err != nil {
		t.Fatalf("Did not expect os.Unsetenv to return an error, but got: %v", err)
	}
	err = os.Unsetenv("LINK_ENV_VAR")
	if err != nil {
		t.Fatalf("Did not expect os.Unsetenv to return an error, but got: %v", err)
	}
	err = os.Unsetenv("ENV_VAR")
	if err != nil {
		t.Fatalf("Did not expect os.Unsetenv to return an error, but got: %v", err)
	}
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

	commandBuilder, err := NewCommandBuilder()
	if err != nil {
		t.Fatalf("Did not expect NewCommandBuilder to return an error, but got: %v", err)
	}
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

	err = os.Unsetenv("VOLUME_ENV_VAR")
	if err != nil {
		t.Fatalf("Did not expect os.Unsetenv to return an error, but got: %v", err)
	}
	err = os.Unsetenv("LINK_ENV_VAR")
	if err != nil {
		t.Fatalf("Did not expect os.Unsetenv to return an error, but got: %v", err)
	}
	err = os.Unsetenv("ENV_VAR")
	if err != nil {
		t.Fatalf("Did not expect os.Unsetenv to return an error, but got: %v", err)
	}
}

func TestBuildCommandFromConfig_ifContainerIsRunning_expectDockerExecCommand(t *testing.T) {

	testDataSet := map[string]struct {
		containerExists       bool
		expectedDockerCommand string
	}{
		"if container does not exists, expect 'docker run' command": {
			containerExists:       false,
			expectedDockerCommand: "docker run",
		},
		"if container already exists, expect 'docker exec' command": {
			containerExists:       true,
			expectedDockerCommand: "docker exec",
		},
	}

	for testCaseName, testData := range testDataSet {
		t.Run(testCaseName, func(t *testing.T) {

			commandName := "some-command"
			configuration := getFullFeatureConfig(commandName)

			cb := &commandBuilder{
				dockerVersionProvider:     newDockerAPIVersionStub("1.25"),
				containerExistenceChecker: newContainerExistenceChecker(testData.containerExists),
			}

			cmd, err := cb.BuildCommandFromConfig(commandName, configuration)
			if err != nil {
				t.Fatalf("Did not expect BuildCommandFromConfig to return an error, but got: %v", err)
			}

			commandString := strings.Join(cmd.Args, " ")

			assert.Contains(t, commandString, testData.expectedDockerCommand)

		})
	}
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
	isDetached := false
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
	workDir := getTestHostDir()
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
		IsDetached:      &isDetached,
		IsDaemon:        &isDetached,
		EntryPoint:      &entryPoint,
		Command:         &command,
		Name:            &name,
		Image:           &image,
		Network:         &network,
		EnvFile:         &envFile,
		IP:              &ip,
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

func getTestHostDir() string {
	hostDir, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("Did not expect os.Getwd() to return an error, but got: %v", err))
	}
	absoluteHostPath, err := filepath.Abs(hostDir)
	if err != nil {
		panic(fmt.Sprintf("could not get absolute path of the tests working dir: %v", err))
	}
	return absoluteHostPath
}

func getFullFeatureDef(commandName string) config.CommandDefinition {
	isTemplate := true
	isDetached := false
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
	workDir := getTestHostDir()
	autoMountWorkDir := true
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
		IsTemplate:       &isTemplate,
		IsDetached:       &isDetached,
		IsDaemon:         &isDetached,
		Template:         &template,
		EntryPoint:       &entryPoint,
		Command:          &command,
		Name:             &name,
		Image:            &image,
		Network:          &network,
		EnvFile:          &envFile,
		IP:               &ip,
		IsInteractive:    &isInteractive,
		AddGroups:        &addGroups,
		Impersonate:      &impersonate,
		RemoveContainer:  &removeContainer,
		WorkDir:          &workDir,
		AutoMountWorkDir: &autoMountWorkDir,
		Volumes:          &volumes,
		Links:            &links,
		EnvVars:          &envVars,
		Ports:            &ports,
		PortsFromParams:  &portsFromParams,
		ReplaceArgs:      &replaceArgs,
		AdditionalArgs:   &additionalArgs,
	}
}

type containerExistenceCheckerStub struct {
	containerExists bool
}

func (c *containerExistenceCheckerStub) exists(containerName string) bool {
	return c.containerExists
}

func newContainerExistenceChecker(containerExists bool) *containerExistenceCheckerStub {
	return &containerExistenceCheckerStub{containerExists}
}

type dockerAPIVersionStub struct {
	dockerAPIVersion string
}

func (s *dockerAPIVersionStub) getAPIVersion() (string, error) {
	return s.dockerAPIVersion, nil
}

func newDockerAPIVersionStub(apiVersion string) *dockerAPIVersionStub {
	return &dockerAPIVersionStub{apiVersion}
}

func TestBuildCommandFromConfig_singleArgumentTest(t *testing.T) {
	testDataSet := map[string]struct {
		dockerAPIVersion         string
		expectedCommandArgString string
		expectCommandContainsArg bool
	}{
		"workdir is not added for versions below 1.35": {
			dockerAPIVersion:         "1.34",
			expectedCommandArgString: "-w",
			expectCommandContainsArg: false,
		},
		"workdir is added for versions from 1.35": {
			dockerAPIVersion:         "1.35",
			expectedCommandArgString: "-w",
			expectCommandContainsArg: true,
		},
		"envvars are not added for versions below 1.25": {
			dockerAPIVersion:         "1.24",
			expectedCommandArgString: "-e",
			expectCommandContainsArg: false,
		},
		"envvars are added for versions from 1.25": {
			dockerAPIVersion:         "1.25",
			expectedCommandArgString: "-e",
			expectCommandContainsArg: true,
		},
	}

	for testCaseName, testData := range testDataSet {
		t.Run(testCaseName, func(t *testing.T) {
			cb := commandBuilder{
				dockerVersionProvider: newDockerAPIVersionStub(testData.dockerAPIVersion),
			}

			workDir := "someWorkDir"
			envVars := []string{"SOME_ENV_VAR"}
			commandDef := &config.CommandDefinition{
				WorkDir: &workDir,
				EnvVars: &envVars,
			}
			cmd, err := cb.buildExecCommand(commandDef)
			if err != nil {
				t.Fatalf("did not expect buildExecCommand to return an error, but got: %v", err)
			}

			if testData.expectCommandContainsArg {
				assert.Contains(t, cmd.Args, testData.expectedCommandArgString)
			} else {
				assert.NotContains(t, cmd.Args, testData.expectedCommandArgString)
			}
		})
	}
}
