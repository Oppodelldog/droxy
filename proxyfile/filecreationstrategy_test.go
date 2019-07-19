package proxyfile

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileCreation_AllStrategies(t *testing.T) {

	testCases := []struct {
		strategy FileCreationStrategy
	}{
		{NewClonesStrategy()},
		{NewHardlinkStrategy()},
		{NewSymlinkStrategy()},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%T", tc.strategy), func(t *testing.T) {
			strategy := tc.strategy

			testFolderPath := "/tmp/droxy/strategy"
			err := os.MkdirAll(testFolderPath, 0776)
			if err != nil {
				t.Fatalf("did not expect os.MkdirAll to return an error, but got: %v", err)
			}
			defer func() {
				err := os.RemoveAll(testFolderPath)
				if err != nil {
					t.Fatalf("Did not expect os.RemoveAll to return an error, but got: %v", err)
				}
			}()

			src := path.Join(testFolderPath, "testFileSrc")

			srcBytes := []byte{1, 2, 3, 4, 5}
			err = ioutil.WriteFile(src, srcBytes, 0776)
			failOnError(t, err, "did not expect ioutil.WriteAndSyncFile to return an error, but got: %v")

			target := path.Join(testFolderPath, "testFileTarget")

			err = strategy.CreateProxyFile(src, target)
			failOnError(t, err, "did not expect CreateProxyFile to return an error, but got: %v")

			targetBytes, err := ioutil.ReadFile(target)
			failOnError(t, err, "did not expect ReadFile to return an error, but got: %v")

			assert.Equal(t, srcBytes, targetBytes)
		},
		)
	}
}

func failOnError(t *testing.T, err error, message string) {
	if err != nil {
		t.Fatalf(message, err)
	}
}
