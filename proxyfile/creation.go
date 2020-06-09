package proxyfile

import (
	"os"
	"path/filepath"
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

// NewClonesStrategy creates a new FileCreationStrategy that produces clones of droxy command.
func NewClonesStrategy() ClonesStrategy {
	return ClonesStrategy{
		FsLinkCreator: FsLinkCreator{strategy: copyFile},
	}
}

//ClonesStrategy contains the implementation of creating clones of droxy executable.
type ClonesStrategy struct {
	FsLinkCreator
}

//CreateProxyFile creates a clone of the given commandBinaryFilePath to commandNameFilePath.
func (s ClonesStrategy) CreateProxyFile(commandBinaryFilePath, commandNameFilePath string) error {
	cleanSrc := filepath.Clean(commandBinaryFilePath)
	cleanDst := filepath.Clean(commandNameFilePath)

	if cleanSrc == cleanDst {
		return nil
	}

	return s.strategy(cleanSrc, cleanDst)
}
