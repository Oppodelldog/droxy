package arguments

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddGroups(t *testing.T) {
	res := addGroups(nil, nil)
	assert.Nil(t, res)
}
