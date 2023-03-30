package proxyfile

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

const writePerm = 0600

func TestCopyFile(t *testing.T) {
	testFolder := "/tmp/droxy/fileCopyTest/"

	err := os.MkdirAll(testFolder, 0776)
	if err != nil {
		t.Fatalf("Did not expect os.MkdirAll to return an error, but got: %v", err)
	}

	srcBytes := []byte("HELLO DROXY!!")
	src := path.Join(testFolder, "fileToCopy")

	err = os.WriteFile(src, srcBytes, writePerm)
	if err != nil {
		t.Fatalf("Did not expect os.WriteFile to return an error, but got: %v", err)
	}

	dst := path.Join(testFolder, "fileCopied")

	err = copyFile(src, dst)
	if err != nil {
		t.Fatalf("Did not expect copyFile to return an error, but got: %v", err)
	}

	dstBytes, err := os.ReadFile(dst)
	if err != nil {
		t.Fatalf("Did not expect os.ReadFile to return an error, but got: %v", err)
	}

	assert.Equal(t, dstBytes, srcBytes)

	err = os.RemoveAll(testFolder)
	if err != nil {
		t.Fatalf("Did not expect os.Remove to return an error, but got: %v", err)
	}
}

func TestCopyFile_srcFileDoesNotExist_expectError(t *testing.T) {
	src := "/tmp/THIS_FILE_DOES_NOT_EXIST"
	dst := ""
	err := copyFile(src, dst)

	assert.Error(t, err)
}
