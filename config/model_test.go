package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const commandNameA = "COMMAND-A"
const commandNameX = "COMMAND-X"
const commandNameC = "COMMAND-Z"

func TestConfiguration_FindCommandByName(t *testing.T) {
	nameA := commandNameA
	commandA := CommandDefinition{Name: &nameA}
	nameX := commandNameX
	commandX := CommandDefinition{Name: &nameX}
	nameZ := commandNameC
	commandZ := CommandDefinition{Name: &nameZ}

	cfg := Configuration{
		Command: []CommandDefinition{
			commandA,
			commandX,
			commandZ,
		},
	}

	commandDef, err := cfg.FindCommandByName(commandNameX)
	if err != nil {
		t.Fatalf("Did not expect cgf.FindCommandByName to return an error, but got: %v", err)
	}

	assert.Equal(t, commandX, commandDef)
}

func TestConfiguration_FindCommandByName_NotFoundError(t *testing.T) {
	nameA := commandNameA
	commandA := CommandDefinition{Name: &nameA}

	cfg := Configuration{
		Command: []CommandDefinition{
			commandA,
		},
	}

	_, err := cfg.FindCommandByName(commandNameX)

	assert.Error(t, err)
}

func TestConfiguration_FindCommandByName_ResolvesTemplate(t *testing.T) {
	templateName := "TEMPLATE-A"
	isTemplate := true
	templateNetwork := "templateNetwork-by-template"
	templateCommand := CommandDefinition{Name: &templateName, IsTemplate: &isTemplate, Network: &templateNetwork}

	nameA := commandNameA
	commandA := CommandDefinition{Name: &nameA, Template: &templateName}

	cfg := Configuration{
		Command: []CommandDefinition{
			commandA,
			templateCommand,
		},
	}

	commandDef, err := cfg.FindCommandByName(commandNameA)
	if err != nil {
		t.Fatalf("Did not expect cgf.FindCommandByName to return an error, but got: %v", err)
	}

	assert.Equal(t, &templateNetwork, commandDef.Network)
}

func TestConfiguration_FindCommandByName_TemplateNotFoundError(t *testing.T) {
	templateName := "yes-template-does-not-exist"
	nameA := commandNameA
	commandA := CommandDefinition{Name: &nameA, Template: &templateName}

	cfg := Configuration{
		Command: []CommandDefinition{
			commandA,
		},
	}

	_, err := cfg.FindCommandByName(commandNameA)

	assert.Error(t, err)
}

func TestConfiguration_FindCommandByName_TemplateHasTemplate(t *testing.T) {
	template1Name := "template1"
	template1EntryPoint := "template1EntryPoint"
	template2Name := "template2"
	template1 := CommandDefinition{Name: &template1Name, EntryPoint: &template1EntryPoint}
	template2 := CommandDefinition{Name: &template2Name, Template: &template1Name}
	nameA := commandNameA
	commandA := CommandDefinition{Name: &nameA, Template: &template2Name}

	cfg := Configuration{
		Command: []CommandDefinition{
			template1,
			template2,
			commandA,
		},
	}

	cmd, err := cfg.FindCommandByName(commandNameA)
	if err != nil {
		t.Fatalf("Did not expect error from cfg.FindCommandByName, but got %v", err)
	}

	expectedEntryPoint := template1EntryPoint
	cmdEntryPoint, _ := cmd.GetEntryPoint()
	assert.Equal(t, expectedEntryPoint, cmdEntryPoint)
}

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
	var (
		base    *[][]string
		overlay = &[][]string{
			{"a", "b"},
		}
	)

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

func TestConfiguration_SetConfigurationFilePath(t *testing.T) {
	cfg := Configuration{}
	somePath := "/tmp/configPath"
	cfg.SetConfigurationFilePath(somePath)

	assert.Equal(t, somePath, cfg.configFilePath)
}

func TestConfiguration_GetConfigurationFilePath(t *testing.T) {
	cfg := Configuration{}
	somePath := "/tmp/configPath"
	cfg.configFilePath = somePath
	configPath := cfg.GetConfigurationFilePath()

	assert.Equal(t, somePath, configPath)
}
