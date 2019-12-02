package parameters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumberSwaggParameter_ToString(t *testing.T) {
	t.Run("Test: NumberSwaggParameter.ToString()", func(t *testing.T) {
		t.Run("Should: return string with empty params", func(t *testing.T) {
			a := &numberSwaggParameter{
				params: nil,
			}
			expected := typeDeficeString + numberType
			actual, _ := a.ToString(false)
			assert.Equal(t, expected, actual)
		})
		t.Run("Should: return string with empty params", func(t *testing.T) {
			params := map[string]interface{}{
				"test": 1,
			}
			a := &numberSwaggParameter{
				params: params,
			}
			expected := typeDeficeString + numberType
			actual, _ := a.ToString(false)
			assert.Equal(t, expected, actual)
		})
		t.Run("Should: return string with all params", func(t *testing.T) {
			params := map[string]interface{}{
				"in":                    "in",
				"name":                  "name",
				"required":              true,
				"defaultValue":          23.9,
				"minimumValue":          -14203.0399394,
				"exclusiveMinimumValue": true,
				"maximumValue":          12321.44623412345,
				"exclusiveMaximumValue": false,
				"multipleOf":            7.9,
				"allowEmptyValue":       false,
				"description":           "description",
			}
			a := &numberSwaggParameter{
				params: params,
			}
			expected := typeDeficeString + numberType +
				inString + "in" +
				nameString + "name" +
				requiredString + "true" +
				defaultValueString + "23.9" +
				minimumValueString + "-14203.0399394" +
				exclusiveMinimumValueString + "true" +
				maximumValueString + "12321.44623412345" +
				exclusiveMaximumValueString + "false" +
				multipleOfString + "7.9" +
				allowEmptyValueString + "false" +
				descriptionString + "description"
			actual, _ := a.ToString(false)
			assert.Equal(t, expected, actual)
		})
	})
}

func TestNewNumberSwagParameter(t *testing.T) {
	t.Run("Test: NewNumberSwagParameter", func(t *testing.T) {
		expect := &numberSwaggParameter{params:nil}
		actual := NewNumberSwagParameter(nil)
		assert.Equal(t, expect, actual)
	})
}

