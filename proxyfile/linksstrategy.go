package proxyfile

import (
	"os"
)

// NewHardlinkStrategy creates a new FsLinkCreator that produces hardlinks of droxy command.
func NewHardlinkStrategy() FsLinkCreator {
	return FsLinkCreator{
		strategy: os.Link,
	}
}

// NewSymlinkStrategy creates a new FsLinkCreator that produces symlinks of droxy command.
func NewSymlinkStrategy() FsLinkCreator {
	return FsLinkCreator{
		strategy: os.Symlink,
	}
}

type FsLinkCreator struct {
	strategy func(string, string) error
}

func (c FsLinkCreator) CreateProxyFile(commandBinaryFilePath, commandNameFilePath string) error {
	return c.strategy(commandBinaryFilePath, commandNameFilePath)
}
