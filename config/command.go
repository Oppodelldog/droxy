package config

import "strings"

// CommandDefinition gives public access to the fields by accessor functions.
type CommandDefinition struct {
	// OS is set to load this configuration only when executed from the given operating system.
	// accepted values are: "windows", "linux" , "darwin"
	OS *string
	// RequireEnvVars requires all EnvVars to be resolvable.
	// If a variable is not set, the command fails.
	RequireEnvVars *bool
	// IsTemplate prevents an executable file to be generated for this command entry.
	IsTemplate *bool
	// Template defines the template commands name that settings will be used as a base for this command.
	Template *string
	// EntryPoint represents the docker flag --entrypoint
	EntryPoint *string
	// Command represents the docker run [COMMAND] part.
	Command *string
	// Name represents the docker run flag --name
	Name *string
	// UniqueNames will make Name unique by appending a unique value to the Name.
	UniqueNames *bool
	// Image represents the docker Image to be used.
	Image *string
	// Network represents the docker run flag --network.
	Network *string
	// EnvFile represents the docker run flag --env-file
	EnvFile *string
	// EnvFile represents the docker run flag --ip
	IP *string
	// IsInteractive sets -i and if IsInteractive is set true, attaches (-a) streams STDIN STDOUT STDERR
	IsInteractive *bool
	// IsDetached represents the docker run flag -d
	IsDetached *bool
	// IsDaemon represents the docker run flag -d
	IsDaemon *bool // deprecated
	// Volumes holds a list of volume mappings which correspond to the docker run flag -v
	Volumes *[]string
	// Links holds a list of links which correspond to the docker run flag --link
	Links *[]string
	// EnvVars holds a list of environment variable mappings which correspond to the docker run flag -e
	EnvVars *[]string
	// EnvVarOverwrites holds a list of environment variable mappings which be taken in favor of env variable values
	// when applying EnvVars to the docker run flag -e
	EnvVarOverwrites *[]string
	// Ports holds a list of port mappings which correspond to the docker run flag -p
	Ports *[]string
	// PortsFromParams holds a list of regular expressions that are used to read a port from a command argument.
	// The parsed port will be applied like follows: -p port:port.
	PortsFromParams *[]string
	// MergeTemplateArrays holds a list of array names that should be applied from the template and merged with
	// The current command. By default a array definition in a command would overwrite a template array.
	// Possible to merge are the following arrays: Volumes, Links, EnvVars, Ports, PortsFromParams, AdditionalArgs.
	MergeTemplateArrays *[]string
	// AddGroups, if set to true, will automatically resolve and add the current user groups to the container.
	// The corresponding docker run flag is --group-add.
	AddGroups *bool
	// Impersonate sets the current user,group for the container.  docker run flag -u.
	Impersonate *bool
	// WorkDir sets the workdir inside the container. docker run flag -w.
	WorkDir *string
	// AutoMountWorkDir will add a volume mapping for the current WorkDir like this: -v WorkDir:WorkDir
	AutoMountWorkDir *bool
	// RemoveContainer removes the container after execution. docker flag --rm.
	RemoveContainer *bool
	// ReplaceArgs allows to manipulate the command arguments.
	// Each entry must be an array of two entries. The first entry defines which argument shall be replaced, the second
	// one defines the replacement.
	// Example: [["arg2", "arg99"]] would replace "arg2" with "arg99".
	ReplaceArgs *[][]string
	// AdditionalArgs contains a list of strings that are prepended the original command arguments.
	AdditionalArgs *[]string
}

// GetOS returns value of OS and an boolean indicating if value is set.
func (c CommandDefinition) GetOS() (string, bool) {
	return getString(c.OS)
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

// GetEnvVarOverwrites returns value of EnvVarOverwrites and an boolean indicating if value is set.
func (c CommandDefinition) GetEnvVarOverwrites() ([]string, bool) {
	return getStringSlice(c.EnvVarOverwrites)
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
