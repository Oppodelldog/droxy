package arguments

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockerrun/builder"
)

// NewUserGroupsArgumentBuilder has no implementation for windows, it is stubbed out.
func NewUserGroupsArgumentBuilder() ArgumentBuilderInterface {
	return &userGroupsArgumentBuilder{}
}

type userGroupsArgumentBuilder struct {
}

func (b *userGroupsArgumentBuilder) BuildArgument(commandDef *config.CommandDefinition, builder builder.Builder) error {
	_ = commandDef
	_ = builder

	return nil
}
