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
		} else {
			continue
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

	return c.mergeCommand(templateDefinition, command), nil

}

func (c *Configuration) mergeCommand(baseCommand *CommandDefinition, overlayCommand *CommandDefinition) *CommandDefinition {
	mergedCommand := new(CommandDefinition)

	mergedCommand.Image = resolvePropertyString(baseCommand.Name, overlayCommand.Name)
	mergedCommand.EntryPoint = resolvePropertyString(baseCommand.EntryPoint, overlayCommand.EntryPoint)
	mergedCommand.Image = resolvePropertyString(baseCommand.Image, overlayCommand.Image)
	mergedCommand.WorkDir = resolvePropertyString(baseCommand.WorkDir, overlayCommand.WorkDir)
	mergedCommand.Network = resolvePropertyString(baseCommand.Network, overlayCommand.Network)

	mergedCommand.AddGroups = resolvePropertyBool(baseCommand.AddGroups, overlayCommand.AddGroups)
	mergedCommand.RemoveContainer = resolvePropertyBool(baseCommand.RemoveContainer, overlayCommand.RemoveContainer)
	mergedCommand.Impersonate = resolvePropertyBool(baseCommand.Impersonate, overlayCommand.Impersonate)
	mergedCommand.IsInteractive = resolvePropertyBool(baseCommand.IsInteractive, overlayCommand.IsInteractive)

	mergedCommand.Volumes = resolvePropertyStringArray(baseCommand.Volumes, overlayCommand.Volumes)
	mergedCommand.EnvVars = resolvePropertyStringArray(baseCommand.EnvVars, overlayCommand.EnvVars)
	mergedCommand.Ports = resolvePropertyStringArray(baseCommand.Ports, overlayCommand.Ports)

	return mergedCommand
}

func resolvePropertyBool(bBase *bool, bOverlay *bool) *bool {
	var b bool

	if bBase != nil {
		b = *bBase
	}

	if bOverlay != nil {
		b = *bOverlay
	}

	return &b
}

func resolvePropertyString(sBase *string, sOverlay *string) *string {
	var s string

	if sBase != nil {
		s = *sBase
	}

	if sOverlay != nil {
		s = *sOverlay
	}

	return &s
}

func resolvePropertyStringArray(sBase *[]string, sOverlay *[]string) *[]string {
	var s []string

	if sBase != nil {
		s = *sBase
	}

	if sOverlay != nil {
		s = *sOverlay
	}

	return &s
}
