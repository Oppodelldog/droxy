package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_resolvePropertyStringArray2D_BaseSet_OverlayNotSet(t *testing.T) {
	base := &[][]string{
		{"a", "b"},
	}
	var overlay *[][]string

	result := resolvePropertyStringArray2D(base, overlay)

	expectedResult := &[][]string{
		{"a", "b"},
	}

	assert.Equal(t, expectedResult, result)
}

func Test_resolvePropertyStringArray2D_BaseNotSet_OverlaySet(t *testing.T) {
	var base *[][]string
	overlay := &[][]string{
		{"a", "b"},
	}

	result := resolvePropertyStringArray2D(base, overlay)

	expectedResult := &[][]string{
		{"a", "b"},
	}

	assert.Equal(t, expectedResult, result)
}

func Test_resolvePropertyStringArray2D_BaseAndOverlaySet(t *testing.T) {
	base := &[][]string{
		{"a", "b"},
	}
	overlay := &[][]string{
		{"c", "d"},
	}

	result := resolvePropertyStringArray2D(base, overlay)

	expectedResult := &[][]string{
		{"c", "d"},
	}

	assert.Equal(t, expectedResult, result)
}

func TestConfiguration_FindCommandByName(t *testing.T) {

	nameA := "COMMAND-A"
	commandA := CommandDefinition{Name: &nameA}
	nameX := "COMMAND-X"
	commandX := CommandDefinition{Name: &nameX}
	nameZ := "COMMAND-Z"
	commandZ := CommandDefinition{Name: &nameZ}

	cfg := Configuration{
		Command: []CommandDefinition{
			commandA,
			commandX,
			commandZ,
		},
	}

	commandDef, err := cfg.FindCommandByName("COMMAND-X")
	if err != nil {
		t.Fatalf("Did not expect cgf.FindCommandByName to return an error, but got: %v", err)
	}

	assert.Equal(t, &commandX, commandDef)
}

func TestConfiguration_FindCommandByName_NotFoundError(t *testing.T) {

	nameA := "COMMAND-A"
	commandA := CommandDefinition{Name: &nameA}

	cfg := Configuration{
		Command: []CommandDefinition{
			commandA,
		},
	}

	_, err := cfg.FindCommandByName("COMMAND-X")

	assert.Error(t, err)
}

func TestConfiguration_FindCommandByName_ResolvesTemplate(t *testing.T) {

	templateName := "TEMPLATE-A"
	isTemplate := true
	templateNetwork := "templateNetwork-by-templte"
	templateCommand := CommandDefinition{Name: &templateName, IsTemplate: &isTemplate, Network: &templateNetwork}

	nameA := "COMMAND-A"
	commandA := CommandDefinition{Name: &nameA, Template: &templateName}

	cfg := Configuration{
		Command: []CommandDefinition{
			commandA,
			templateCommand,
		},
	}

	commandDef, err := cfg.FindCommandByName("COMMAND-A")
	if err != nil {
		t.Fatalf("Did not expect cgf.FindCommandByName to return an error, but got: %v", err)
	}

	assert.Equal(t, &templateNetwork, commandDef.Network)
}

func TestConfiguration_FindCommandByName_TemplateNotFoundError(t *testing.T) {

	templateName := "yes-template-does-not-exist"
	nameA := "COMMAND-A"
	commandA := CommandDefinition{Name: &nameA, Template: &templateName}

	cfg := Configuration{
		Command: []CommandDefinition{
			commandA,
		},
	}

	_, err := cfg.FindCommandByName("COMMAND-A")

	assert.Error(t, err)
}

func TestConfiguration_FindCommandByName_TemplateHasTemplate(t *testing.T) {

	template1Name := "template1"
	template1EntryPoint := "template1EntryPoint"
	template2Name := "template2"
	template1 := CommandDefinition{Name: &template1Name, EntryPoint: &template1EntryPoint}
	template2 := CommandDefinition{Name: &template2Name, Template: &template1Name}
	nameA := "COMMAND-A"
	commandA := CommandDefinition{Name: &nameA, Template: &template2Name}

	cfg := Configuration{
		Command: []CommandDefinition{
			template1,
			template2,
			commandA,
		},
	}

	cmd, err := cfg.FindCommandByName("COMMAND-A")
	if err != nil {
		t.Fatalf("Did not expect error from cfg.FindCommandByName, but got %v", err)
	}

	expectedEntryPoint := template1EntryPoint
	cmdEntryPoint, _ := cmd.GetEntryPoint()
	assert.Equal(t, expectedEntryPoint, cmdEntryPoint)
}

func TestConfiguration_SetConfigurationFilePath(t *testing.T) {
	cfg := Configuration{}
	somePath := "/tmp/configpath"
	cfg.SetConfigurationFilePath(somePath)

	assert.Equal(t, somePath, cfg.configFilePath)
}

func TestConfiguration_GetConfigurationFilePath(t *testing.T) {
	cfg := Configuration{}
	somePath := "/tmp/configpath"
	cfg.configFilePath = somePath
	configPath := cfg.GetConfigurationFilePath()

	assert.Equal(t, somePath, configPath)
}
