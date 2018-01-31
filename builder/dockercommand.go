package builder

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

type (
	DockerCommandBuilder struct {
		command         string
		subCommand      string
		imageName       string
		entryPoint      string
		args            []string
		portMappings    []string
		volumeMappings  []string
		envVarMappings  []string
		attachedStreams []string
		workingDir      []string
		containerName   []string
		addedUserGroups []string
		containerUser   []string
		cmdArgs         []string
		stdIn           io.Reader
		stdOut          io.Writer
		stdErr          io.Writer

		buildArgs []string
	}
)

func NewDockerCommandBuilder() *DockerCommandBuilder {
	return &DockerCommandBuilder{
		command:    "docker",
		subCommand: "run",
		stdIn:      os.Stdin,
		stdOut:     os.Stdout,
		stdErr:     os.Stderr,
	}
}
func (b *DockerCommandBuilder) SetStdIn(r io.Reader) *DockerCommandBuilder {
	b.stdIn = r

	return b
}

func (b *DockerCommandBuilder) SetStdOut(w io.Writer) *DockerCommandBuilder {
	b.stdOut = w

	return b
}

func (b *DockerCommandBuilder) SetStdErr(w io.Writer) *DockerCommandBuilder {
	b.stdErr = w

	return b
}

func (b *DockerCommandBuilder) AddPortMapping(hostPort string, containerPort string) *DockerCommandBuilder {
	b.portMappings = append(b.portMappings, "-p", fmt.Sprintf("%s:%s", hostPort, containerPort))
	return b
}

func (b *DockerCommandBuilder) AddCmdArguments(arguments []string) *DockerCommandBuilder {
	b.cmdArgs = append(b.cmdArgs, arguments...)
	return b
}

func (b *DockerCommandBuilder) AddArgument(argument string) *DockerCommandBuilder {
	b.args = append(b.args, argument)
	return b
}

func (b *DockerCommandBuilder) AttachTo(stream string) *DockerCommandBuilder {
	b.attachedStreams = append(b.attachedStreams, "-a", stream)
	return b
}

func (b *DockerCommandBuilder) AddMirrorVolume(path string) *DockerCommandBuilder {
	b.volumeMappings = append(b.volumeMappings, "-v", fmt.Sprintf("%s:%s", path, path))
	return b
}

func (b *DockerCommandBuilder) AddVolume(hostPath string, containerPath string) *DockerCommandBuilder {
	b.volumeMappings = append(b.volumeMappings, "-v", fmt.Sprintf("%s:%s", hostPath, containerPath))
	return b
}

func (b *DockerCommandBuilder) AddVolumePlain(hostPath, containerPath, options string) *DockerCommandBuilder {
	s := bytes.NewBufferString("")
	if hostPath != "" {
		s.WriteString(hostPath)
	}
	if containerPath != "" {
		if s.Len() > 0 {
			s.WriteString(":")
		}
		s.WriteString(containerPath)
	}
	if options != "" {
		if s.Len() > 0 {
			s.WriteString(":")
		}
		s.WriteString(options)
	}
	b.volumeMappings = append(b.volumeMappings, "-v", s.String())
	return b
}

func (b *DockerCommandBuilder) AddMirrorVolumeReadOnly(path string) *DockerCommandBuilder {
	b.volumeMappings = append(b.volumeMappings, "-v", fmt.Sprintf("%s:%s:ro", path, path))
	return b
}

func (b *DockerCommandBuilder) AddEnvVar(envVarDeclaration string) *DockerCommandBuilder {
	b.envVarMappings = append(b.envVarMappings, "-e", envVarDeclaration)
	return b
}

func (b *DockerCommandBuilder) AddEnvVars(envVarDeclarations []string) *DockerCommandBuilder {
	for _, envVarDeclaration := range envVarDeclarations {
		b.envVarMappings = append(b.envVarMappings, "-e", envVarDeclaration)
	}

	return b
}

func (b *DockerCommandBuilder) AdduserGroup(userGroup string) *DockerCommandBuilder {
	b.addedUserGroups = append(b.addedUserGroups, "--group-add", userGroup)
	return b
}

func (b *DockerCommandBuilder) SetEntryPoint(entryPoint string) *DockerCommandBuilder {
	b.entryPoint = entryPoint
	return b
}

func (b *DockerCommandBuilder) SetImageName(imageName string) *DockerCommandBuilder {
	b.imageName = imageName
	return b
}

func (b *DockerCommandBuilder) SetWorkingDir(workingDir string) *DockerCommandBuilder {
	b.workingDir = []string{"-w", workingDir}
	return b
}

func (b *DockerCommandBuilder) SetContainerName(containerName string) *DockerCommandBuilder {
	b.containerName = []string{"--name", containerName}

	return b
}

func (b *DockerCommandBuilder) SetContainerUserAndGroup(userId string, groupId string) *DockerCommandBuilder {
	b.containerUser = []string{"-u", fmt.Sprintf("%s:%s", userId, groupId)}

	return b
}

func (b *DockerCommandBuilder) Build() *exec.Cmd {

	cmd := exec.Command(b.command, b.subCommand)

	b.buildArgsAppend(b.args...)
	b.buildArgsAppend(b.containerName...)
	b.buildArgsAppend(b.workingDir...)
	b.buildArgsAppend(b.portMappings...)
	b.buildArgsAppend(b.volumeMappings...)
	b.buildArgsAppend(b.envVarMappings...)
	b.buildArgsAppend(b.addedUserGroups...)
	b.buildArgsAppend(b.containerUser...)
	b.buildArgsAppend(b.attachedStreams...)

	b.buildArgAppend(b.imageName)
	b.buildArgAppend(b.entryPoint)

	b.buildArgsAppend(b.cmdArgs...)

	cmd.Args = append(cmd.Args, b.buildArgs...)

	cmd.Stdout = b.stdOut
	cmd.Stderr = b.stdErr
	cmd.Stdin = b.stdIn

	return cmd
}

func (b *DockerCommandBuilder) buildArgAppend(arg string) {
	b.buildArgs = append(b.buildArgs, arg)
}

func (b *DockerCommandBuilder) buildArgsAppend(args ...string) {
	b.buildArgs = append(b.buildArgs, args...)
}
