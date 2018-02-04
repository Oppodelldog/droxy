package proxyfile

import (
	"os"
)

// NewHardlinkStrategy creates a new FileCreationStrategy that produces hardlinks of docker-proxy command
func NewHardlinkStrategy() FileCreationStrategy {
	return &HardlinkStrategy{}
}

//HardlinkStrategy contains the implementation of creating a hardlink to docker-proxy execuable
type HardlinkStrategy struct{}

//CreateProxyFile creates a hardlink from commandNameFilePath to commandBinaryFilePath
func (s *HardlinkStrategy) CreateProxyFile(commandBinaryFilePath, commandNameFilePath string) error {
	return os.Link(commandBinaryFilePath, commandNameFilePath)
}
