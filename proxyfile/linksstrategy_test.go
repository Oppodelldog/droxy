package proxyfile

import (
	"errors"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHardlinkStrategy_configuresTheAppropriateSystemFunction(t *testing.T) {
	strategy := NewHardlinkStrategy()

	strategyFunction := strategy.(*HardlinkStrategy).hardLinkFunction
	expectedFunction := os.Link

	if reflect.ValueOf(expectedFunction).Pointer() != reflect.ValueOf(strategyFunction).Pointer() {
		t.Fail()
	}
}
func TestNewSymlinkStrategy_callsTheAppropriateSystemFunction(t *testing.T) {
	strategy := NewSymlinkStrategy()

	strategyFunction := strategy.(*SymlinkStrategy).symlinkFunction
	expectedFunction := os.Symlink

	if reflect.ValueOf(expectedFunction).Pointer() != reflect.ValueOf(strategyFunction).Pointer() {
		t.Fail()
	}
}

func TestNewHardlinkStrategy_callsConfiguredSystemFunction(t *testing.T) {
	mock := fileCreationFunctionMock{}
	strategy := NewHardlinkStrategy()
	strategy.(*HardlinkStrategy).hardLinkFunction = mock.FileCreationFunc

	expectedSrc := "A"
	expectedDst := "B"

	err := strategy.CreateProxyFile(expectedSrc, expectedDst)
	if err != nil {
		t.Fatalf("Did not expect CreateProxyFile to return an error, but got: %v", err)
	}

	assert.Equal(t, expectedSrc, mock.parmSrc)
	assert.Equal(t, expectedDst, mock.parmDst)
	assert.Equal(t, 1, mock.calls)
}
func TestNewSymlinkStrategy_callsConfiguredSystemFunction(t *testing.T) {
	mock := fileCreationFunctionMock{}
	strategy := NewSymlinkStrategy()
	strategy.(*SymlinkStrategy).symlinkFunction = mock.FileCreationFunc

	expectedSrc := "A"
	expectedDst := "B"

	err := strategy.CreateProxyFile(expectedSrc, expectedDst)
	if err != nil {
		t.Fatalf("Did not expect CreateProxyFile to return an error, but got: %v", err)
	}

	assert.Equal(t, expectedSrc, mock.parmSrc)
	assert.Equal(t, expectedDst, mock.parmDst)
	assert.Equal(t, 1, mock.calls)
}

func TestNewHardlinkStrategy_returnsErrorIfSystemFunctionReturnsError(t *testing.T) {
	expectedError := errors.New("error from configured system function")

	mock := fileCreationFunctionMock{returnValue: expectedError}
	strategy := NewHardlinkStrategy()
	strategy.(*HardlinkStrategy).hardLinkFunction = mock.FileCreationFunc

	err := strategy.CreateProxyFile("A", "B")

	assert.Equal(t, expectedError, err)
}
func TestNewSymlinkStrategy_returnsErrorIfSystemFunctionReturnsError(t *testing.T) {
	expectedError := errors.New("error from configured system function")

	mock := fileCreationFunctionMock{returnValue: expectedError}
	strategy := NewSymlinkStrategy()
	strategy.(*SymlinkStrategy).symlinkFunction = mock.FileCreationFunc

	err := strategy.CreateProxyFile("A", "B")

	assert.Equal(t, expectedError, err)
}
