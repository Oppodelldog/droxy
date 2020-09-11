package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	commandNameA = "COMMAND-A"
	commandNameX = "COMMAND-X"
	commandNameC = "COMMAND-Z"
)

func TestConfiguration_FindCommandByName(t *testing.T) {
	nameA := commandNameA
	nameZ := commandNameC
	commandA := CommandDefinition{Name: &nameA, Command: stringP("A")}
	commandA1 := CommandDefinition{Name: &nameA, Command: stringP("A1")}
	commandZ := CommandDefinition{Name: &nameZ}

	cfg := Configuration{
		osNameMatcher: defaultOSNameMatcher,
		Command: []CommandDefinition{
			commandA,
			commandA1,
			commandZ,
		},
	}

	commandDef, err := cfg.FindCommandByName(commandNameA)
	if err != nil {
		t.Fatalf("Did not expect cgf.FindCommandByName to return an error, but got: %v", err)
	}

	assert.Equal(t, commandA1, commandDef)
}

func TestConfiguration_FindCommandByName_OSSpecific(t *testing.T) {
	name := commandNameA
	command1 := CommandDefinition{Name: &name, OS: stringP("windows")}
	command2 := CommandDefinition{Name: &name, Command: stringP("linux")}

	cfg := Configuration{
		osNameMatcher: defaultOSNameMatcher,
		Command: []CommandDefinition{
			command1,
			command2,
		},
	}

	commandDef, err := cfg.FindCommandByName(commandNameA)
	if err != nil {
		t.Fatalf("Did not expect cgf.FindCommandByName to return an error, but got: %v", err)
	}

	assert.Equal(t, command2, commandDef)
}

func TestConfiguration_FindCommandByName_OSSpecific_Templating(t *testing.T) {
	tplName := stringP("base1")
	commandName := stringP("command")
	baseTplName := stringP("base")
	baseCommand := stringP("base-cmd")
	baseEntryPoint := stringP("base-entry-point")
	windowsCommand := stringP("windows-cmd")
	linuxCommand := stringP("linux-cmd")

	template1 := CommandDefinition{Name: baseTplName, Command: baseCommand, EntryPoint: baseEntryPoint}
	template2 := CommandDefinition{Name: tplName, Template: baseTplName, OS: stringP("windows"), Command: windowsCommand}
	template3 := CommandDefinition{Name: tplName, Template: baseTplName, OS: stringP("linux"), Command: linuxCommand}

	command := CommandDefinition{Name: commandName, Template: tplName, OS: stringP("linux")}

	cfg := Configuration{
		osNameMatcher: fakeOSMatcher("linux"),
		Command: []CommandDefinition{
			template1,
			template2,
			template3,
			command,
		},
	}

	commandDef, err := cfg.FindCommandByName("command")
	if err != nil {
		t.Fatalf("Did not expect cgf.FindCommandByName to return an error, but got: %v", err)
	}

	assert.Equal(t, CommandDefinition{
		OS:         nil,
		EntryPoint: baseEntryPoint,
		Command:    linuxCommand,
		Name:       commandName,
	}, commandDef)
}

func TestConfiguration_FindCommandByName_NotFoundError(t *testing.T) {
	nameA := commandNameA
	commandA := CommandDefinition{Name: &nameA}

	cfg := Configuration{
		osNameMatcher: defaultOSNameMatcher,
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
		osNameMatcher: defaultOSNameMatcher,
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
		osNameMatcher: defaultOSNameMatcher,
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
		osNameMatcher: defaultOSNameMatcher,
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
	cfg.ConfigFilePath = somePath

	assert.Equal(t, somePath, cfg.ConfigFilePath)
}

func TestConfiguration_GetConfigurationFilePath(t *testing.T) {
	cfg := Configuration{}
	somePath := "/tmp/configPath"
	cfg.ConfigFilePath = somePath
	configPath := cfg.GetConfigurationFilePath()

	assert.Equal(t, somePath, configPath)
}

func stringP(s string) *string {
	return &s
}

func fakeOSMatcher(fakeCurrentOS string) func(string) bool {
	return func(s string) bool {
		return fakeCurrentOS == s
	}
}
