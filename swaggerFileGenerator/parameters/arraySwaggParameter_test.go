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
				params: nil,
				items:  nil,
			}
			_, expectedError := a.ToString(false)
			assert.Equal(t, expectedError, ErrorNillItemsParameter)
		})
		t.Run("Should: return error empty Object name", func(t *testing.T) {
			a := &arraySwaggParameter{
				params: nil,
				items:  NewObjectSwaggerParameter(nil, nil),
			}
			_, expectedError := a.ToString(false)
			assert.Equal(t, expectedError, errorNilObjectName)
		})
		t.Run("Should: return string with empty params", func(t *testing.T) {
			a := &arraySwaggParameter{
				params: nil,
				items:  NewStringSwagParameter(nil),
			}
			str, _ := a.items.ToString(true)
			expected := typeDeficeString + arrayType + itemsString + strings.Replace(str, "\n", "\n  ", -1)
			actual, _ := a.ToString(false)
			assert.Equal(t, expected, actual)
		})
		t.Run("Should: return string with empty params", func(t *testing.T) {
			params := map[string]interface{}{
				"test": 1,
			}
			a := &arraySwaggParameter{
				params: params,
				items:  NewBoolSwagParameter(nil),
			}
			str, _ := a.items.ToString(true)
			expected := typeDeficeString + arrayType + itemsString + strings.Replace(str, "\n", "\n  ", -1)
			actual, _ := a.ToString(false)
			assert.Equal(t, expected, actual)
		})
		t.Run("Should: return string with all params", func(t *testing.T) {
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
				params: params,
				items:  NewIntegerSwagParameter(nil),
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
			params: nil,
			items:  NewBoolSwagParameter(nil),
		}
		actual := NewArraySwaggParameter(nil, NewBoolSwagParameter(nil))
		assert.Equal(t, expect, actual)
	})
}
