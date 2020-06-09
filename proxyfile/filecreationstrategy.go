package proxyfile

// FileCreationStrategy defines the interface for creation of a droxy commands in filesystem.
type FileCreationStrategy interface {
	CreateProxyFile(string, string) error
}
