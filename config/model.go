package config

import (
	"fmt"
)

type Configuration Commands

type Commands struct {
	Command        []CommandDefinition
	Version        string
	configFilePath string
	Logging        bool
}

func (c *Configuration) SetConfigurationFilePath(configFilePath string) {
	c.configFilePath = configFilePath
}
func (c *Configuration) GetConfigurationFilePath() string {
	return c.configFilePath
}

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
	AddGroups       *bool
	Impersonate     *bool
	WorkDir         *string
	RemoveContainer *bool
}

func (c *CommandDefinition) HasTemplate() bool                { return c.Template != nil && *c.Template != "" }
func (c *CommandDefinition) HasPropertyTemplate() bool        { return c.Template != nil }
func (c *CommandDefinition) HasPropertyIsTemplate() bool      { return c.IsTemplate != nil }
func (c *CommandDefinition) HasPropertyEntryPoint() bool      { return c.EntryPoint != nil }
func (c *CommandDefinition) HasPropertyName() bool            { return c.Name != nil }
func (c *CommandDefinition) HasPropertyImage() bool           { return c.Image != nil }
func (c *CommandDefinition) HasPropertyVolumes() bool         { return c.Volumes != nil }
func (c *CommandDefinition) HasPropertyEnvVars() bool         { return c.EnvVars != nil }
func (c *CommandDefinition) HasPropertyAddGroups() bool       { return c.AddGroups != nil }
func (c *CommandDefinition) HasPropertyImpersonate() bool     { return c.Impersonate != nil }
func (c *CommandDefinition) HasPropertyWorkDir() bool         { return c.WorkDir != nil }
func (c *CommandDefinition) HasPropertyRemoveContainer() bool { return c.RemoveContainer != nil }
func (c *CommandDefinition) HasPropertyNetwork() bool         { return c.Network != nil }
func (c *CommandDefinition) HasPropertyIsInteractive() bool   { return c.IsInteractive != nil }

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
		return nil, fmt.Errorf("could not find template '%s' to resolve config of '%s'", command.Template, command.Name)
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
