package proxyfile

import (
	"os"
)

// NewHardlinkStrategy creates a new FileCreationStrategy that produces hardlinks of droxy command.
func NewHardlinkStrategy() FileCreationStrategy {
	return &HardlinkStrategy{
		hardLinkFunction: os.Link,
	}
}

//HardlinkStrategy contains the implementation of creating a hardlink to droxy executable.
type HardlinkStrategy struct {
	hardLinkFunction hardlinkFunctionDef
}

type hardlinkFunctionDef func(string, string) error

//CreateProxyFile creates a hardlink from commandNameFilePath to commandBinaryFilePath.
func (s *HardlinkStrategy) CreateProxyFile(commandBinaryFilePath, commandNameFilePath string) error {
	return s.hardLinkFunction(commandBinaryFilePath, commandNameFilePath)
}

// NewSymlinkStrategy creates a new FileCreationStrategy that produces symlinks of droxy command.
func NewSymlinkStrategy() FileCreationStrategy {
	return &SymlinkStrategy{
		symlinkFunction: os.Symlink,
	}
}

//SymlinkStrategy contains the implementation of creating a symlink to droxy executable.
type SymlinkStrategy struct {
	symlinkFunction symlinkFunctionDef
}

type symlinkFunctionDef func(string, string) error

//CreateProxyFile creates a symlink from commandNameFilePath to commandBinaryFilePath.
func (s *SymlinkStrategy) CreateProxyFile(commandBinaryFilePath, commandNameFilePath string) error {
	return s.symlinkFunction(commandBinaryFilePath, commandNameFilePath)
}
