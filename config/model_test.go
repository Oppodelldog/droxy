package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_resolvePropertyStringArray2D_BaseSet_OverlayNotSet(t *testing.T) {
	base := &[][]string{
		{"a", "b"},
	}
	var overlay *[][]string

	result := resolvePropertyStringArray2D(base, overlay)

	expectedResult := &[][]string{
		{"a", "b"},
	}

	assert.Equal(t, expectedResult, result)
}

func Test_resolvePropertyStringArray2D_BaseNotSet_OverlaySet(t *testing.T) {
	var base *[][]string
	overlay := &[][]string{
		{"a", "b"},
	}

	result := resolvePropertyStringArray2D(base, overlay)

	expectedResult := &[][]string{
		{"a", "b"},
	}

	assert.Equal(t, expectedResult, result)
}

func Test_resolvePropertyStringArray2D_BaseAndOverlaySet(t *testing.T) {
	base := &[][]string{
		{"a", "b"},
	}
	overlay := &[][]string{
		{"c", "d"},
	}

	result := resolvePropertyStringArray2D(base, overlay)

	expectedResult := &[][]string{
		{"c", "d"},
	}

	assert.Equal(t, expectedResult, result)
}
