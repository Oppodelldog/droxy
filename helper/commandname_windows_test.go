package helper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCommandName(t *testing.T) {
	assert.Equal(t, "droxy.exe", GetCommandName())
}

func TestParseCommandNameFromCommandLine(t *testing.T) {
	possibleTestRunners := []string{
		"___TestParseCommandNameFromCommandLine_in_commandname_windows_test_go",
		"helper.test",
	}

	assert.Contains(t, possibleTestRunners, ParseCommandNameFromCommandLine())
}

func TestGetCommandNameFilename(t *testing.T) {
	assert.Equal(t, "testFileName.exe", GetCommandNameFilename("testFileName"))
}
