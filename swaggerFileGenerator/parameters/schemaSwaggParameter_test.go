package parameters

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestSchemaSwaggerParameter_ToString(t *testing.T) {
	t.Run("Test: SchemaSwaggerParameter.ToString()", func(t *testing.T) {
		t.Run("Should: return error "+errorEmptyIn.Error(), func(t *testing.T) {
			params := NewBoolSwaggerParameter(nil)
			a := NewSchemaSwaggParameter(params)
			_, actual := a.ToString()
			expected := errorEmptyIn
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return error "+errorEmptyIn.Error(), func(t *testing.T) {
			config := map[string]interface{}{}
			parameter := NewBoolSwaggerParameter(config)
			a := NewSchemaSwaggParameter(parameter)
			_, actual := a.ToString()
			expected := errorEmptyIn
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return error "+errorEmptyName.Error(), func(t *testing.T) {
			config := map[string]interface{}{
				"in": "body",
			}
			parameter := NewBoolSwaggerParameter(config)
			a := NewSchemaSwaggParameter(parameter)
			_, actual := a.ToString()

			expected := errorEmptyName
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return error "+ErrorNillItemsParameter.Error(), func(t *testing.T) {
			config := map[string]interface{}{
				"in":   "body",
				"name": "name",
			}
			parameter := NewArraySwaggerParameter(config, nil)
			a := NewSchemaSwaggParameter(parameter)
			_, actual := a.ToString()

			expected := ErrorNillItemsParameter
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return with config", func(t *testing.T) {
			config := map[string]interface{}{
				"in":   "body",
				"name": "name",
			}
			parameter := NewBoolSwaggerParameter(config)
			a := NewSchemaSwaggParameter(parameter)
			actual, err := a.ToString()
			assert.NoError(t, err)

			str, _ := parameter.ToString()
			expected := inDeficeString + "body" +
				nameString + "name" +
				linkOnSchemaString + strings.ReplaceAll(str, "\n", "\n  ")
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return with config", func(t *testing.T) {
			config := map[string]interface{}{
				"in":       "header",
				"name":     "name",
				"required": true,
			}
			parameter := NewBoolSwaggerParameter(config)
			a := NewSchemaSwaggParameter(parameter)
			actual, err := a.ToString()
			assert.NoError(t, err)

			str, _ := parameter.ToString()
			expected := inDeficeString + "header" +
				nameString + "name" +
				requiredString + "true" +
				str
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return with full config", func(t *testing.T) {
			config := map[string]interface{}{
				"in":             "header",
				"name":           "name",
				"description":    "description",
				"nameOfVariable": "parameter",
			}
			parameter := NewBoolSwaggerParameter(config)
			a := NewSchemaSwaggParameter(parameter)
			actual, err := a.ToString()
			assert.NoError(t, err)

			expected := inDeficeString + "header" +
				nameString + "name" +
				descriptionString + "description" +
				linkOnSchemaString + refString + "parameter'"
			assert.Equal(t, expected, actual)
		})
	})
}

func TestSchemaSwaggerParameter_IsObject(t *testing.T) {
	t.Run("Test: SchemaSwaggerParameter.IsObject()", func(t *testing.T) {
		t.Run("Should: return true", func(t *testing.T) {
			config := map[string]interface{}{
				"in":   "body",
				"name": "name",
			}
			parameter := NewBoolSwaggerParameter(config)
			a := NewSchemaSwaggParameter(parameter)
			assert.False(t, a.IsObject())
		})

		t.Run("Should: return true", func(t *testing.T) {
			config := map[string]interface{}{
				"in":             "body",
				"name":           "name",
				"nameOfVariable": "parameter",
			}
			parameter := NewBoolSwaggerParameter(config)
			a := NewSchemaSwaggParameter(parameter)
			assert.False(t, a.IsObject())
		})
	})
}

func TestSchemaSwaggerParameter_getConfigs(t *testing.T) {
	t.Run("Test: SchemaSwaggerParameter.getConfigs()", func(t *testing.T) {
		t.Run("Should: return config with nameOfVariable", func(t *testing.T) {
			expected := map[string]interface{}{
				"in":   "body",
				"name": "name",
			}
			parameter := NewBoolSwaggerParameter(expected)
			a := NewSchemaSwaggParameter(parameter)
			actual := a.getConfigs()
			assert.Equal(t, expected, actual)
		})
	})
}
