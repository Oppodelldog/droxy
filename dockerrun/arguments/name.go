package arguments

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockerrun/builder"
)

//NewUserGroupsArgumentBuilder has no implementation for windows, it is stubbed out
func NewNameArgumentBuilder() ArgumentBuilderInterface {
	return &nameArgumentBuilder{
		nameRandomizerFunc: defaultNameRandomizerFunc,
	}
}

type nameArgumentBuilder struct {
	nameRandomizerFunc nameRandomizerFuncDef
}

type nameRandomizerFuncDef func(string) string

func (b *nameArgumentBuilder) BuildArgument(commandDef *config.CommandDefinition, builder builder.Builder) error {

	if containerName, ok := commandDef.GetName(); ok {
		if uniqueContainerNames, ok := commandDef.GetUniqueNames(); ok && uniqueContainerNames {
			containerName = b.nameRandomizerFunc(containerName)
		}

		builder.SetContainerName(containerName)
	}

	return nil
}

func defaultNameRandomizerFunc(containerName string) string {

	rand.Seed(time.Now().UnixNano())
	randomValue := rand.Int31()
	return fmt.Sprintf("%s%v", containerName, randomValue)
}
