package config

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type searchableStringSlice []string

var deprecatedFieldNames = searchableStringSlice{"IsDaemon"}

func (sss searchableStringSlice) Contain(needle string) bool {
	for _, v := range sss {
		if v == needle {
			return true
		}
	}

	return false
}

func TestProperty_Unset(t *testing.T) {
	command := CommandDefinition{}
	checkCommandDefinitionGettersUnset(t, command)
}

func checkCommandDefinitionGettersUnset(t *testing.T, command CommandDefinition) {
	val := reflect.ValueOf(command)
	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)

		if deprecatedFieldNames.Contain(typeField.Name) {
			return
		}

		getterName := fmt.Sprintf("Get%s", typeField.Name)
		method := reflect.ValueOf(&command).MethodByName(getterName)

		zero := reflect.Value{}
		if method == zero {
			t.Fatalf("missing getter '%s'", getterName)
		}

		result := method.Call([]reflect.Value{})

		// 2nd return value must be false, since there is no valid configuration value
		ok := result[1].Bool()
		assert.False(t, ok, getterName)
		// returned value does not matter, but should be a valid type
		assert.True(t, result[0].IsValid(), getterName)
	}
}

func TestProperty_Set(t *testing.T) {
	command := getFullFeatureCommandDefinition()
	checkCommandDefinitionGetters(t, command)
}

func checkCommandDefinitionGetters(t *testing.T, command CommandDefinition) {
	val := reflect.ValueOf(command)
	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)

		if deprecatedFieldNames.Contain(typeField.Name) {
			return
		}

		getterName := fmt.Sprintf("Get%s", typeField.Name)
		method := reflect.ValueOf(&command).MethodByName(getterName)

		zero := reflect.Value{}
		if method == zero {
			t.Fatalf("missing getter '%s'", getterName)
		}

		result := method.Call([]reflect.Value{})

		// second return value must be true to indicate there is a valid configuration value
		ok := result[1].Bool()
		assert.True(t, ok, getterName)
		// first return parameter must be the same value as the struct has, but not a pointer anymore
		returnedValue := result[0]
		structValue := val.Field(i).Elem()

		assert.Equal(t, structValue.Interface(), returnedValue.Interface(), getterName)
	}
}

func TestIsTemplateArrayMerged_configNotSet(t *testing.T) {
	commandDef := CommandDefinition{}
	assert.False(t, commandDef.IsTemplateArrayMerged("volumes"))
}

func TestIsTemplateArrayMerged_configIsSetSet(t *testing.T) {
	arrayKeysToBeMerged := []string{"Volumes"}
	commandDef := CommandDefinition{MergeTemplateArrays: &arrayKeysToBeMerged}
	assert.True(t, commandDef.IsTemplateArrayMerged("volumes"))
}

func TestHasTemplate_configNotSet(t *testing.T) {
	commandDef := CommandDefinition{}
	assert.False(t, commandDef.HasTemplate())
}

func TestHasTemplate_configSetButEmpty(t *testing.T) {
	template := ""
	commandDef := CommandDefinition{
		Template: &template,
	}
	assert.False(t, commandDef.HasTemplate())
}

func TestHasTemplate_configSetAndNotEmpty(t *testing.T) {
	template := "some-template"
	commandDef := CommandDefinition{
		Template: &template,
	}
	assert.True(t, commandDef.HasTemplate())
}

func TestHasName_configNotSet(t *testing.T) {
	commandDef := CommandDefinition{}
	assert.False(t, commandDef.HasName())
}

func TestHasName_configSetButEmpty(t *testing.T) {
	name := ""
	commandDef := CommandDefinition{
		Name: &name,
	}
	assert.False(t, commandDef.HasName())
}

func TestHasName_configSetAndNotEmpty(t *testing.T) {
	name := "some-name"
	commandDef := CommandDefinition{
		Name: &name,
	}
	assert.True(t, commandDef.HasName())
}

