package proxyfile

import (
	"errors"
	"reflect"
	"testing"

	"github.com/Oppodelldog/droxy/helper"
	"github.com/stretchr/testify/assert"
)

func TestNewClonesStrategy_configuresTheAppropriateSystemFunction(t *testing.T) {
	strategy := NewClonesStrategy()

	strategyFunction := strategy.(*ClonesStrategy).copyFileFunction
	expectedFunction := helper.CopyFile

	if reflect.ValueOf(expectedFunction).Pointer() != reflect.ValueOf(strategyFunction).Pointer() {
		t.Fail()
	}
}

func TestNewClonesStrategy_callsConfiguredSystemFunction(t *testing.T) {
	mock := fileCreationFunctionMock{}
	strategy := NewClonesStrategy()
	strategy.(*ClonesStrategy).copyFileFunction = mock.FileCreationFunc

	expectedSrc := "A"
	expectedDst := "B"
	strategy.CreateProxyFile(expectedSrc, expectedDst)

	assert.Equal(t, expectedSrc, mock.parmSrc)
	assert.Equal(t, expectedDst, mock.parmDst)
	assert.Equal(t, 1, mock.calls)
}

func TestNewClonesStrategy_returnsErrorIfSystemFunctionReturnsError(t *testing.T) {
	expectedError := errors.New("error from configured system function")

	mock := fileCreationFunctionMock{returnValue: expectedError}
	strategy := NewClonesStrategy()
	strategy.(*ClonesStrategy).copyFileFunction = mock.FileCreationFunc

	err := strategy.CreateProxyFile("A", "B")

	assert.Equal(t, expectedError, err)
}

func TestNewClonesStrategy_returnsNilIfFilePathsAreSame(t *testing.T) {
	strategy := NewClonesStrategy()
	err := strategy.CreateProxyFile("A", "A")

	assert.Nil(t, err)
}
