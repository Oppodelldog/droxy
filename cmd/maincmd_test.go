package cmd

import (
	"os"
	"testing"

	"github.com/spf13/cobra"
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
	originalArgs := os.Args
	defer func() {
		os.Args = originalArgs
	}()

	testDataSet := map[string]struct {
		args           []string
		expectedResult bool
	}{
		"flag is not set": {args: []string{}, expectedResult: false},
		"flag is set":     {args: []string{"--is-it-droxy"}, expectedResult: true},
	}

	for testCaseName, testData := range testDataSet {
		t.Run(testCaseName, func(t *testing.T) {
			os.Args = testData.args
			assert.Exactly(t, testData.expectedResult, shallRevealItsDroxy())
		})
	}
}

func Test_isSubCommand(t *testing.T) {
	type args struct {
		commandName string
		commands    []*cobra.Command
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "empty search-for, no commands", args: args{}, want: false},
		{name: "search-for given, no commands", args: args{commandName: "some-sub-command"}, want: false},
		{name: "search-for given, one non matching command", args: args{commandName: "command-A", commands: []*cobra.Command{{Use: "command-B"}}}, want: false},
		{name: "search-for given, one matching command", args: args{commandName: "command-B", commands: []*cobra.Command{{Use: "command-B"}}}, want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubCommand(tt.args.commandName, tt.args.commands); got != tt.want {
				t.Errorf("isSubCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
