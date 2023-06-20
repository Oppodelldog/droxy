package arguments

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder"
)

func BuildTmpfsMounts(commandDef config.CommandDefinition, builder builder.Builder) error {
	if tmpfsMounts, ok := commandDef.GetTmpfsMounts(); ok {
		for _, tmpfsMount := range tmpfsMounts {
			resolvedTmpfsMount, err := resolveEnvVar(tmpfsMount)
			if err != nil {
				return err
			}

			builder.AddTmpfsMount(resolvedTmpfsMount)
		}
	}

	return nil
}
