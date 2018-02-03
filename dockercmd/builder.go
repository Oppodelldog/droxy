package dockercmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

type (
	Builder struct {
		command         string
		subCommand      string
		imageName       string
		entryPoint      string
		network         []string
		args            []string
		portMappings    []string
		volumeMappings  []string
		envVarMappings  []string
		attachedStreams []string
		workingDir      []string
		containerName   []string
		addedGroups     []string
		containerUser   []string
		cmdArgs         []string
		stdIn           io.Reader
		stdOut          io.Writer
		stdErr          io.Writer

		buildArgs []string
	}
)

func NewBuilder() *Builder {
	return &Builder{
		command:    "docker",
		subCommand: "run",
		stdIn:      os.Stdin,
		stdOut:     os.Stdout,
		stdErr:     os.Stderr,
	}
}
func (b *Builder) SetStdIn(r io.Reader) *Builder {
	b.stdIn = r

	return b
}

func (b *Builder) SetStdOut(w io.Writer) *Builder {
	b.stdOut = w

	return b
}

func (b *Builder) SetStdErr(w io.Writer) *Builder {
	b.stdErr = w

	return b
}

func (b *Builder) AddPortMapping(hostPort string, containerPort string) *Builder {
	b.portMappings = append(b.portMappings, "-p", fmt.Sprintf("%s:%s", hostPort, containerPort))
	return b
}

func (b *Builder) AddCmdArguments(arguments []string) *Builder {
	b.cmdArgs = append(b.cmdArgs, arguments...)
	return b
}

func (b *Builder) AddArgument(argument string) *Builder {
	b.args = append(b.args, argument)
	return b
}

func (b *Builder) AttachTo(stream string) *Builder {
	b.attachedStreams = append(b.attachedStreams, "-a", stream)
	return b
}

func (b *Builder) AddVolumeMapping(hostPath, containerPath, options string) *Builder {
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

func (b *Builder) AddEnvVar(envVarDeclaration string) *Builder {
	b.envVarMappings = append(b.envVarMappings, "-e", envVarDeclaration)
	return b
}

func (b *Builder) AddGroup(groupName string) *Builder {
	b.addedGroups = append(b.addedGroups, "--group-add", groupName)
	return b
}

func (b *Builder) SetEntryPoint(entryPoint string) *Builder {
	b.entryPoint = entryPoint
	return b
}

func (b *Builder) SetNetwork(network string) *Builder {
	b.network = []string{"--network", network}
	return b
}

func (b *Builder) SetImageName(imageName string) *Builder {
	b.imageName = imageName
	return b
}

func (b *Builder) SetWorkingDir(workingDir string) *Builder {
	b.workingDir = []string{"-w", workingDir}
	return b
}

func (b *Builder) SetContainerName(containerName string) *Builder {
	b.containerName = []string{"--name", containerName}

	return b
}

func (b *Builder) SetContainerUserAndGroup(userId string, groupId string) *Builder {
	b.containerUser = []string{"-u", fmt.Sprintf("%s:%s", userId, groupId)}

	return b
}

func (b *Builder) Build() *exec.Cmd {

	cmd := exec.Command(b.command, b.subCommand)

	b.buildArgsAppend(b.args...)
	b.buildArgsAppend(b.containerName...)
	b.buildArgsAppend(b.workingDir...)
	b.buildArgsAppend(b.portMappings...)
	b.buildArgsAppend(b.volumeMappings...)
	b.buildArgsAppend(b.envVarMappings...)
	b.buildArgsAppend(b.addedGroups...)
	b.buildArgsAppend(b.containerUser...)
	b.buildArgsAppend(b.attachedStreams...)
	b.buildArgsAppend(b.network...)

	b.buildArgAppend(b.imageName)
	b.buildArgAppend(b.entryPoint)

	b.buildArgsAppend(b.cmdArgs...)

	cmd.Args = append(cmd.Args, b.buildArgs...)

	cmd.Stdout = b.stdOut
	cmd.Stderr = b.stdErr
	cmd.Stdin = b.stdIn

	return cmd
}

func (b *Builder) buildArgAppend(arg string) {
	b.buildArgs = append(b.buildArgs, arg)
}

func (b *Builder) buildArgsAppend(args ...string) {
	b.buildArgs = append(b.buildArgs, args...)
}
