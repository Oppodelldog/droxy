package proxyexecution

import (
	"github.com/Oppodelldog/droxy/proxyfile"
)

type executableNameParser struct{}

func newExecutableNameParser() executableNameParser {
	return executableNameParser{}
}

//ParseCommandNameFromCommandLine parsed the command name of the currently executed binary from cli arguments.
func (p executableNameParser) ParseCommandNameFromCommandLine() string {
	return proxyfile.ParseCommandNameFromCommandLine()
}
