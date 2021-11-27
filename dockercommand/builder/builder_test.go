package builder

import (
	"bytes"
	"io"
	"os/exec"
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

	cmd := buildCmd(testWriterA, testWriterB, testReader)

	got := strings.Join(cmd.Args, " ")
	want := getExpectedCmdString()

	assert.Equal(t, want, got)
	assert.Equal(t, cmd.Stdout, testWriterA)
	assert.Equal(t, cmd.Stderr, testWriterB)
	assert.Equal(t, cmd.Stdin, testReader)
}

func buildCmd(w1 io.Writer, w2 io.Writer, r io.Reader) *exec.Cmd {
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
	b.SetStdOut(w1)
	b.SetStdErr(w2)
	b.SetStdIn(r)
	b.SetWorkingDir("workingDir")

	return b.Build()
}

func getExpectedCmdString() string {
	return strings.ReplaceAll(`docker run
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
		"\n", " ")
}

func TestBuilder_DefaultOutput(t *testing.T) {
	b := New()
	cmd := b.Build()

	commandString := strings.Join(cmd.Args, " ")
	expectedCommandString := `docker run`

	assert.Equal(t, expectedCommandString, commandString)
}
