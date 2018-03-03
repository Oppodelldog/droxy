package helper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCommandName(t *testing.T) {
	assert.Equal(t, "droxy.exe", GetCommandName())
}

func TestParseCommandNameFromCommandLine(t *testing.T) {
	assert.Equal(t, "helper.test", ParseCommandNameFromCommandLine())
}

func TestGetCommandNameFilename(t *testing.T) {
	assert.Equal(t, "testFileName.exe", GetCommandNameFilename("testFileName"))
}
