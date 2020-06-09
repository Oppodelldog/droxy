package builder

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBuilder(t *testing.T) {
	assert.IsType(t, new(builder), New())
}

func TestBuilder_FullFeature(t *testing.T) {
	testWriterA := bytes.NewBufferString("")
	testWriterB := bytes.NewBufferString("")
	testReader := bytes.NewBufferString("")

	b := New()
	b.AddArgument("arg1")
	b.AddArgument("arg2")
	b.AddCmdArguments([]string{"cmdArg1", "cmdArg2"})
	b.AddEnvVar("envVar1")
	b.AddEnvVar("envVar2")
	b.AddGroup("group1")
	b.AddGroup("group2")
	b.AddPortMapping("portMapping1")
	b.AddPortMapping("portMapping2")
	b.AddVolumeMapping("volumeMapping1Host:volumeMapping1Container:volumeMapping1Options")
	b.AddVolumeMapping("volumeMapping2Host:volumeMapping2Container:volumeMapping2Options")
	b.AttachTo("Stdin")
	b.AttachTo("Stdout")
	b.AttachTo("Stderr")
	b.SetContainerName("containerName")
	b.SetContainerUserAndGroup("userId", "userGroup")
	b.SetCommand("command")
	b.SetImageName("imageName")
	b.SetNetwork("network")
	b.SetEnvFile(".env")
	b.SetIP("127.1.2.3")
	b.SetStdOut(testWriterA)
	b.SetStdErr(testWriterB)
	b.SetStdIn(testReader)
	b.SetWorkingDir("workingDir")
	cmd := b.Build()

	commandString := strings.Join(cmd.Args, " ")
	expectedCommandString := strings.Replace(`docker run
arg1
arg2
--name containerName
-w workingDir
-p portMapping1
-p portMapping2
-v volumeMapping1Host:volumeMapping1Container:volumeMapping1Options
-v volumeMapping2Host:volumeMapping2Container:volumeMapping2Options
-e envVar1
-e envVar2
--group-add group1
--group-add group2
-u userId:userGroup
-a Stdin
-a Stdout
-a Stderr
--network network
--env-file .env
--ip 127.1.2.3
imageName
command
cmdArg1
cmdArg2`,
		"\n", " ", -1)

	assert.Equal(t, expectedCommandString, commandString)
	assert.Equal(t, cmd.Stdout, testWriterA)
	assert.Equal(t, cmd.Stderr, testWriterB)
	assert.Equal(t, cmd.Stdin, testReader)
}

func TestBuilder_DefaultOutput(t *testing.T) {
	b := New()
	cmd := b.Build()

	commandString := strings.Join(cmd.Args, " ")
	expectedCommandString := `docker run`

	assert.Equal(t, expectedCommandString, commandString)
}
