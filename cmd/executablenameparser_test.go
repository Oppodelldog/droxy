package cmd

import (
	"testing"

	"github.com/Oppodelldog/droxy/helper"
	"github.com/stretchr/testify/assert"
)

func TestNewExecutableNameParser(t *testing.T) {
	assert.Implements(t, new(ExecutableNameParser), NewExecutableNameParser())
}

func TestExecutableNameParser_ParseCommandNameFromCommandLine(t *testing.T) {
	executableNameParser := NewExecutableNameParser()
	parsedCommandName := executableNameParser.ParseCommandNameFromCommandLine()

	assert.Equal(t, helper.ParseCommandNameFromCommandLine(), parsedCommandName)
}
