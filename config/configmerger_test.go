package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResolvePropertyString_nilInputs_nilOutput(t *testing.T) {
	assert.Nil(t, resolvePropertyString(nil, nil))
}

func TestResolvePropertyString_parm1Defined_outputAddressIsDifferentToInput(t *testing.T) {
	s := "A"
	saddr := &s
	res := resolvePropertyString(&s, nil)

	assert.True(t, saddr != res)
}

func TestResolvePropertyString_parm2Defined_outputAddressIsDifferentToInput(t *testing.T) {
	s := "A"
	saddr := &s
	res := resolvePropertyString(nil, &s)

	assert.True(t, saddr != res)
}

func TestResolvePropertyString_allParamsDefined_outputLatestParamValue(t *testing.T) {
	s1 := "A"
	s2 := "B"

	res := resolvePropertyString(&s1, &s2)

	assert.Equal(t, s2, *res)
}

func TestResolvePropertyBool_nilInputs_nilOutput(t *testing.T) {
	assert.Nil(t, resolvePropertyBool(nil, nil))
}

func TestResolvePropertyBool_parm1Defined_outputAddressIsDifferentToInput(t *testing.T) {
	b := true
	baddr := &b
	res := resolvePropertyBool(&b, nil)

	assert.True(t, baddr != res)
}

func TestResolvePropertyBool_parm2Defined_outputAddressIsDifferentToInput(t *testing.T) {
	b := true
	baddr := &b
	res := resolvePropertyBool(nil, &b)

	assert.True(t, baddr != res)
}

func TestResolvePropertyBool_allParamsDefined_outputLatestParamValue(t *testing.T) {
	b1 := true
	b2 := false

	res := resolvePropertyBool(&b1, &b2)

	assert.Equal(t, b2, *res)
}

func TestResolvePropertyStringArray_nilInputs_nilOutput(t *testing.T) {
	assert.Nil(t, resolvePropertyStringArray(false, nil, nil))
}

func TestResolvePropertyStringArray_parm1Defined_outputAddressIsDifferentToInput(t *testing.T) {
	var arr []string
	arrAddr := &arr
	res := resolvePropertyStringArray(false, &arr, nil)

	assert.True(t, arrAddr != res)
}

func TestResolvePropertyStringArray_parm2Defined_outputAddressIsDifferentToInput(t *testing.T) {
	var arr []string
	arrAddr := &arr
	res := resolvePropertyStringArray(false, nil, &arr)

	assert.True(t, arrAddr != res)
}

func TestResolvePropertyStringArray_allParamsDefined_outputLatestParamValue(t *testing.T) {
	var (
		arr1 []string
		arr2 []string
	)

	res := resolvePropertyStringArray(false, &arr1, &arr2)

	assert.Equal(t, arr2, *res)
}

func TestResolvePropertyStringArray_mergeIsTrue_nilInputs_nilOutput(t *testing.T) {
	assert.Nil(t, resolvePropertyStringArray(true, nil, nil))
}

func TestResolvePropertyStringArray_mergeIsTrue_parm1Defined_outputAddressIsDifferentToInput(t *testing.T) {
	var arr []string
	arrAddr := &arr
	res := resolvePropertyStringArray(true, &arr, nil)

	assert.True(t, arrAddr != res)
}

func TestResolvePropertyStringArray_mergeIsTrue_parm2Defined_outputAddressIsDifferentToInput(t *testing.T) {
	var arr []string
	arrAddr := &arr
	res := resolvePropertyStringArray(true, nil, &arr)

	assert.True(t, arrAddr != res)
}

func TestResolvePropertyStringArray_mergeIsTrue_allParamsDefined_outputLatestParamValue(t *testing.T) {
	arr1 := []string{"A"}
	arr2 := []string{"B"}

	res := resolvePropertyStringArray(true, &arr1, &arr2)

	assert.Equal(t, append(arr1, arr2...), *res)
}

func TestResolvePropertyStringArray2D_nilInputs_nilOutput(t *testing.T) {
	assert.Nil(t, resolvePropertyStringArray2D(nil, nil))
}

func TestResolvePropertyStringArray2D_parm1Defined_outputAddressIsDifferentToInput(t *testing.T) {
	var arr [][]string
	arrAddr := &arr
	res := resolvePropertyStringArray2D(&arr, nil)

	assert.True(t, arrAddr != res)
}

func TestResolvePropertyStringArray2D_parm2Defined_outputAddressIsDifferentToInput(t *testing.T) {
	var arr [][]string
	arrAddr := &arr
	res := resolvePropertyStringArray2D(nil, &arr)

	assert.True(t, arrAddr != res)
}

func TestResolvePropertyStringArray2D_allParamsDefined_outputLatestParamValue(t *testing.T) {
	var (
		arr1 [][]string
		arr2 [][]string
	)

	res := resolvePropertyStringArray2D(&arr1, &arr2)

	assert.Equal(t, arr2, *res)
}
