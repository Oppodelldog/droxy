package helper

import (
	"fmt"
	"path/filepath"
	"os"
	"strings"
)

const windowsCommandExtension = ".exe"

func GetCommandName() string {
	return fmt.Sprintf("%s%s", commandFileName, windowsCommandExtension)
}

func ParseCommandNameFromCommandLine() string {
	return strings.Replace(filepath.Base(os.Args[0]), windowsCommandExtension, "", -1)
}

func GetCommandNameFilename(commandName string) string {
	return fmt.Sprintf("%s%s", commandName, windowsCommandExtension)
}
