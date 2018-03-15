package arguments

import (
	"testing"

	"fmt"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockerrun/builder/mocks"
	"github.com/stretchr/testify/assert"
)

func TestBuildName_NameIsSet_AndNotUnique_ExpectAppropriateBuilderCall(t *testing.T) {
	randomNamePartStub := "123"

	testCases := []struct {
		testNo                int
		containerName         string
		uniqueNames           bool
		expectedContainerName string
	}{
		{testNo: 1, containerName: "my-container", uniqueNames: false, expectedContainerName: "my-container"},
		{testNo: 2, containerName: "my-container", uniqueNames: true, expectedContainerName: "my-container123"},
	}

	for _, testCase := range testCases {

		t.Logf("testCase no: %v", testCase.testNo)

		containerName := testCase.containerName
		uniqueNames := testCase.uniqueNames
		commandDef := &config.CommandDefinition{
			Name:        &containerName,
			UniqueNames: &uniqueNames,
		}
		builder := &mocks.Builder{}
		builder.On("SetContainerName", testCase.expectedContainerName).Return(builder)

		nameBuilder := nameArgumentBuilder{nameRandomizerFunc: func(containerName string) string { return fmt.Sprintf("%s%v", containerName, randomNamePartStub) }}
		nameBuilder.BuildArgument(commandDef, builder)

		builder.AssertExpectations(t)
	}
}

func TestBuildName_NameIsNotSet_AndNotUniqueNames(t *testing.T) {
	commandDef := &config.CommandDefinition{}
	builder := &mocks.Builder{}

	nameBuilder := NewNameArgumentBuilder()
	nameBuilder.BuildArgument(commandDef, builder)

	assert.Empty(t, builder.Calls)
}
