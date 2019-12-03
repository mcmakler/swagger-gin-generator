package parameters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringSwaggParameter_ToString(t *testing.T) {
	t.Run("Test: StringSwaggParameter.ToString()", func(t *testing.T) {
		t.Run("Should: return string with empty configs", func(t *testing.T) {
			a := &stringSwaggParameter{
				configs: nil,
			}
			expected := typeDeficeString + stringType
			actual, _ := a.ToString(false)
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return string with empty configs", func(t *testing.T) {
			a := &stringSwaggParameter{
				configs: nil,
			}
			expected := typeString + stringType
			actual, _ := a.ToString(true)
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return string with empty configs", func(t *testing.T) {
			params := map[string]interface{}{
				"test": 1,
			}
			a := &stringSwaggParameter{
				configs: params,
			}
			expected := errorEmptyIn
			_, actual := a.ToString(false)
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return string with empty configs", func(t *testing.T) {
			params := map[string]interface{}{
				"in": "in",
			}
			a := &stringSwaggParameter{
				configs: params,
			}
			expected := errorEmptyName
			_, actual := a.ToString(false)
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return string with all configs", func(t *testing.T) {
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
				configs: params,
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
		expect := &stringSwaggParameter{configs: nil}
		actual := NewStringSwagParameter(nil)
		assert.Equal(t, expect, actual)
	})
}
