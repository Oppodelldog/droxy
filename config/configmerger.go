package config

func mergeCommand(baseCommand *CommandDefinition, overlayCommand *CommandDefinition) *CommandDefinition {
	mergedCommand := new(CommandDefinition)

	mergedCommand.Name = resolvePropertyString(baseCommand.Name, overlayCommand.Name)
	mergedCommand.EntryPoint = resolvePropertyString(baseCommand.EntryPoint, overlayCommand.EntryPoint)
	mergedCommand.Command = resolvePropertyString(baseCommand.Command, overlayCommand.Command)
	mergedCommand.Image = resolvePropertyString(baseCommand.Image, overlayCommand.Image)
	mergedCommand.WorkDir = resolvePropertyString(baseCommand.WorkDir, overlayCommand.WorkDir)
	mergedCommand.Network = resolvePropertyString(baseCommand.Network, overlayCommand.Network)
	mergedCommand.EnvFile = resolvePropertyString(baseCommand.EnvFile, overlayCommand.EnvFile)
	mergedCommand.Ip = resolvePropertyString(baseCommand.Ip, overlayCommand.Ip)

	mergedCommand.AddGroups = resolvePropertyBool(baseCommand.AddGroups, overlayCommand.AddGroups)
	mergedCommand.RemoveContainer = resolvePropertyBool(baseCommand.RemoveContainer, overlayCommand.RemoveContainer)
	mergedCommand.Impersonate = resolvePropertyBool(baseCommand.Impersonate, overlayCommand.Impersonate)
	mergedCommand.IsInteractive = resolvePropertyBool(baseCommand.IsInteractive, overlayCommand.IsInteractive)
	mergedCommand.IsDaemon = resolvePropertyBool(baseCommand.IsDaemon, overlayCommand.IsDaemon)
	mergedCommand.UniqueNames = resolvePropertyBool(baseCommand.UniqueNames, overlayCommand.UniqueNames)

	mergedCommand.Volumes = resolvePropertyStringArray(baseCommand.Volumes, overlayCommand.Volumes)
	mergedCommand.Links = resolvePropertyStringArray(baseCommand.Links, overlayCommand.Links)
	mergedCommand.EnvVars = resolvePropertyStringArray(baseCommand.EnvVars, overlayCommand.EnvVars)
	mergedCommand.Ports = resolvePropertyStringArray(baseCommand.Ports, overlayCommand.Ports)
	mergedCommand.AdditionalArgs = resolvePropertyStringArray(baseCommand.AdditionalArgs, overlayCommand.AdditionalArgs)

	mergedCommand.ReplaceArgs = resolvePropertyStringArray2D(baseCommand.ReplaceArgs, overlayCommand.ReplaceArgs)

	return mergedCommand
}

func resolveProperty(base interface{}, overlay interface{}) interface{} {

	if base == nil && overlay == nil {
		return nil
	}

	if overlay != nil {
		return overlay
	}
	if base == nil {
		return nil
	}

	return base
}

func resolvePropertyBool(bBase *bool, bOverlay *bool) *bool {
	res := resolveProperty(bBase, bOverlay)
	if v, ok := res.(*bool); ok && v != nil {
		c := *v
		return &c
	}

	return nil
}

func resolvePropertyString(sBase *string, sOverlay *string) *string {
	res := resolveProperty(sBase, sOverlay)
	if v, ok := res.(*string); ok && v != nil {
		c := *v
		return &c
	}

	return nil
}

func resolvePropertyStringArray(sBase *[]string, sOverlay *[]string) *[]string {
	res := resolveProperty(sBase, sOverlay)
	if v, ok := res.(*[]string); ok && v != nil {
		c := *v
		return &c
	}

	return nil
}

func resolvePropertyStringArray2D(sBase *[][]string, sOverlay *[][]string) *[][]string {
	res := resolveProperty(sBase, sOverlay)
	if v, ok := res.(*[][]string); ok && v != nil {
		c := *v
		return &c
	}

	return nil
}
