package builder

import (
	"io"
	"os/exec"
)

//Builder describes the interface of the "docker run" command builder
type Builder interface {
	SetStdIn(r io.Reader) Builder
	SetStdOut(w io.Writer) Builder
	SetStdErr(w io.Writer) Builder
	AddPortMapping(portMapping string) Builder
	AddCmdArguments(arguments []string) Builder
	AddArgument(argument string) Builder
	AttachTo(stream string) Builder
	AddVolumeMapping(volumeMapping string) Builder
	AddEnvVar(envVarDeclaration string) Builder
	AddGroup(groupName string) Builder
	SetEntryPoint(entryPoint string) Builder
	SetCommand(command string) Builder
	SetNetwork(network string) Builder
	SetEnvFile(envFile string) Builder
	SetIp(ip string) Builder
	SetImageName(imageName string) Builder
	SetWorkingDir(workingDir string) Builder
	SetContainerName(containerName string) Builder
	SetContainerUserAndGroup(userID string, groupID string) Builder
	Build() *exec.Cmd
}
