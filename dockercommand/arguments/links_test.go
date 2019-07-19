package arguments

import (
	"os"
	"testing"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder/mocks"
	"github.com/stretchr/testify/assert"
)

func TestBuildLinks_LinksAreSet(t *testing.T) {
	Links := []string{"/home/samwise/:/app/walktovulcano"}

	commandDef := &config.CommandDefinition{
		Links: &Links,
	}

	builder := &mocks.Builder{}
	builder.On("AddLinkMapping", Links[0]).Return(builder)

	err := BuildLinks(commandDef, builder)
	if err != nil {
		t.Fatalf("Did not expect BuildLinks to return an error, but got: %v", err)
	}

	builder.AssertExpectations(t)
}

func TestBuildLinks_LinksAreNotSet(t *testing.T) {
	commandDef := &config.CommandDefinition{
		Links: nil,
	}

	builder := &mocks.Builder{}

	err := BuildLinks(commandDef, builder)
	if err != nil {
		t.Fatalf("Did not expect BuildLinks to return an error, but got: %v", err)
	}

	assert.Empty(t, builder.Calls)
}

func TestBuildLinks_LinksEnvVarsAreResolves(t *testing.T) {
	err := os.Setenv("WHERE_THE_HECK_AM_I", "NO_CLUE")
	if err != nil {
		t.Fatalf("Did not expect os.Setenv to return an error, but got: %v", err)
	}
	Links := []string{"${WHERE_THE_HECK_AM_I}"}

	commandDef := &config.CommandDefinition{
		Links: &Links,
	}

	builder := &mocks.Builder{}
	builder.On("AddLinkMapping", "NO_CLUE").Return(builder)

	err = BuildLinks(commandDef, builder)
	if err != nil {
		t.Fatalf("Did not expect BuildLinks to return an error, but got: %v", err)
	}

	builder.AssertExpectations(t)
}
