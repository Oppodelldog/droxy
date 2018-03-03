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

func TestNewHardlinkStrategy_callsConfiguredSystemFunction(t *testing.T) {
	mock := fileCreationFunctionMock{}
	strategy := NewHardlinkStrategy()
	strategy.(*HardlinkStrategy).hardLinkFunction = mock.FileCreationFunc

	expectedSrc := "A"
	expectedDst := "B"
	strategy.CreateProxyFile(expectedSrc, expectedDst)

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
