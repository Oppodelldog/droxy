package proxyfile

import (
	"os"
)

func NewHardlinkStrategy() *HardlinkStrategy {
	return &HardlinkStrategy{}
}

type HardlinkStrategy struct{}

func (s *HardlinkStrategy) CreateProxyFile(commandBinaryFilePath, commandNameFilePath string) error {
	return os.Link(commandBinaryFilePath, commandNameFilePath)
}
