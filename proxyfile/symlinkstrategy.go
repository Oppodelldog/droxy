package proxyfile

import (
	"os"
)

// NewSymlinkStrategy creates a new FileCreationStrategy that produces symlinks of droxy command
func NewSymlinkStrategy() FileCreationStrategy {
	return &SymlinkStrategy{}
}

//SymlinkStrategy contains the implementation of creating a symlink to droxy execuable
type SymlinkStrategy struct{}

//CreateProxyFile creates a symlink from commandNameFilePath to commandBinaryFilePath
func (s *SymlinkStrategy) CreateProxyFile(commandBinaryFilePath, commandNameFilePath string) error {
	return os.Symlink(commandBinaryFilePath, commandNameFilePath)
}
