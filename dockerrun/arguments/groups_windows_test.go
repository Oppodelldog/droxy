package arguments

import (
	"testing"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockerrun/builder/mocks"
	"github.com/stretchr/testify/assert"
)

func TestAddGroups(t *testing.T) {
	assert.Nil(t, NewUserGroupsArgumentBuilder().BuildArgument(&config.CommandDefinition{}, &mocks.Builder{}))
}
