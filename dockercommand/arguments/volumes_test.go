package arguments

import (
	"os"
	"testing"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder/mocks"
	"github.com/stretchr/testify/assert"
)

func TestBuildVolumes_VolumesAreSet(t *testing.T) {
	volumes := []string{"/home/samwise/:/app/walktovulcano"}

	commandDef := &config.CommandDefinition{
		Volumes: &volumes,
	}

	builder := &mocks.Builder{}
	builder.On("AddVolumeMapping", volumes[0]).Return(builder)

	err := BuildVolumes(commandDef, builder)
	if err != nil {
		t.Fatalf("Did not expect BuildVolumes to return an error, but got: %v", err)
	}

	builder.AssertExpectations(t)
}

func TestBuildVolumes_VolumesAreNotSet(t *testing.T) {
	commandDef := &config.CommandDefinition{
		Volumes: nil,
	}

	builder := &mocks.Builder{}

	err := BuildVolumes(commandDef, builder)
	if err != nil {
		t.Fatalf("Did not expect BuildVolumes to return an error, but got: %v", err)
	}

	assert.Empty(t, builder.Calls)
}

func TestBuildVolumes_VolumesEnvVarsAreResolves(t *testing.T) {
	err := os.Setenv("WHERE_THE_HECK_AM_I", "NO_CLUE")
	if err != nil {
		t.Fatalf("Did not expect os.Setenv to return an error, but got: %v", err)
	}

	volumes := []string{"${WHERE_THE_HECK_AM_I}"}

	commandDef := &config.CommandDefinition{
		Volumes: &volumes,
	}

	builder := &mocks.Builder{}
	builder.On("AddVolumeMapping", "NO_CLUE").Return(builder)

	err = BuildVolumes(commandDef, builder)
	if err != nil {
		t.Fatalf("Did not expect BuildVolumes to return an error, but got: %v", err)
	}

	builder.AssertExpectations(t)
}
