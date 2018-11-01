package proxyfile

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const windowsCommandExtension = ".exe"

// GetCommandName returns the filename of the droxy binary
func GetCommandName() string {
	return fmt.Sprintf("%s%s", commandFileName, windowsCommandExtension)
}

// ParseCommandNameFromCommandLine returns the called proxy command from cli args
func ParseCommandNameFromCommandLine() string {
	return strings.Replace(filepath.Base(os.Args[0]), windowsCommandExtension, "", -1)
}

// GetCommandNameFilename returns the binary filename of the given proxy command name
func GetCommandNameFilename(commandName string) string {
	return fmt.Sprintf("%s%s", commandName, windowsCommandExtension)
}
