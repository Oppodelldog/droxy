package proxyfile

import (
	"errors"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkStrategies_configuresTheAppropriateSystemFunction(t *testing.T) {
	testCases := map[string]struct {
		linkCreator      FsLinkCreator
		expectedFunction func(string, string) error
	}{
		"hardlink": {
			linkCreator:      NewHardlinkStrategy(),
			expectedFunction: os.Link,
		},
		"symlink": {
			linkCreator:      NewSymlinkStrategy(),
			expectedFunction: os.Symlink,
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			want := reflect.ValueOf(testCase.expectedFunction).Pointer()
			got := reflect.ValueOf(testCase.linkCreator.strategy).Pointer()

			if want != got {
				t.Fatalf("want: %v, got: %v", want, got)
			}
		})
	}
}

func TestFsLinkCreator_callsConfiguredSystemFunction(t *testing.T) {
	mock := &linkMock{}
	linkCreator := FsLinkCreator{
		strategy: mock.Call,
	}

	expectedSrc := "A"
	expectedDst := "B"

	err := linkCreator.CreateProxyFile(expectedSrc, expectedDst)
	if err != nil {
		t.Fatalf("Did not expect CreateProxyFile to return an error, but got: %v", err)
	}

	assert.Equal(t, expectedSrc, mock.arg1)
	assert.Equal(t, expectedDst, mock.arg2)
	assert.Equal(t, 1, mock.calls)
}

func TestFsLinkCreator_returnsErrorIfSystemFunctionReturnsError(t *testing.T) {
	expectedError := errors.New("error from configured system function")

	mock := &linkMock{err: expectedError}
	linkCreator := FsLinkCreator{
		strategy: mock.Call,
	}

	err := linkCreator.CreateProxyFile("A", "B")

	assert.Equal(t, expectedError, err)
}

type linkMock struct {
	arg1  string
	arg2  string
	calls int
	err   error
}

func (m *linkMock) Call(p1, p2 string) error {
	m.arg1 = p1
	m.arg2 = p2
	m.calls++

	return m.err
}
