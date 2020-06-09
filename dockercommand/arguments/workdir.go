package arguments

import (
	"fmt"
	"os"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder"
)

// BuildWorkDir sets the working directory inside the container
// if the directory exists on the host, it is automatically mounted when the appropriate option is set.
func BuildWorkDir(commandDef *config.CommandDefinition, builder builder.Builder) error {
	if workDir, ok := commandDef.GetWorkDir(); ok {
		resolvedWorkDir, err := resolveEnvVar(workDir)
		if err != nil {
			return err
		}

		builder.SetWorkingDir(resolvedWorkDir)

		if isAutoMount, ok := commandDef.GetAutoMountWorkDir(); ok && isAutoMount {
			if _, err := os.Stat(resolvedWorkDir); !os.IsNotExist(err) {
				builder.AddVolumeMapping(fmt.Sprintf("%s:%s", resolvedWorkDir, resolvedWorkDir))
			}
		}
	}

	return nil
}
