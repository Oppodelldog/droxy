package helper

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCommandName(t *testing.T) {
	assert.Equal(t, "droxy", GetCommandName())
}

func TestParseCommandNameFromCommandLine(t *testing.T) {
	originalValue := os.Args[0]
	defer func() {
		os.Args[0] = originalValue
	}()

	os.Args[0] = "/tmp/test123"
	assert.Equal(t, "test123", ParseCommandNameFromCommandLine())
}

func TestGetCommandNameFilename(t *testing.T) {
	assert.Equal(t, "testFileName", GetCommandNameFilename("testFileName"))
}
