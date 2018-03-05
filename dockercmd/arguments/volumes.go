package arguments

import (
	"fmt"
	"strings"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercmd/builder"
)

func BuildVolumes(commandDef *config.CommandDefinition, builder builder.Builder) error {
	if volumes, ok := commandDef.GetVolumes(); ok {
		return buildVolumes(volumes, builder)
	}
	return nil
}

func buildVolumes(volumes []string, builder builder.Builder) error {
	for _, volume := range volumes {
		volumeParts := strings.Split(volume, ":")
		if len(volumeParts) < 2 || len(volumeParts) > 3 {
			return fmt.Errorf("invalid number of volume parts in '%s'", volume)
		}

		var hostPart, containerPart, options string
		var resolveErr error

		if len(volumeParts) > 0 {
			hostPart, resolveErr = resolveEnvVar(volumeParts[0])
			if resolveErr != nil {
				return resolveErr
			}
		}
		if len(volumeParts) > 1 {
			containerPart, resolveErr = resolveEnvVar(volumeParts[1])
			if resolveErr != nil {
				return resolveErr
			}
		}
		if len(volumeParts) > 2 {
			options, resolveErr = resolveEnvVar(volumeParts[2])
			if resolveErr != nil {
				return resolveErr
			}
		}

		builder.AddVolumeMapping(hostPart, containerPart, options)
	}

	return nil
}