func TestGetBoolConfigValues(t *testing.T) {
	testCases := map[string]struct {
		initField     func(*bool) CommandDefinition
		getFieldValue func(CommandDefinition) (bool, bool)
	}{
		"AddGroups": {
			initField:     func(b *bool) CommandDefinition { return CommandDefinition{AddGroups: b} },
			getFieldValue: func(d CommandDefinition) (bool, bool) { return d.GetAddGroups() },
		},
		"AutoMountWorkDir": {
			initField:     func(b *bool) CommandDefinition { return CommandDefinition{AutoMountWorkDir: b} },
			getFieldValue: func(d CommandDefinition) (bool, bool) { return d.GetAutoMountWorkDir() },
		},
		"Impersonate": {
			initField:     func(b *bool) CommandDefinition { return CommandDefinition{Impersonate: b} },
			getFieldValue: func(d CommandDefinition) (bool, bool) { return d.GetImpersonate() },
		},
		"IsDetached": {
			initField:     func(b *bool) CommandDefinition { return CommandDefinition{IsDetached: b} },
			getFieldValue: func(d CommandDefinition) (bool, bool) { return d.GetIsDetached() },
		},
		"IsInteractive": {
			initField:     func(b *bool) CommandDefinition { return CommandDefinition{IsInteractive: b} },
			getFieldValue: func(d CommandDefinition) (bool, bool) { return d.GetIsInteractive() },
		},
		"IsTemplate": {
			initField:     func(b *bool) CommandDefinition { return CommandDefinition{IsTemplate: b} },
			getFieldValue: func(d CommandDefinition) (bool, bool) { return d.GetIsTemplate() },
		},
		"RemoveContainer": {
			initField:     func(b *bool) CommandDefinition { return CommandDefinition{RemoveContainer: b} },
			getFieldValue: func(d CommandDefinition) (bool, bool) { return d.GetRemoveContainer() },
		},
		"RequireEnvVars": {
			initField:     func(b *bool) CommandDefinition { return CommandDefinition{RequireEnvVars: b} },
			getFieldValue: func(d CommandDefinition) (bool, bool) { return d.GetRequireEnvVars() },
		},
		"UniqueNames": {
			initField:     func(b *bool) CommandDefinition { return CommandDefinition{UniqueNames: b} },
			getFieldValue: func(d CommandDefinition) (bool, bool) { return d.GetUniqueNames() },
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			commandDef := testCase.initField(boolP(false))
			v, ok := testCase.getFieldValue(commandDef)
			assert.True(t, ok)
			assert.False(t, v)

			commandDef = testCase.initField(boolP(true))
			v, ok = testCase.getFieldValue(commandDef)
			assert.True(t, ok)
			assert.True(t, v)

			commandDef = testCase.initField(nil)
			v, ok = testCase.getFieldValue(commandDef)
			assert.False(t, ok)
			assert.False(t, v)
		})
	}
}

func boolP(b bool) *bool {
	return &b
}

func TestGetStringConfigValues(t *testing.T) {
	testCases := map[string]struct {
		initField     func(*string) CommandDefinition
		getFieldValue func(CommandDefinition) (string, bool)
	}{
		"Template": {
			initField:     func(s *string) CommandDefinition { return CommandDefinition{Template: s} },
			getFieldValue: func(d CommandDefinition) (string, bool) { return d.GetTemplate() },
		},
		"EntryPoint": {
			initField:     func(s *string) CommandDefinition { return CommandDefinition{EntryPoint: s} },
			getFieldValue: func(d CommandDefinition) (string, bool) { return d.GetEntryPoint() },
		},
		"Command": {
			initField:     func(s *string) CommandDefinition { return CommandDefinition{Command: s} },
			getFieldValue: func(d CommandDefinition) (string, bool) { return d.GetCommand() },
		},
		"Name": {
			initField:     func(s *string) CommandDefinition { return CommandDefinition{Name: s} },
			getFieldValue: func(d CommandDefinition) (string, bool) { return d.GetName() },
		},
		"Image": {
			initField:     func(s *string) CommandDefinition { return CommandDefinition{Image: s} },
			getFieldValue: func(d CommandDefinition) (string, bool) { return d.GetImage() },
		},
		"Network": {
			initField:     func(s *string) CommandDefinition { return CommandDefinition{Network: s} },
			getFieldValue: func(d CommandDefinition) (string, bool) { return d.GetNetwork() },
		},
		"EnvFile": {
			initField:     func(s *string) CommandDefinition { return CommandDefinition{EnvFile: s} },
			getFieldValue: func(d CommandDefinition) (string, bool) { return d.GetEnvFile() },
		},
		"IP": {
			initField:     func(s *string) CommandDefinition { return CommandDefinition{IP: s} },
			getFieldValue: func(d CommandDefinition) (string, bool) { return d.GetIP() },
		},
		"WorkDir": {
			initField:     func(s *string) CommandDefinition { return CommandDefinition{WorkDir: s} },
			getFieldValue: func(d CommandDefinition) (string, bool) { return d.GetWorkDir() },
		},
	}

	const want = "TEST"

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			commandDef := testCase.initField(stringP(want))
			v, ok := testCase.getFieldValue(commandDef)
			assert.True(t, ok)
			assert.Equal(t, want, v)

			commandDef = testCase.initField(nil)
			v, ok = testCase.getFieldValue(commandDef)
			assert.False(t, ok)
			assert.Equal(t, "", v)
		})
	}
}

