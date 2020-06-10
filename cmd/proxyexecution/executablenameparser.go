package proxyexecution

import (
	"github.com/Oppodelldog/droxy/crossplatform"
)

type executableNameParser struct{}

func newExecutableNameParser() executableNameParser {
	return executableNameParser{}
}

//ParseCommandNameFromCommandLine parsed the command name of the currently executed binary from cli arguments.
func (p executableNameParser) ParseCommandNameFromCommandLine() string {
	return crossplatform.ParseCommandNameFromCommandLine()
}
