package proxyfile

import (
	"path/filepath"
)

// NewClonesStrategy creates a new FileCreationStrategy that produces clones of droxy command.
func NewClonesStrategy() FileCreationStrategy {
	return &ClonesStrategy{
		copyFileFunction: copyFile,
	}
}

//ClonesStrategy contains the implementation of creating clones of droxy executable.
type ClonesStrategy struct {
	copyFileFunction copyFileFunctionDef
}

type copyFileFunctionDef func(string, string) error

//CreateProxyFile creates a clone of the given commandBinaryFilePath to commandNameFilePath.
func (s *ClonesStrategy) CreateProxyFile(commandBinaryFilePath, commandNameFilePath string) error {
	cleanSrc := filepath.Clean(commandBinaryFilePath)
	cleanDst := filepath.Clean(commandNameFilePath)

	if cleanSrc == cleanDst {
		return nil
	}

	return s.copyFileFunction(cleanSrc, cleanDst)
}
