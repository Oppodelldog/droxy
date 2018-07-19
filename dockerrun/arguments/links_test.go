package arguments

import (
	"os"
	"testing"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockerrun/builder/mocks"
	"github.com/stretchr/testify/assert"
)

func TestBuildLinks_LinksAreSet(t *testing.T) {
	Links := []string{"/home/samwise/:/app/walktovulcano"}

	commandDef := &config.CommandDefinition{
		Links: &Links,
	}

	builder := &mocks.Builder{}
	builder.On("AddLinkMapping", Links[0]).Return(builder)

	BuildLinks(commandDef, builder)

	builder.AssertExpectations(t)
}

func TestBuildLinks_LinksAreNotSet(t *testing.T) {
	commandDef := &config.CommandDefinition{
		Links: nil,
	}

	builder := &mocks.Builder{}

	BuildLinks(commandDef, builder)

	assert.Empty(t, builder.Calls)
}

func TestBuildLinks_LinksEnvVarsAreResolves(t *testing.T) {
	os.Setenv("WHERE_THE_HECK_AM_I", "NO_CLUE")
	Links := []string{"${WHERE_THE_HECK_AM_I}"}

	commandDef := &config.CommandDefinition{
		Links: &Links,
	}

	builder := &mocks.Builder{}
	builder.On("AddLinkMapping", "NO_CLUE").Return(builder)

	BuildLinks(commandDef, builder)

	builder.AssertExpectations(t)
}
