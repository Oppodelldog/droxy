package config

import "reflect"

func mergeCommand(baseCommand CommandDefinition, overlayCommand CommandDefinition) CommandDefinition {
	mergedCommand := CommandDefinition{}

	mergedCommand.Name = resolvePropertyString(baseCommand.Name, overlayCommand.Name)
	mergedCommand.EntryPoint = resolvePropertyString(baseCommand.EntryPoint, overlayCommand.EntryPoint)
	mergedCommand.Command = resolvePropertyString(baseCommand.Command, overlayCommand.Command)
	mergedCommand.Image = resolvePropertyString(baseCommand.Image, overlayCommand.Image)
	mergedCommand.WorkDir = resolvePropertyString(baseCommand.WorkDir, overlayCommand.WorkDir)
	mergedCommand.AutoMountWorkDir = resolvePropertyBool(baseCommand.AutoMountWorkDir, overlayCommand.AutoMountWorkDir)
	mergedCommand.Network = resolvePropertyString(baseCommand.Network, overlayCommand.Network)
	mergedCommand.EnvFile = resolvePropertyString(baseCommand.EnvFile, overlayCommand.EnvFile)
	mergedCommand.IP = resolvePropertyString(baseCommand.IP, overlayCommand.IP)

	mergedCommand.AddGroups = resolvePropertyBool(baseCommand.AddGroups, overlayCommand.AddGroups)
	mergedCommand.RemoveContainer = resolvePropertyBool(baseCommand.RemoveContainer, overlayCommand.RemoveContainer)
	mergedCommand.Impersonate = resolvePropertyBool(baseCommand.Impersonate, overlayCommand.Impersonate)
	mergedCommand.IsInteractive = resolvePropertyBool(baseCommand.IsInteractive, overlayCommand.IsInteractive)
	mergedCommand.IsDetached = resolvePropertyBool(baseCommand.IsDetached, overlayCommand.IsDetached)
	mergedCommand.IsDaemon = resolvePropertyBool(baseCommand.IsDaemon, overlayCommand.IsDaemon)
	mergedCommand.UniqueNames = resolvePropertyBool(baseCommand.UniqueNames, overlayCommand.UniqueNames)

	mergedCommand.MergeTemplateArrays = resolvePropertyStringArray(
		false,
		baseCommand.MergeTemplateArrays,
		overlayCommand.MergeTemplateArrays,
	)
	mergedCommand.Volumes = resolvePropertyStringArray(
		mergedCommand.IsTemplateArrayMerged("Volumes"),
		baseCommand.Volumes,
		overlayCommand.Volumes,
	)
	mergedCommand.Tmpfs = resolvePropertyStringArray(
		mergedCommand.IsTemplateArrayMerged("Tmpfs"),
		baseCommand.Tmpfs,
		overlayCommand.Tmpfs,
	)
	mergedCommand.Links = resolvePropertyStringArray(
		mergedCommand.IsTemplateArrayMerged("Links"),
		baseCommand.Links,
		overlayCommand.Links,
	)
	mergedCommand.EnvVars = resolvePropertyStringArray(
		mergedCommand.IsTemplateArrayMerged("EnvVars"),
		baseCommand.EnvVars,
		overlayCommand.EnvVars,
	)
	mergedCommand.Ports = resolvePropertyStringArray(
		mergedCommand.IsTemplateArrayMerged("Ports"),
		baseCommand.Ports,
		overlayCommand.Ports,
	)
	mergedCommand.PortsFromParams = resolvePropertyStringArray(
		mergedCommand.IsTemplateArrayMerged("PortsFromParams"),
		baseCommand.PortsFromParams,
		overlayCommand.PortsFromParams,
	)
	mergedCommand.AdditionalArgs = resolvePropertyStringArray(
		mergedCommand.IsTemplateArrayMerged("AdditionalArgs"),
		baseCommand.AdditionalArgs,
		overlayCommand.AdditionalArgs,
	)

	mergedCommand.ReplaceArgs = resolvePropertyStringArray2D(baseCommand.ReplaceArgs, overlayCommand.ReplaceArgs)

	return mergedCommand
}

func resolveProperty(base interface{}, overlay interface{}) interface{} {
	if base == nil && overlay == nil {
		return nil
	}

	if !reflect.ValueOf(overlay).IsNil() {
		return overlay
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

func resolvePropertyStringArray(isMerged bool, sBase *[]string, sOverlay *[]string) *[]string {
	if isMerged {
		if sBase != nil && sOverlay != nil {
			mergedArray := append(*sBase, *sOverlay...)

			return &mergedArray
		}
	}

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
