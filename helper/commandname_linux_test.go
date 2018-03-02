package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCommandName(t *testing.T) {
	assert.Equal(t, "droxy", GetCommandName())
}

func TestParseCommandNameFromCommandLine(t *testing.T) {
	assert.Equal(t, "___commandname_linux_test_go", ParseCommandNameFromCommandLine())
}

func TestGetCommandNameFilename(t *testing.T) {
	assert.Equal(t, "testFileName", GetCommandNameFilename("testFileName"))
}
