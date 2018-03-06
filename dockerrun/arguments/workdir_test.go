package arguments

import (
	"testing"
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockerrun/builder/mocks"
	"github.com/stretchr/testify/assert"
)

func TestBuildWorkDir_WorkDirIsSet(t *testing.T) {
	workDir := "/someDir"
	commandDef := &config.CommandDefinition{
		WorkDir: &workDir,
	}

	builder := &mocks.Builder{}

	builder.On("SetWorkingDir", workDir).Return(builder)

	BuildWorkDir(commandDef, builder)

	builder.AssertExpectations(t)
}

func TestBuildWorkDir_WorkDirIsNotSet(t *testing.T) {
	commandDef := &config.CommandDefinition{
		WorkDir: nil,
	}

	builder := &mocks.Builder{}

	BuildWorkDir(commandDef, builder)

	assert.Empty(t, builder.Calls)
}
