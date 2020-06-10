package config

import "strings"

// CommandDefinition gives public access to the fields by accessor functions.
type CommandDefinition struct {
	RequireEnvVars      *bool
	IsTemplate          *bool
	Template            *string
	EntryPoint          *string
	Command             *string
	Name                *string
	UniqueNames         *bool
	Image               *string
	Network             *string
	EnvFile             *string
	IP                  *string
	IsInteractive       *bool
	IsDetached          *bool
	IsDaemon            *bool // deprecated
	Volumes             *[]string
	Links               *[]string
	EnvVars             *[]string
	Ports               *[]string
	PortsFromParams     *[]string
	MergeTemplateArrays *[]string
	AddGroups           *bool
	Impersonate         *bool
	WorkDir             *string
	AutoMountWorkDir    *bool
	RemoveContainer     *bool
	ReplaceArgs         *[][]string
	AdditionalArgs      *[]string
}

// GetEntryPoint returns entrypoint and an boolean indicating if value is set.
func (c CommandDefinition) GetEntryPoint() (string, bool) {
	return getString(c.EntryPoint)
}

// getCommand returns value of Command (CMD) and an boolean indicating if value is set.
func (c CommandDefinition) GetCommand() (string, bool) {
	return getString(c.Command)
}

// GetName returns value of Name and an boolean indicating if value is set.
func (c CommandDefinition) GetName() (string, bool) {
	return getString(c.Name)
}

// GetImage returns value of Image and an boolean indicating if value is set.
func (c CommandDefinition) GetImage() (string, bool) {
	return getString(c.Image)
}

// GetNetwork returns value of Network and an boolean indicating if value is set.
func (c CommandDefinition) GetNetwork() (string, bool) {
	return getString(c.Network)
}

// GetEnvFile returns value of EnvFile and an boolean indicating if value is set.
func (c CommandDefinition) GetEnvFile() (string, bool) {
	return getString(c.EnvFile)
}

// GetIP returns value of Ip and an boolean indicating if value is set.
func (c CommandDefinition) GetIP() (string, bool) {
	return getString(c.IP)
}

// GetWorkDir returns value of Impersonate and an boolean indicating if value is set.
func (c CommandDefinition) GetWorkDir() (string, bool) {
	return getString(c.WorkDir)
}

// GetTemplate returns value of Template and an boolean indicating if value is set.
func (c CommandDefinition) GetTemplate() (string, bool) {
	return getString(c.Template)
}

// GetIsDetached returns value of GetIsDetached and an boolean indicating if value is set.
func (c CommandDefinition) GetIsDetached() (bool, bool) {
	v, ok := getBool(c.IsDetached)
	if !ok {
		return getBool(c.IsDaemon)
	}

	return v, true
}

// GetRequireEnvVars returns value of RequireEnvVars and an boolean indicating if value is set.
func (c CommandDefinition) GetRequireEnvVars() (bool, bool) {
	return getBool(c.RequireEnvVars)
}

// GetIsTemplate returns value of IsTemplate and an boolean indicating if value is set.
func (c CommandDefinition) GetIsTemplate() (bool, bool) {
	return getBool(c.IsTemplate)
}

// GetIsInteractive returns value of IsInteractive and an boolean indicating if value is set.
func (c CommandDefinition) GetIsInteractive() (bool, bool) {
	return getBool(c.IsInteractive)
}

// GetAddGroups returns value of IsInteractive and an boolean indicating if value is set.
func (c CommandDefinition) GetAddGroups() (bool, bool) {
	return getBool(c.AddGroups)
}

// GetImpersonate returns value of Impersonate and an boolean indicating if value is set.
func (c CommandDefinition) GetImpersonate() (bool, bool) {
	return getBool(c.Impersonate)
}

// GetAutoMountWorkDir returns value of AutoMountWorkDir and an boolean indicating if value is set.
func (c CommandDefinition) GetAutoMountWorkDir() (bool, bool) {
	return getBool(c.AutoMountWorkDir)
}

// GetRemoveContainer returns value of RemoveContainer and an boolean indicating if value is set.
func (c CommandDefinition) GetRemoveContainer() (bool, bool) {
	return getBool(c.RemoveContainer)
}

// GetUniqueNames returns value of UniqueNames and an boolean indicating if value is set.
func (c CommandDefinition) GetUniqueNames() (bool, bool) {
	return getBool(c.UniqueNames)
}

// GetVolumes returns value of Volumes and an boolean indicating if value is set.
func (c CommandDefinition) GetVolumes() ([]string, bool) {
	return getStringSlice(c.Volumes)
}

// GetLinks returns value of Links and an boolean indicating if value is set.
func (c CommandDefinition) GetLinks() ([]string, bool) {
	return getStringSlice(c.Links)
}

// GetEnvVars returns value of EnvVars and an boolean indicating if value is set.
func (c CommandDefinition) GetEnvVars() ([]string, bool) {
	return getStringSlice(c.EnvVars)
}

// GetPorts returns value of Ports and an boolean indicating if value is set.
func (c CommandDefinition) GetPorts() ([]string, bool) {
	return getStringSlice(c.Ports)
}

// GetPortsFromParams returns value of Ports and an boolean indicating if value is set.
func (c CommandDefinition) GetPortsFromParams() ([]string, bool) {
	return getStringSlice(c.PortsFromParams)
}

// GetMergeTemplateArrays returns value of MergeTemplateArrays and an boolean indicating if value is set.
func (c CommandDefinition) GetMergeTemplateArrays() ([]string, bool) {
	return getStringSlice(c.MergeTemplateArrays)
}

// GetAdditionalArgs returns value of AdditionalArgs and an boolean indicating if value is set.
func (c CommandDefinition) GetAdditionalArgs() ([]string, bool) {
	return getStringSlice(c.AdditionalArgs)
}

// GetReplaceArgs returns value of ReplaceArgs and an boolean indicating if value is set.
func (c CommandDefinition) GetReplaceArgs() ([][]string, bool) {
	return get2DStringSlice(c.ReplaceArgs)
}

// HasTemplate indicates if the command definition has a template set.
func (c CommandDefinition) HasTemplate() bool { return c.Template != nil && *c.Template != "" }

// HasName indicates if the command definition has Name.
func (c CommandDefinition) HasName() bool { return c.Name != nil && *c.Name != "" }

// IsTemplateArrayMerged returns true if the given identifier is part of MergeTemplateArrays.
func (c CommandDefinition) IsTemplateArrayMerged(arrayKeyName string) bool {
	if identifiers, ok := c.GetMergeTemplateArrays(); ok {
		for _, identifier := range identifiers {
			if strings.EqualFold(identifier, arrayKeyName) {
				return true
			}
		}
	}

	return false
}

func getBool(b *bool) (bool, bool) {
	if b != nil {
		return *b, true
	}

	return false, false
}

func getString(s *string) (string, bool) {
	if s != nil {
		return *s, true
	}

	return "", false
}

func getStringSlice(s *[]string) ([]string, bool) {
	if s != nil {
		return *s, true
	}

	return []string{}, false
}

func get2DStringSlice(s *[][]string) ([][]string, bool) {
	if s != nil {
		return *s, true
	}

	return [][]string{}, false
}
