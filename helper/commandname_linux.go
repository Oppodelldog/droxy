package helper

import (
	"os"
	"path/filepath"
)

// GetCommandName returns the filename of the droxy binary
func GetCommandName() string {
	return commandFileName
}

// ParseCommandNameFromCommandLine returns the called proxy command from cli args
func ParseCommandNameFromCommandLine() string {
	return filepath.Base(os.Args[0])
}

// GetCommandNameFilename returns the binary filename of the given proxy command name
func GetCommandNameFilename(commandName string) string {
	return commandName
}
