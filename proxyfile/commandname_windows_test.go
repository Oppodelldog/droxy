package proxyfile

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCommandName(t *testing.T) {
	assert.Equal(t, "droxy.exe", GetCommandName())
}

func TestParseCommandNameFromCommandLine(t *testing.T) {
	possibleTestRunners := []string{
		"___TestParseCommandNameFromCommandLine_in_commandname_windows_test_go",
		"proxyfile.test",
	}

	assert.Contains(t, possibleTestRunners, ParseCommandNameFromCommandLine())
}

func TestGetCommandNameFilename(t *testing.T) {
	assert.Equal(t, "testFileName.exe", GetCommandNameFilename("testFileName"))
}
