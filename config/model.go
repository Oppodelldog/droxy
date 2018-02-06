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

// SetConfigurationFilePath sets the filepath the configuration was load from. this is for debugging purpose.
func (c *Configuration) SetConfigurationFilePath(configFilePath string) {
	c.configFilePath = configFilePath
}

// GetConfigurationFilePath returns the path the configuration was load from. this is for debugging purpose.
func (c *Configuration) GetConfigurationFilePath() string {
	return c.configFilePath
}

// CommandDefinition defines the configuration fields for a docker-proxy-command
type CommandDefinition struct {
	IsTemplate      *bool
	Template        *string
	EntryPoint      *string
	Name            *string
	Image           *string
	Network         *string
	IsInteractive   *bool
	Volumes         *[]string
	EnvVars         *[]string
	Ports           *[]string
	AddGroups       *bool
	Impersonate     *bool
	WorkDir         *string
	RemoveContainer *bool
}

// HasTemplate indicates if the command definition has a template set
func (c *CommandDefinition) HasTemplate() bool { return c.Template != nil && *c.Template != "" }

// HasPropertyTemplate indicates if the command definition has the property Template set
func (c *CommandDefinition) HasPropertyTemplate() bool { return c.Template != nil }

// HasPropertyIsTemplate indicates if the command definition has the property IsTemplate
func (c *CommandDefinition) HasPropertyIsTemplate() bool { return c.IsTemplate != nil }

// HasPropertyEntryPoint indicates if the command definition has the property EntryPoint set
func (c *CommandDefinition) HasPropertyEntryPoint() bool { return c.EntryPoint != nil }

// HasPropertyName indicates if the command definition has the property Name set
func (c *CommandDefinition) HasPropertyName() bool { return c.Name != nil }

// HasPropertyImage indicates if the command definition has the property Image set
func (c *CommandDefinition) HasPropertyImage() bool { return c.Image != nil }

// HasPropertyVolumes indicates if the command definition has the property Volumes set
func (c *CommandDefinition) HasPropertyVolumes() bool { return c.Volumes != nil }

// HasPropertyEnvVars indicates if the command definition has the property EnvVars set
func (c *CommandDefinition) HasPropertyEnvVars() bool { return c.EnvVars != nil }

// HasPropertyPorts indicates if the command definition has the property ports set
func (c *CommandDefinition) HasPropertyPorts() bool { return c.Ports != nil }

// HasPropertyAddGroups indicates if the command definition has the property AddGroups set
func (c *CommandDefinition) HasPropertyAddGroups() bool { return c.AddGroups != nil }

// HasPropertyImpersonate indicates if the command definition has the property Impersonate set
func (c *CommandDefinition) HasPropertyImpersonate() bool { return c.Impersonate != nil }

// HasPropertyWorkDir indicates if the command definition has the property WorkDir set
func (c *CommandDefinition) HasPropertyWorkDir() bool { return c.WorkDir != nil }

// HasPropertyRemoveContainer indicates if the command definition has the property RemoveContainer set
func (c *CommandDefinition) HasPropertyRemoveContainer() bool { return c.RemoveContainer != nil }

// HasPropertyNetwork indicates if the command definition has the property Network set
func (c *CommandDefinition) HasPropertyNetwork() bool { return c.Network != nil }

// HasPropertyIsInteractive indicates if the command definition has the property IsInteractive set
func (c *CommandDefinition) HasPropertyIsInteractive() bool { return c.IsInteractive != nil }

// FindCommandByName finds a command by the given name
func (c *Configuration) FindCommandByName(commandName string) (*CommandDefinition, error) {
	for _, command := range c.Command {
		if !command.HasPropertyName() {
			continue
		}

		if *command.Name == commandName {
			return c.resolveConfig(&command)
		}
	}

	return nil, fmt.Errorf("command not defined: '%s'", commandName)
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
