package config

import (
	"errors"
	"fmt"
	"runtime"
)

var errCommandNotDefined = errors.New("command not defined")
var errCouldNotFindTemplate = errors.New("could not find template")

// Configuration defines the fields/types of the configuration file.
type Configuration struct {
	Command        []CommandDefinition
	Version        string
	ConfigFilePath string
	Logging        bool
	osNameMatcher  func(string) bool
}

// FindCommandByName finds a command by the given name.
func (c Configuration) FindCommandByName(commandName string) (CommandDefinition, error) {
	var commandDef CommandDefinition
	var found bool

	for _, command := range c.Command {
		if configCommandName, ok := command.GetName(); ok {
			os, _ := command.GetOS()
			if configCommandName == commandName && c.matchesOS(os) {
				cd, err := c.resolveConfig(command)
				if err != nil {
					return CommandDefinition{}, fmt.Errorf("error finding command '%s': %v", commandName, err)
				}

				commandDef = cd
				found = true
			}
		}
	}

	if found {
		return commandDef, nil
	}

	return CommandDefinition{}, fmt.Errorf("%w: '%s'", errCommandNotDefined, commandName)
}

func (c Configuration) matchesOS(osName string) bool {
	if osName == "" {
		return true
	}

	return c.osNameMatcher(osName)
}

// GetConfigurationFilePath returns the path the configuration was load from. this is for debugging purpose.
func (c Configuration) GetConfigurationFilePath() string {
	return c.ConfigFilePath
}

func (c Configuration) resolveConfig(command CommandDefinition) (CommandDefinition, error) {
	if !command.HasTemplate() {
		return command, nil
	}

	templateDefinition, err := c.FindCommandByName(*command.Template)
	if err != nil {
		return CommandDefinition{},
			fmt.Errorf(
				"%w '%s' to resolve config of '%s'",
				errCouldNotFindTemplate,
				*command.Template,
				*command.Name,
			)
	}

	return mergeCommand(templateDefinition, command), nil
}

func defaultOSNameMatcher(osName string) bool {
	return runtime.GOOS == osName
}
