package cmd

import "github.com/Oppodelldog/droxy/helper"

type executableNameParser struct{}

//NewExecutableNameParser returns a new executableNameParser
func NewExecutableNameParser() *executableNameParser {
	return &executableNameParser{}
}

//ParseCommandNameFromCommandLine parsed the command name of the currently executed binary from cli aruments
func (p *executableNameParser) ParseCommandNameFromCommandLine() string {
	return helper.ParseCommandNameFromCommandLine()
}
