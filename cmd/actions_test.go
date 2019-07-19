package cmd

import (
	"reflect"
	"testing"

	"github.com/pkg/errors"

	"github.com/stretchr/testify/assert"
)

func Test_stringInSlice_StringIsInSlice(t *testing.T) {
	s := "abc"
	slice := []string{"iop", "zui", "abc"}
	res := stringInSlice(s, slice)

	assert.True(t, res)
}

func Test_stringInSlice_StringIsNotInSlice(t *testing.T) {
	s := "abc"
	slice := []string{"iop", "zui", "abf"}
	res := stringInSlice(s, slice)

	assert.False(t, res)
}

func Test_shallRevealItsDroxy(t *testing.T) {
	testDataSet := map[string]struct {
		args           []string
		expectedResult bool
	}{
		"flag is not set": {args: []string{}, expectedResult: false},
		"flag is set":     {args: []string{"--is-it-droxy"}, expectedResult: true},
	}

	for testCaseName, testData := range testDataSet {
		t.Run(testCaseName, func(t *testing.T) {
			assert.Exactly(t, testData.expectedResult, shallRevealItsDroxy(testData.args))
		})
	}
}

func Test_getActionChain(t *testing.T) {
	tests := []struct {
		name  string
		index int
		want  actionChainElement
	}{
		{
			name:  "1st element is subcommand action",
			index: 0,
			want:  newSubCommandAction(nil),
		},
		{
			name:  "2nd element is help display action",
			index: 1,
			want:  newHelpDisplayAction(nil),
		},
		{
			name:  "3rd element is revealThatItsDroxy action",
			index: 2,
			want:  newRevealItsDroxyAction(),
		},
		{
			name:  "last element is droxy command action",
			index: len(getActionChain()) - 1,
			want:  newDroxyCommandAction(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ObjectsAreEqual(tt.want, getActionChain()[tt.index])
		})
	}
}

func Test_newSubCommandAction_isResponsible(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want bool
	}{
		{name: "no arguments", args: []string{}, want: false},
		{name: "one arguments", args: []string{testCommandName}, want: false},
		{name: "two arguments, but no subcommand", args: []string{testCommandName, "nonexistent"}, want: false},
		{name: "three arguments, but no subcommand", args: []string{testCommandName}, want: false},
		{name: "two arguments with subcommand", args: []string{testCommandName, "clones"}, want: true},
		{name: "three arguments with subcommand", args: []string{testCommandName, "clones"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newSubCommandAction(nil).(*action).isResponsibleFunc(tt.args); got != tt.want {
				t.Errorf("newSubCommandAction().isResponsibleFunc(%v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}

func Test_newRevealItsDroxyAction_isResponsible(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want bool
	}{
		{name: "no arguments", args: []string{}, want: false},
		{name: "one arguments", args: []string{testCommandName}, want: false},
		{name: "two arguments, but no match", args: []string{testCommandName, "hello"}, want: false},
		{name: "three arguments, but no match", args: []string{testCommandName}, want: false},
		{name: "four arguments with a match", args: []string{testCommandName, "--is-it-droxy", "test", "clones"}, want: true},
		{name: "one argument which is a match", args: []string{"--is-it-droxy"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newRevealItsDroxyAction().(*action).isResponsibleFunc(tt.args); got != tt.want {
				t.Errorf("newRevealItsDroxyAction().isResponsibleFunc(%v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}

func Test_newHelpDisplayAction_isResponsible(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want bool
	}{
		{name: "no arguments", args: []string{}, want: false},
		{name: "one argument", args: []string{testCommandName}, want: true},
		{name: "two arguments, but no match", args: []string{testCommandName, "hello"}, want: true},
		{name: "three arguments, but no match", args: []string{testCommandName, "test", "123"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newHelpDisplayAction(nil).(*action).isResponsibleFunc(tt.args); got != tt.want {
				t.Errorf("newHelpDisplayAction().isResponsibleFunc(%v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}

func Test_newDroxyCommandAction(t *testing.T) {
	assert.True(t, newDroxyCommandAction().(*action).isResponsibleFunc(nil))
}

func Test_newDroxyCommandAction_defaultExecuteFunc(t *testing.T) {

	if reflect.ValueOf(newDroxyCommandAction().(*action).executeFunc).Pointer() != reflect.ValueOf(defaultExecuteFunc).Pointer() {
		t.Fatal("expected newDroxyCommand to be configured with 'defaultExecuteFunc', but was not")
	}
}

func Test_revealTheTruth(t *testing.T) {
	res := revealTheTruth()

	assert.Exactly(t, 0, res)
}

type executerMock struct {
	wasCalled bool
	result    error
}

func (e *executerMock) Execute() error {
	e.wasCalled = true
	return e.result
}

func Test_execSubCommand(t *testing.T) {

	mock := &executerMock{result: errors.New("for code coverage")}
	newSubCommandAction(mock).Execute()

	assert.True(t, mock.wasCalled)
}

type helperMock struct {
	wasCalled bool
	result    error
}

func (e *helperMock) Help() error {
	e.wasCalled = true
	return e.result
}

func Test_displayHelp(t *testing.T) {

	mock := &helperMock{result: errors.New("for code coverage")}
	newHelpDisplayAction(mock).Execute()

	assert.True(t, mock.wasCalled)
}
