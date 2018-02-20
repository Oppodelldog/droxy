package proxyfile

import (
	"os"
)

// NewHardlinkStrategy creates a new FileCreationStrategy that produces hardlinks of droxy command
func NewHardlinkStrategy() FileCreationStrategy {
	return &HardlinkStrategy{}
}

//HardlinkStrategy contains the implementation of creating a hardlink to droxy execuable
type HardlinkStrategy struct{}

//CreateProxyFile creates a hardlink from commandNameFilePath to commandBinaryFilePath
func (s *HardlinkStrategy) CreateProxyFile(commandBinaryFilePath, commandNameFilePath string) error {
	return os.Link(commandBinaryFilePath, commandNameFilePath)
}
