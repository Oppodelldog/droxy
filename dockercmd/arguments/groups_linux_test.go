package arguments

import (
	"testing"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercmd/builder/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

type groupIdResolverStub struct {
	returnValue1 []string
	returnValue2 error
}

func (stub *groupIdResolverStub) getUserGroupIDs() ([]string, error) {
	return stub.returnValue1, stub.returnValue2
}

func TestNewUserGroupsArgumentBuilder(t *testing.T) {
	assert.IsType(t, new(UserGroupsArgumentBuilder), NewUserGroupsArgumentBuilder())
}

func TestAddGroups_AddGroupsTrue_ExpectAllUsersGroupsAddedToBuilder(t *testing.T) {

	addGroupsFlag := true
	commandDef := &config.CommandDefinition{
		AddGroups: &addGroupsFlag,
	}

	userGroupIDs := []string{"1", "99", "1002"}

	builder := &mocks.Builder{}
	builder.
		On("AddGroup", "1").Return(builder).
		On("AddGroup", "99").Return(builder).
		On("AddGroup", "1002").Return(builder)

	argumentBuilder := UserGroupsArgumentBuilder{userGroupIdsResolver: &groupIdResolverStub{userGroupIDs, nil}}

	argumentBuilder.BuildArgument(commandDef, builder)

	builder.AssertExpectations(t)
}

func TestAddGroups_AddGroupsTrue_IdResolverReturnsError_ExpectError(t *testing.T) {

	addGroupsFlag := true
	commandDef := &config.CommandDefinition{
		AddGroups: &addGroupsFlag,
	}
	builder := &mocks.Builder{}

	idResolverError := errors.New("Id resolver error")
	argumentBuilder := UserGroupsArgumentBuilder{userGroupIdsResolver: &groupIdResolverStub{nil, idResolverError}}

	err := argumentBuilder.BuildArgument(commandDef, builder)

	assert.Equal(t, idResolverError, err)
}

func TestAddGroups_AddGroupsFalse_ExpectBuilderNotCalled(t *testing.T) {
	commandDef := &config.CommandDefinition{}
	builder := &mocks.Builder{}

	argumentBuilder := UserGroupsArgumentBuilder{userGroupIdsResolver: &groupIdResolverStub{}}
	argumentBuilder.BuildArgument(commandDef, builder)

	assert.Empty(t, builder.Calls)
}

func TestCurrentUserGroupIDsResolver_getUserGroupIDs(t *testing.T) {
	resolver := currentUserGroupIDsResolver{}
	userGroupIds, err := resolver.getUserGroupIDs()

	assert.NoError(t, err)
	assert.NotEmpty(t, userGroupIds)
}
