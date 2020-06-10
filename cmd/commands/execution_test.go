package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert.NotPanics(t, func() { Run([]string{testCommandName}) })
}
