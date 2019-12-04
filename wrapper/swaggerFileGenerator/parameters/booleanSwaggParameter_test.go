package parameters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoolSwaggParameter_ToString(t *testing.T) {
	t.Run("Test: BoolSwaggParameter.ToString()", func(t *testing.T) {
		t.Run("Should: return string with empty configs", func(t *testing.T) {
			a := &boolSwaggParameter{
				configs: nil,
			}
			expected := typeDeficeString + boolType
			actual, _ := a.ToString(false)
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return string with empty configs", func(t *testing.T) {
			params := map[string]interface{}{
				"test": 1,
			}
			a := &boolSwaggParameter{
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
			a := &boolSwaggParameter{
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
				"allowEmptyValue": false,
				"description":     "description",
			}
			a := &boolSwaggParameter{
				configs: params,
			}
			expected := typeDeficeString + boolType +
				inString + "in" +
				nameString + "name" +
				requiredString + "true" +
				allowEmptyValueString + "false" +
				descriptionString + "description"
			actual, _ := a.ToString(false)
			assert.Equal(t, expected, actual)
		})
	})
}

func TestNewBoolSwagParameter(t *testing.T) {
	t.Run("Test: NewBoolSwagParameter", func(t *testing.T) {
		expect := &boolSwaggParameter{configs: nil}
		actual := NewBoolSwagParameter(nil)
		assert.Equal(t, expect, actual)
	})
}
