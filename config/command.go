package config

// CommandDefinition gives public access to the fields by accessor functions
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
	ReplaceArgs     *[][]string
	AdditionalArgs  *[]string
}

// GetIsTemplate returns value of IsTemplate and an boolean indicating if value is set.
func (c *CommandDefinition) GetIsTemplate() (bool, bool) {
	if c.IsTemplate != nil {
		return *c.IsTemplate, true
	}
	return false, false
}

// GetTemplate returns value of Template and an boolean indicating if value is set.
func (c *CommandDefinition) GetTemplate() (string, bool) {
	if c.Template != nil {
		return *c.Template, true
	}
	return "", false
}

// GetEntryPoint returns value of EntryPoint and an boolean indicating if value is set.
func (c *CommandDefinition) GetEntryPoint() (string, bool) {
	if c.EntryPoint != nil {
		return *c.EntryPoint, true
	}
	return "", false
}

// GetName returns value of Name and an boolean indicating if value is set.
func (c *CommandDefinition) GetName() (string, bool) {
	if c.Name != nil {
		return *c.Name, true
	}
	return "", false
}

// GetImage returns value of Image and an boolean indicating if value is set.
func (c *CommandDefinition) GetImage() (string, bool) {
	if c.Image != nil {
		return *c.Image, true
	}
	return "", false
}

// GetNetwork returns value of Network and an boolean indicating if value is set.
func (c *CommandDefinition) GetNetwork() (string, bool) {
	if c.Network != nil {
		return *c.Network, true
	}
	return "", false
}

// GetIsInteractive returns value of IsInteractive and an boolean indicating if value is set.
func (c *CommandDefinition) GetIsInteractive() (bool, bool) {
	if c.IsInteractive != nil {
		return *c.IsInteractive, true
	}
	return false, false
}

// GetAddGroups returns value of IsInteractive and an boolean indicating if value is set.
func (c *CommandDefinition) GetAddGroups() (bool, bool) {
	if c.AddGroups != nil {
		return *c.AddGroups, true
	}
	return false, false
}

// GetImpersonate returns value of Impersonate and an boolean indicating if value is set.
func (c *CommandDefinition) GetImpersonate() (bool, bool) {
	if c.Impersonate != nil {
		return *c.Impersonate, true
	}
	return false, false
}

// GetWorkDir returns value of Impersonate and an boolean indicating if value is set.
func (c *CommandDefinition) GetWorkDir() (string, bool) {
	if c.WorkDir != nil {
		return *c.WorkDir, true
	}
	return "", false
}

// GetRemoveContainer returns value of RemoveContainer and an boolean indicating if value is set.
func (c *CommandDefinition) GetRemoveContainer() (bool, bool) {
	if c.RemoveContainer != nil {
		return *c.RemoveContainer, true
	}
	return false, false
}

// GetVolumes returns value of Volumes and an boolean indicating if value is set.
func (c *CommandDefinition) GetVolumes() ([]string, bool) {
	if c.Volumes != nil {
		return *c.Volumes, true
	}
	return []string{}, false
}

// GetEnvVars returns value of EnvVars and an boolean indicating if value is set.
func (c *CommandDefinition) GetEnvVars() ([]string, bool) {
	if c.EnvVars != nil {
		return *c.EnvVars, true
	}
	return []string{}, false
}

// GetPorts returns value of Ports and an boolean indicating if value is set.
func (c *CommandDefinition) GetPorts() ([]string, bool) {
	if c.Ports != nil {
		return *c.Ports, true
	}
	return []string{}, false
}

// GetReplaceArgs returns value of ReplaceArgs and an boolean indicating if value is set.
func (c *CommandDefinition) GetReplaceArgs() ([][]string, bool) {
	if c.ReplaceArgs != nil {
		return *c.ReplaceArgs, true
	}
	return [][]string{}, false
}

// GetAdditionalArgs returns value of AdditionalArgs and an boolean indicating if value is set.
func (c *CommandDefinition) GetAdditionalArgs() ([]string, bool) {
	if c.AdditionalArgs != nil {
		return *c.AdditionalArgs, true
	}
	return []string{}, false
}

// HasTemplate indicates if the command definition has a template set
func (c *CommandDefinition) HasTemplate() bool { return c.Template != nil && *c.Template != "" }

// HasName indicates if the command definition has Name
func (c *CommandDefinition) HasName() bool { return c.Name != nil && *c.Name != "" }
