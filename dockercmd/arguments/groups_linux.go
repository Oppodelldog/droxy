package arguments

import (
	"os/user"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercmd/builder"
)

func NewUserGroupsArgumentBuilder() *UserGroupsArgumentBuilder {
	return &UserGroupsArgumentBuilder{
		userGroupIdsResolver: &currentUserGroupIDsResolver{},
	}
}

type UserGroupsArgumentBuilder struct {
	userGroupIdsResolver userGroupIdsResolverInterface
}

func (b *UserGroupsArgumentBuilder) BuildArgument(commandDef *config.CommandDefinition, builder builder.Builder) error {

	if addGroups, ok := commandDef.GetAddGroups(); ok && addGroups {

		groupIDs, err := b.userGroupIdsResolver.getUserGroupIDs()
		if err != nil {
			return err
		}

		for _, groupID := range groupIDs {
			builder.AddGroup(groupID)
		}
	}
	return nil
}

type userGroupIdsResolverInterface interface {
	getUserGroupIDs() ([]string, error)
}

type currentUserGroupIDsResolver struct{}

func (r *currentUserGroupIDsResolver) getUserGroupIDs() ([]string, error) {

	currentUser, err := user.Current()
	if err != nil {
		return nil, err
	}

	return currentUser.GroupIds()

}
