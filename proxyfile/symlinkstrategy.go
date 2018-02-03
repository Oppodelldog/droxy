package proxyfile

import (
	"os"
)

func NewSymlinkStrategy() *SymlinkStrategy {
	return &SymlinkStrategy{}
}

type SymlinkStrategy struct{}

func (s *SymlinkStrategy) CreateProxyFile(commandBinaryFilePath, commandNameFilePath string) error {
	return os.Symlink(commandBinaryFilePath, commandNameFilePath)
}
