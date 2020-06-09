package arguments

import (
	"testing"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

type groupIDResolverStub struct {
	returnValue1 []string
	returnValue2 error
}

func (stub *groupIDResolverStub) getUserGroupIDs() ([]string, error) {
	return stub.returnValue1, stub.returnValue2
}

func TestNewUserGroupsArgumentBuilder(t *testing.T) {
	assert.IsType(t, new(userGroupsArgumentBuilder), NewUserGroupsArgumentBuilder())
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

	argumentBuilder := userGroupsArgumentBuilder{userGroupIdsResolver: &groupIDResolverStub{userGroupIDs, nil}}

	err := argumentBuilder.BuildArgument(commandDef, builder)
	if err != nil {
		t.Fatalf("Did not expect BuildArgument to return an error, but got: %v", err)
	}

	builder.AssertExpectations(t)
}

func TestAddGroups_AddGroupsTrue_IdResolverReturnsError_ExpectError(t *testing.T) {
	addGroupsFlag := true
	commandDef := &config.CommandDefinition{
		AddGroups: &addGroupsFlag,
	}
	builder := &mocks.Builder{}

	idResolverError := errors.New("Id resolver error")
	argumentBuilder := userGroupsArgumentBuilder{userGroupIdsResolver: &groupIDResolverStub{nil, idResolverError}}

	err := argumentBuilder.BuildArgument(commandDef, builder)

	assert.Equal(t, idResolverError, err)
}

func TestAddGroups_AddGroupsFalse_ExpectBuilderNotCalled(t *testing.T) {
	commandDef := &config.CommandDefinition{}
	builder := &mocks.Builder{}

	argumentBuilder := userGroupsArgumentBuilder{userGroupIdsResolver: &groupIDResolverStub{}}

	err := argumentBuilder.BuildArgument(commandDef, builder)
	if err != nil {
		t.Fatalf("Did not expect BuildArgument to return an error, but got: %v", err)
	}

	assert.Empty(t, builder.Calls)
}

func TestCurrentUserGroupIDsResolver_getUserGroupIDs(t *testing.T) {
	resolver := currentUserGroupIDsResolver{}
	userGroupIds, err := resolver.getUserGroupIDs()

	assert.NoError(t, err)
	assert.NotEmpty(t, userGroupIds)
}
