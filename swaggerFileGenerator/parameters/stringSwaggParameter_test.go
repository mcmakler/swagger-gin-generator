package parameters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringSwaggParameter_ToString(t *testing.T) {
	t.Run("Test: StringSwaggParameter.ToString()", func(t *testing.T) {
		t.Run("Should: return string with empty params", func(t *testing.T) {
			a := &stringSwaggParameter{
				params: nil,
			}
			expected := typeDeficeString + stringType
			actual, _ := a.ToString(false)
			assert.Equal(t, expected, actual)
		})
		t.Run("Should: return string with empty params", func(t *testing.T) {
			params := map[string]interface{}{
				"test": 1,
			}
			a := &stringSwaggParameter{
				params: params,
			}
			expected := typeDeficeString + stringType
			actual, _ := a.ToString(false)
			assert.Equal(t, expected, actual)
		})
		t.Run("Should: return string with all params", func(t *testing.T) {
			params := map[string]interface{}{
				"in":              "in",
				"name":            "name",
				"required":        true,
				"format":          "date",
				"minLength":       79,
				"maxLength":       239,
				"pattern":         "pattern",
				"allowEmptyValue": false,
				"description":     "description",
				"enum":            []string{"EIN", "ZWEI", "DREI"},
			}
			a := &stringSwaggParameter{
				params: params,
			}
			expected := typeDeficeString + stringType +
				inString + "in" +
				nameString + "name" +
				requiredString + "true" +
				formatString + "date" +
				minLengthString + "79" +
				maxLengthString + "239" +
				patternString + "pattern" +
				allowEmptyValueString + "false" +
				descriptionString + "description" +
				enumString + enumNewString + "EIN" + enumNewString + "ZWEI" + enumNewString + "DREI"
			actual, _ := a.ToString(false)
			assert.Equal(t, expected, actual)
		})
	})
}

func TestNewStringSwagParameter(t *testing.T) {
	t.Run("Test: NewStringSwagParameter", func(t *testing.T) {
		expect := &stringSwaggParameter{params:nil}
		actual := NewStringSwagParameter(nil)
		assert.Equal(t, expect, actual)
	})
}
