package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_stringInSlice_StringIsInSlice(t *testing.T) {
	s := "abc"
	slice := []string{"iop", "zui", "abc"}
	res := stringInSlice(s, slice)

	assert.True(t, res)
}

func Test_stringInSlice_StringIsNotInSlice(t *testing.T) {
	s := "abc"
	slice := []string{"iop", "zui", "abf"}
	res := stringInSlice(s, slice)

	assert.False(t, res)
}
