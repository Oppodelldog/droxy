package helper

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestStringInSlice_StringIsInSlice(t *testing.T) {
	s := "abc"
	slice := []string{"iop", "zui", "abc"}
	res := StringInSlice(s, slice)

	assert.True(t, res)
}

func TestStringInSlice_StringIsNotInSlice(t *testing.T) {
	s := "abc"
	slice := []string{"iop", "zui", "abf"}
	res := StringInSlice(s, slice)

	assert.False(t, res)
}
