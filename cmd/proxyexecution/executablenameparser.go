package proxyexecution

import (
	"github.com/Oppodelldog/droxy/proxyfile"
)

type executableNameParser struct{}

//NewExecutableNameParser returns a new executableNameParser.
func NewExecutableNameParser() ExecutableNameParser {
	return &executableNameParser{}
}

//ParseCommandNameFromCommandLine parsed the command name of the currently executed binary from cli arguments.
func (p *executableNameParser) ParseCommandNameFromCommandLine() string {
	return proxyfile.ParseCommandNameFromCommandLine()
}
