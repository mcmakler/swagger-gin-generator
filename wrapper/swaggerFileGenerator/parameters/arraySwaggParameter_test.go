package parameters

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestArraySwaggParameter_ToString(t *testing.T) {
	t.Run("Test: IntegerSwaggParameter.ToString()", func(t *testing.T) {
		t.Run("Should: return error empty items", func(t *testing.T) {
			a := &arraySwaggParameter{
				configs: nil,
				items:   nil,
			}
			_, expectedError := a.ToString(false)
			assert.Equal(t, expectedError, ErrorNillItemsParameter)
		})

		t.Run("Should: return error empty Object name", func(t *testing.T) {
			a := &arraySwaggParameter{
				configs: nil,
				items:   NewObjectSwaggerParameter(nil, nil, false),
			}
			expected := errorEmptyIn
			_, actual := a.ToString(false)
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return string with empty configs", func(t *testing.T) {
			a := &arraySwaggParameter{
				configs: nil,
				items:   NewStringSwagParameter(nil),
			}
			expected := errorEmptyIn
			_, actual := a.ToString(false)
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return string with empty configs", func(t *testing.T) {
			params := map[string]interface{}{
				"in": "in",
			}
			a := &arraySwaggParameter{
				configs: params,
				items:   NewBoolSwagParameter(nil),
			}
			expected := errorEmptyName
			_, actual := a.ToString(false)
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return string with empty configs", func(t *testing.T) {
			params := map[string]interface{}{
				"in":   "in",
				"name": "name",
			}
			a := &arraySwaggParameter{
				configs: params,
				items:   NewArraySwaggParameter(nil, nil),
			}
			expected := ErrorNillItemsParameter
			_, actual := a.ToString(false)
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return string with empty configs", func(t *testing.T) {
			params := map[string]interface{}{
				"test": 1,
			}
			a := &arraySwaggParameter{
				configs: params,
				items:   NewBoolSwagParameter(nil),
			}
			str, _ := a.items.ToString(true)
			expected := typeString + arrayType + itemsString + strings.Replace(str, "\n", "\n  ", -1)
			actual, _ := a.ToString(true)
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return string with all configs", func(t *testing.T) {
			params := map[string]interface{}{
				"in":              "in",
				"name":            "name",
				"required":        true,
				"minItems":        79,
				"maxItems":        239,
				"uniqueItems":     false,
				"allowEmptyValue": false,
				"description":     "description",
				"enum":            []string{"EIN", "ZWEI", "DREI"},
			}
			a := &arraySwaggParameter{
				configs: params,
				items:   NewIntegerSwagParameter(nil),
			}
			str, _ := a.items.ToString(true)
			expected := typeDeficeString + arrayType +
				inString + "in" +
				nameString + "name" +
				requiredString + "true" +
				minItemsString + "79" +
				maxItemsString + "239" +
				uniqueItemsString + "false" +
				allowEmptyValueString + "false" +
				descriptionString + "description" +
				itemsString + strings.Replace(str, "\n", "\n  ", -1)
			actual, _ := a.ToString(false)
			assert.Equal(t, expected, actual)
		})
	})
}

func TestNewArraySwaggParameter(t *testing.T) {
	t.Run("Test: NewArraySwaggParameter", func(t *testing.T) {
		expect := &arraySwaggParameter{
			configs: nil,
			items:   NewBoolSwagParameter(nil),
		}
		actual := NewArraySwaggParameter(nil, NewBoolSwagParameter(nil))
		assert.Equal(t, expect, actual)
	})
}
