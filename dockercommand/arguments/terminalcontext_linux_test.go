package arguments

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isTerminalContext(t *testing.T) {
	assert.False(t, isTerminalContext())
}
