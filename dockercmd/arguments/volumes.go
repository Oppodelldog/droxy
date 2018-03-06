package arguments

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercmd/builder"
)

func BuildVolumes(commandDef *config.CommandDefinition, builder builder.Builder) error {
	if volumes, ok := commandDef.GetVolumes(); ok {
		for _, volume := range volumes {
			resolvedVolume, err := resolveEnvVar(volume)
			if err != nil {
				return err
			}
			builder.AddVolumeMapping(resolvedVolume)
		}
	}
	return nil
}
