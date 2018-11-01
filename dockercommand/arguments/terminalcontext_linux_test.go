package arguments

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isTerminalContext(t *testing.T) {
	assert.False(t, isTerminalContext())
}
