package config

import (
	"fmt"
)

// Configuration is the data model for a configuration file
type Configuration ConfigurationDefinition

// ConfigurationDefinition defines the fields/types of the configuration file
type ConfigurationDefinition struct {
	Command        []CommandDefinition
	Version        string
	configFilePath string
	Logging        bool
}

// FindCommandByName finds a command by the given name
func (c *Configuration) FindCommandByName(commandName string) (*CommandDefinition, error) {
	for _, command := range c.Command {
		if configCommandName, ok := command.GetName(); ok {
			if configCommandName == commandName {
				return c.resolveConfig(&command)
			}
		}
	}

	return nil, fmt.Errorf("command not defined: '%s'", commandName)
}

// SetConfigurationFilePath sets the filepath the configuration was load from. this is for debugging purpose.
func (c *Configuration) SetConfigurationFilePath(configFilePath string) {
	c.configFilePath = configFilePath
}

// GetConfigurationFilePath returns the path the configuration was load from. this is for debugging purpose.
func (c *Configuration) GetConfigurationFilePath() string {
	return c.configFilePath
}

func (c *Configuration) resolveConfig(command *CommandDefinition) (*CommandDefinition, error) {

	if !command.HasTemplate() {
		return command, nil
	}

	templateDefinition, err := c.FindCommandByName(*command.Template)
	if err != nil {
		return nil, fmt.Errorf("could not find template '%s' to resolve config of '%s'", *command.Template, *command.Name)
	}

	return mergeCommand(templateDefinition, command), nil
}
