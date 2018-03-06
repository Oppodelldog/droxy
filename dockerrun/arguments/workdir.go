package arguments

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockerrun/builder"
)

// BuildWorkDir sets the working directory inside the container
func BuildWorkDir(commandDef *config.CommandDefinition, builder builder.Builder) error {
	if workDir, ok := commandDef.GetWorkDir(); ok {
		resolvedWorkDir,err := resolveEnvVar(workDir)
		if err != nil {
			return err
		}
		builder.SetWorkingDir(resolvedWorkDir)
	}
	return nil
}

