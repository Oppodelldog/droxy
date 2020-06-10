package proxyexecution

import (
	"testing"

	"github.com/Oppodelldog/droxy/crossplatform"

	"github.com/stretchr/testify/assert"
)

func TestNewExecutableNameParser(t *testing.T) {
	assert.IsType(t, executableNameParser{}, newExecutableNameParser())
}

func TestExecutableNameParser_ParseCommandNameFromCommandLine(t *testing.T) {
	executableNameParser := newExecutableNameParser()
	parsedCommandName := executableNameParser.ParseCommandNameFromCommandLine()

	assert.Equal(t, crossplatform.ParseCommandNameFromCommandLine(), parsedCommandName)
}