func stringP(s string) *string {
	return &s
}

func TestGetStringSliceConfigValues(t *testing.T) {
	testCases := map[string]struct {
		initField     func(*[]string) CommandDefinition
		getFieldValue func(CommandDefinition) ([]string, bool)
	}{
		"PortsFromParams": {
			initField:     func(s *[]string) CommandDefinition { return CommandDefinition{PortsFromParams: s} },
			getFieldValue: func(d CommandDefinition) ([]string, bool) { return d.GetPortsFromParams() },
		},
		"MergeTemplateArrays": {
			initField:     func(s *[]string) CommandDefinition { return CommandDefinition{MergeTemplateArrays: s} },
			getFieldValue: func(d CommandDefinition) ([]string, bool) { return d.GetMergeTemplateArrays() },
		},
		"AdditionalArgs": {
			initField:     func(s *[]string) CommandDefinition { return CommandDefinition{AdditionalArgs: s} },
			getFieldValue: func(d CommandDefinition) ([]string, bool) { return d.GetAdditionalArgs() },
		},
		"Volumes": {
			initField:     func(s *[]string) CommandDefinition { return CommandDefinition{Volumes: s} },
			getFieldValue: func(d CommandDefinition) ([]string, bool) { return d.GetVolumes() },
		},
		"Links": {
			initField:     func(s *[]string) CommandDefinition { return CommandDefinition{Links: s} },
			getFieldValue: func(d CommandDefinition) ([]string, bool) { return d.GetLinks() },
		},
		"EnvVars": {
			initField:     func(s *[]string) CommandDefinition { return CommandDefinition{EnvVars: s} },
			getFieldValue: func(d CommandDefinition) ([]string, bool) { return d.GetEnvVars() },
		},
		"Ports": {
			initField:     func(s *[]string) CommandDefinition { return CommandDefinition{Ports: s} },
			getFieldValue: func(d CommandDefinition) ([]string, bool) { return d.GetPorts() },
		},
	}

	const want = "TEST"

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			commandDef := testCase.initField(&[]string{"TEST"})
			v, ok := testCase.getFieldValue(commandDef)
			assert.True(t, ok)
			assert.Equal(t, want, v[0])

			commandDef = testCase.initField(nil)
			v, ok = testCase.getFieldValue(commandDef)
			assert.False(t, ok)
			assert.Equal(t, []string{}, v)
		})
	}
}

func TestGet2DStringSliceConfigValues(t *testing.T) {
	testCases := map[string]struct {
		initField     func(*[][]string) CommandDefinition
		getFieldValue func(CommandDefinition) ([][]string, bool)
	}{
		"ReplaceArgs": {
			initField:     func(s *[][]string) CommandDefinition { return CommandDefinition{ReplaceArgs: s} },
			getFieldValue: func(d CommandDefinition) ([][]string, bool) { return d.GetReplaceArgs() },
		},
	}

	const want = "TEST"

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			commandDef := testCase.initField(&[][]string{{"TEST"}})
			v, ok := testCase.getFieldValue(commandDef)
			assert.True(t, ok)
			assert.Equal(t, want, v[0][0])

			commandDef = testCase.initField(nil)
			v, ok = testCase.getFieldValue(commandDef)
			assert.False(t, ok)
			assert.Equal(t, [][]string{}, v)
		})
	}
}
