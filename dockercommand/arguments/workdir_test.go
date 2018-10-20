package arguments

import (
	"os"
	"testing"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder/mocks"
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

func TestBuildWorkDir_ResolvesEnvVars(t *testing.T) {
	expectedWorkingDir := "/home/somewhere"
	os.Setenv("CURRENT_WORKING_DIR", expectedWorkingDir)
	defer os.Unsetenv("CURRENT_WORKING_DIR")

	workDir := "${CURRENT_WORKING_DIR}"
	commandDef := &config.CommandDefinition{
		WorkDir: &workDir,
	}

	builder := &mocks.Builder{}

	builder.On("SetWorkingDir", expectedWorkingDir).Return(builder)

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
