package arguments

import (
	"reflect"
	"testing"

	"fmt"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder/mocks"
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

		nameBuilder := nameArgumentBuilder{
			nameRandomizerFunc: func(containerName string) string {
				return fmt.Sprintf("%s%v", containerName, randomNamePartStub)
			},
		}

		err := nameBuilder.BuildArgument(commandDef, builder)
		if err != nil {
			t.Fatalf("Did not expect BuildArgument to return an error, but got: %v", err)
		}

		builder.AssertExpectations(t)
	}
}

func TestBuildName_NameIsNotSet_AndNotUniqueNames(t *testing.T) {
	commandDef := &config.CommandDefinition{}
	builder := &mocks.Builder{}

	nameBuilder := NewNameArgumentBuilder()
	err := nameBuilder.BuildArgument(commandDef, builder)
	if err != nil {
		t.Fatalf(
			"Did not expect BuildArgument to return an error, but got: %v",
			err,
		)
	}

	assert.Empty(t, builder.Calls)
}

func TestBuildName_ArgumentsBuilderNameRandomizerFunc(t *testing.T) {
	got := reflect.ValueOf(NewNameArgumentBuilder().(*nameArgumentBuilder).nameRandomizerFunc).Pointer()
	want := reflect.ValueOf(defaultNameRandomizerFunc).Pointer()

	if got != want {
		t.Fatalf("nameArgumentBuilder.nameRandomizerFunc is not set to defaultNameRandomizerFunc")
	}
}

func Test_defaultNameRandomizerFunc(t *testing.T) {
	const name = "some-name"
	randomizedName := defaultNameRandomizerFunc(name)

	assert.Contains(t, randomizedName, name)
	assert.NotEqual(t, randomizedName, name)
}
