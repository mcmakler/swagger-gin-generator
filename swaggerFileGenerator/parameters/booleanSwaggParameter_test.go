package parameters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoolSwaggParameter_ToString(t *testing.T) {
	t.Run("Test: BoolSwaggParameter.ToString()", func(t *testing.T) {
		t.Run("Should: return string with empty params", func(t *testing.T) {
			a := &boolSwaggParameter{
				params: nil,
			}
			expected := typeDeficeString + boolType
			actual, _ := a.ToString(false)
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return string with empty params", func(t *testing.T) {
			params := map[string]interface{}{
				"test": 1,
			}
			a := &boolSwaggParameter{
				params: params,
			}
			expected := errorEmptyIn
			_, actual := a.ToString(false)
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return string with empty params", func(t *testing.T) {
			params := map[string]interface{}{
				"in": "in",
			}
			a := &boolSwaggParameter{
				params: params,
			}
			expected := errorEmptyName
			_, actual := a.ToString(false)
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return string with all params", func(t *testing.T) {
			params := map[string]interface{}{
				"in":              "in",
				"name":            "name",
				"required":        true,
				"allowEmptyValue": false,
				"description":     "description",
			}
			a := &boolSwaggParameter{
				params: params,
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
		expect := &boolSwaggParameter{params:nil}
		actual := NewBoolSwagParameter(nil)
		assert.Equal(t, expect, actual)
	})
}
