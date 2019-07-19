package proxyfile

import (
	"os"
)

// NewSymlinkStrategy creates a new FileCreationStrategy that produces symlinks of droxy command
func NewSymlinkStrategy() FileCreationStrategy {
	return &SymlinkStrategy{
		symlinkFunction: os.Symlink,
	}
}

//SymlinkStrategy contains the implementation of creating a symlink to droxy executable
type SymlinkStrategy struct {
	symlinkFunction symlinkFunctionDef
}

type symlinkFunctionDef func(string, string) error

//CreateProxyFile creates a symlink from commandNameFilePath to commandBinaryFilePath
func (s *SymlinkStrategy) CreateProxyFile(commandBinaryFilePath, commandNameFilePath string) error {
	return s.symlinkFunction(commandBinaryFilePath, commandNameFilePath)
}
