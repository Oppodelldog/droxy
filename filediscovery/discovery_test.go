package filediscovery

import (
	"os"
	"path"
	"testing"

	"bytes"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert.Implements(t, new(FileDiscovery), New([]FileLocationProvider{}))
}

func TestFileDiscovery_Discover_callsFileLocationProviders(t *testing.T) {

	mock1, provider1 := newFileLocationProviderMock()
	mock2, provider2 := newFileLocationProviderMock()

	providers := []FileLocationProvider{
		provider1,
		provider2,
	}

	discovery := New(providers)
	discovery.Discover("")

	assert.True(t, mock1.WasCalled())
	assert.True(t, mock2.WasCalled())
}

func TestFileDiscovery_Discover_callsFileLocationProvidersWithFilename(t *testing.T) {

	mock, provider := newFileLocationProviderMock()
	providers := []FileLocationProvider{provider}

	discovery := New(providers)
	testFilename := "test-file"
	discovery.Discover(testFilename)

	assert.Equal(t, testFilename, mock.GetCalledFilenameParameter())
}

func TestFileDiscovery_Discover_providerErrorsAreAppendedToError(t *testing.T) {

	errorMessage := "stub-error"
	errStub := errors.New(errorMessage)
	mock := &fileLocationProviderMock{}
	provider := mock.GetFunc("", errStub)

	providers := []FileLocationProvider{provider}

	discovery := New(providers)
	testFilename := "test-file"
	_, err := discovery.Discover(testFilename)

	assert.Contains(t, err.Error(), errorMessage)
}

func TestFileDiscovery_Discover_ifFileNotFoundReturnsError(t *testing.T) {

	mock := &fileLocationProviderMock{}
	provider := mock.GetFunc("", nil)

	providers := []FileLocationProvider{provider}

	discovery := New(providers)
	testFilename := "test-file"
	_, err := discovery.Discover(testFilename)

	expectedError := bytes.NewBufferString("could not find config file at ''")
	expectedError.WriteString("\n")

	assert.Equal(t, expectedError.String(), err.Error())
}

func TestFileDiscovery_Discover_ifFileWasFoundReturnsFilePath(t *testing.T) {

	testFilename := "test-file"
	testFilePath := path.Join(os.TempDir(), testFilename)
	f, err := os.Create(testFilePath)
	if err != nil {
		t.Fatalf("did not expect os.Create to return an error, but got: %v", err)
	}
	f.Close()

	tempDirProvider := func(fileName string) (string, error) {
		return path.Join(os.TempDir(), testFilename), nil
	}

	providers := []FileLocationProvider{tempDirProvider}

	discovery := New(providers)

	result, err := discovery.Discover(testFilename)
	if err != nil {
		t.Fatalf("did not expect discovery.Discover to return an error, but got: %v", err)
	}

	assert.Equal(t, testFilePath, result)

	err = os.Remove(testFilePath)
	if err != nil {
		t.Fatalf("did not expect os.Remove to return an error, but got: %v", err)
	}
}
