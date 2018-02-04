package proxyfile

import (
	"os"
)

// NewSymlinkStrategy creates a new FileCreationStrategy that produces symlinks of docker-proxy command
func NewSymlinkStrategy() FileCreationStrategy {
	return &SymlinkStrategy{}
}

//SymlinkStrategy contains the implementation of creating a symlink to docker-proxy execuable
type SymlinkStrategy struct{}

//CreateProxyFile creates a symlink from commandNameFilePath to commandBinaryFilePath
func (s *SymlinkStrategy) CreateProxyFile(commandBinaryFilePath, commandNameFilePath string) error {
	return os.Symlink(commandBinaryFilePath, commandNameFilePath)
}
