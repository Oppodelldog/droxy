package dockercommand

import (
	"errors"
	"os/exec"
	"reflect"
	"testing"

	"github.com/Oppodelldog/droxy/config"
)

//nolint:funlen
func TestBuildCommandFromConfig(t *testing.T) {
	var (
		cmdStub    = exec.Command("test")
		errStub    = errors.New("error stub")
		versionErr = errors.New("version error stub")
	)

	testCases := map[string]struct { //nolint:maligned
		containerExists bool
		version         string
		versionErr      error
		wantCmd         *exec.Cmd
		wantErr         error
		wantRun         bool
		wantExec        bool
	}{
		"container not exists, uses run builder": {
			containerExists: false,
			wantRun:         true,
			wantCmd:         cmdStub,
			wantErr:         errStub,
		},
		"container exists, uses exec builder": {
			containerExists: true,
			wantExec:        true,
			wantCmd:         cmdStub,
			wantErr:         errStub,
		},
		"version provider returns err, expect no cmd and version error": {
			versionErr: versionErr,
			wantCmd:    nil,
			wantErr:    versionErr,
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			var (
				containerName = "container1"
				runBuilder    = &builderMock{command: testCase.wantCmd, err: errStub}
				execBuilder   = &builderMock{command: testCase.wantCmd, err: errStub}
			)

			versionProvider := &versionMock{
				resValue: testCase.version,
				resErr:   testCase.versionErr,
			}
			existenceChecker := &existenceMock{
				result: testCase.containerExists,
			}

			b := Builder{
				containerExistenceChecker: existenceChecker,
				versionProvider:           versionProvider,
				newRun:                    func(dockerVersion string) CommandBuilder { return runBuilder },
				newExec:                   func(dockerVersion string) CommandBuilder { return execBuilder },
			}
			commandDef := config.CommandDefinition{
				Name: &containerName,
			}

			cmd, err := b.BuildCommandFromConfig(commandDef)
			if !reflect.DeepEqual(testCase.wantCmd, cmd) {
				t.Fatalf("\nwant:\n%v\ngot:\n%v\n", testCase.wantCmd, cmd)
			}

			if !reflect.DeepEqual(testCase.wantErr, err) {
				t.Fatalf("\nwant:\n%v\ngot:\n%v\n", testCase.wantErr, err)
			}

			if testCase.wantRun {
				assertBuilder(t, runBuilder, commandDef)
			}

			if testCase.wantExec {
				assertBuilder(t, execBuilder, commandDef)
			}
		})
	}
}

func assertBuilder(t *testing.T, b *builderMock, expectedCommandDef config.CommandDefinition) {
	if b.calls != 1 {
		t.Fatalf("expected builder to be called, but got %v calls", b.calls)
	}

	if !reflect.DeepEqual(expectedCommandDef, b.commandDef) {
		t.Fatalf("expected command def failed\nwant:\n%#v\ngot:\n%#v\n", expectedCommandDef, b.commandDef)
	}
}

type existenceMock struct {
	calls         int
	containerName string
	result        bool
}

func (e *existenceMock) exists(containerName string) bool {
	e.calls++
	e.containerName = containerName

	return e.result
}

type versionMock struct {
	calls    int
	resValue string
	resErr   error
}

func (v *versionMock) getAPIVersion() (string, error) {
	v.calls++

	return v.resValue, v.resErr
}

type builderMock struct {
	calls      int
	commandDef config.CommandDefinition
	command    *exec.Cmd
	err        error
}

func (b *builderMock) BuildCommandFromConfig(commandDef config.CommandDefinition) (*exec.Cmd, error) {
	b.calls++
	b.commandDef = commandDef

	return b.command, b.err
}
