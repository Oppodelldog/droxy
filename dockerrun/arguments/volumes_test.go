package arguments

import (
	"os"
	"testing"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockerrun/builder/mocks"
	"github.com/stretchr/testify/assert"
)

func TestBuildVolumes_VolumesAreSet(t *testing.T) {
	volumes := []string{"/home/samwise/:/app/walktovulcano"}

	commandDef := &config.CommandDefinition{
		Volumes: &volumes,
	}

	builder := &mocks.Builder{}
	builder.On("AddVolumeMapping", volumes[0]).Return(builder)

	BuildVolumes(commandDef, builder)

	builder.AssertExpectations(t)
}

func TestBuildVolumes_VolumesAreNotSet(t *testing.T) {
	commandDef := &config.CommandDefinition{
		Volumes: nil,
	}

	builder := &mocks.Builder{}

	BuildVolumes(commandDef, builder)

	assert.Empty(t, builder.Calls)
}

func TestBuildVolumes_VolumesEnvVarsAreResolves(t *testing.T) {
	os.Setenv("WHERE_THE_HECK_AM_I", "NO_CLUE")
	volumes := []string{"${WHERE_THE_HECK_AM_I}"}

	commandDef := &config.CommandDefinition{
		Volumes: &volumes,
	}

	builder := &mocks.Builder{}
	builder.On("AddVolumeMapping", "NO_CLUE").Return(builder)

	BuildVolumes(commandDef, builder)

	builder.AssertExpectations(t)
}
