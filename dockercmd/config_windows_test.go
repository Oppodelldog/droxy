package dockercmd

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAddGroups(t *testing.T) {
	res := addGroups(nil, nil)
	assert.Nil(t, res)
}

func TestAddImpersonation(t *testing.T) {
	res := addImpersonation(nil, nil)
	assert.Nil(t, res)
}
