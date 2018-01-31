package config

import (
	"fmt"
)

type Configuration Commands

type Commands struct {
	Command []CommandDefinition
	Version string
}

type CommandDefinition struct {
	EntryPoint      string
	Name            string
	Image           string
	Volumes         []string
	EnvVars         []string
	AddGroups       bool
	Impersonate     bool
	WorkDir         string
	RemoveContainer bool
}

func (c *Configuration) FindCommandByName(commandName string) (*CommandDefinition, error) {
	for _, command := range c.Command {
		if command.Name == commandName {
			return &command, nil
		}
	}

	return nil, fmt.Errorf("command not defined: '%s'", commandName)
}
