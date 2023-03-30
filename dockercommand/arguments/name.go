package arguments

import (
	"crypto/rand"
	"fmt"
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder"
)

// NewNameArgumentBuilder has no implementation for windows, it is stubbed out.
func NewNameArgumentBuilder() ArgumentBuilderInterface {
	return &nameArgumentBuilder{
		nameRandomizerFunc: defaultNameRandomizerFunc,
	}
}

type nameArgumentBuilder struct {
	nameRandomizerFunc nameRandomizerFuncDef
}

type nameRandomizerFuncDef func(string) string

func (b *nameArgumentBuilder) BuildArgument(commandDef config.CommandDefinition, builder builder.Builder) error {
	if containerName, ok := commandDef.GetName(); ok {
		if uniqueContainerNames, ok := commandDef.GetUniqueNames(); ok && uniqueContainerNames {
			containerName = b.nameRandomizerFunc(containerName)
		}

		builder.SetContainerName(containerName)
	}

	return nil
}

func defaultNameRandomizerFunc(containerName string) string {
	randomValue := make([]byte, 4)
	_, _ = rand.Read(randomValue)

	return fmt.Sprintf("%s%v", containerName, randomValue)
}
