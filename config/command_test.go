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
