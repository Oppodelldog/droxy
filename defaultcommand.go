package docker_proxy_command

import (
	"docker-proxy-command/builder"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"syscall"
	"time"
	"unsafe"
)

func getUserHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	home := usr.HomeDir

	return home
}

func isTerminalContext() bool {
	var termios syscall.Termios
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, os.Stdin.Fd(), syscall.TCGETS, uintptr(unsafe.Pointer(&termios)))
	return err == 0
}

func getExecutableInfo() (string, string) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fileInfo, err := os.Stat(ex)
	if err != nil {
		panic(err)
	}

	return fileInfo.Name(), exPath
}

func BuildDefaultCommand(dockerImage, dockerEntryPoint string) (*builder.DockerCommandBuilder, error) {

	workingDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	home := getUserHomeDir()
	builder := builder.NewDockerCommandBuilder()
	builder.
		SetImageName(dockerImage).
		SetEntryPoint(dockerEntryPoint)
	builder.
		AddArgument("--rm").
		AddArgument("-i")
	if isTerminalContext() {
		builder.AddArgument("-t")
	}
	builder.
		AddMirrorVolume(home).
		AddMirrorVolume(workingDir).
		AddMirrorVolume("/run/docker.sock")
	builder.SetWorkingDir(workingDir)
	builder.AddVolume(workingDir, "/app")
	builder.AddMirrorVolume("/tmp")
	builder.
		AddMirrorVolumeReadOnly("/etc/passwd").
		AddMirrorVolumeReadOnly("/etc/group")
	builder.AddEnvVar("DOCKER_HOST=unix:///run/docker.sock")
	if authSock, isSet := os.LookupEnv("SSH_AUTH_SOCK"); isSet {
		builder.AddVolume(authSock, "/run/ssh.sock")
		builder.AddEnvVar("SSH_AUTH_SOCK=/run/ssh.sock")
	}

	//really map all user environment??? builder.AddEnvVars(os.Environ())
	builder.
		AttachTo("STDIN").
		AttachTo("STDOUT").
		AttachTo("STDERR")
	commandName, _ := getExecutableInfo()
	commandInstanceName := fmt.Sprintf("%s_%s_%v", commandName, "cmd_proxy", time.Now().Nanosecond())
	builder.SetContainerName(commandInstanceName)
	currentUser, err := user.Current()
	if err != nil {
		return nil, err
	}
	groupIds, err := currentUser.GroupIds()
	if err != nil {
		return nil, err
	}
	if len(groupIds) > 0 {
		for _, groupId := range groupIds {
			builder.AdduserGroup(groupId)
		}
	}
	builder.SetContainerUserAndGroup(currentUser.Uid, currentUser.Gid)
	builder.AddCmdArguments(os.Args[1:])

	return builder, nil
}
