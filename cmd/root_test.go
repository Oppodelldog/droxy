package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/spf13/cobra"
)

func TestNewRoot(t *testing.T) {
	assert.NotNil(t, NewRoot())
	assert.IsType(t, new(cobra.Command), NewRoot())
}

func TestRoot_Use(t *testing.T) {
	assert.Equal(t, "droxy", NewRoot().Use)
}
