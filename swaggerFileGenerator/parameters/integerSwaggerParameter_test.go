package parameters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegerSwaggParameter_ToString(t *testing.T) {
	t.Run("Test: IntegerSwaggerParameter.ToString()", func(t *testing.T) {
		t.Run("Should: return with empty config", func(t *testing.T) {
			a := NewIntegerSwaggerParameter(nil)
			actual, err := a.ToString()
			assert.NoError(t, err)

			expected := typeString + integerType
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return with full config", func(t *testing.T) {
			config := map[string]interface{}{
				"defaultValue":          10,
				"minimumValue":          -2000000000,
				"exclusiveMinimumValue": true,
				"maximumValue":          2000000000,
				"exclusiveMaximumValue": false,
				"multipleOf":            4,
			}
			a := NewIntegerSwaggerParameter(config)
			actual, err := a.ToString()
			assert.NoError(t, err)

			expected := typeString + integerType +
				defaultValueString + "10" +
				minimumValueString + "-2000000000" +
				exclusiveMinimumValueString + "true" +
				maximumValueString + "2000000000" +
				exclusiveMaximumValueString + "false" +
				multipleOfString + "4"
			assert.Equal(t, expected, actual)
		})
	})
}

func TestIntegerSwaggerParameter_IsObject(t *testing.T) {
	t.Run("Test: IntegerSwaggerParameter.IsObject()", func(t *testing.T) {
		t.Run("Should: return false", func(t *testing.T) {
			a := NewIntegerSwaggerParameter(nil)
			actual := a.IsObject()
			assert.False(t, actual)
		})
	})
}

func TestIntegerSwaggerParameter_getConfigs(t *testing.T) {
	t.Run("Test: IntegerSwaggerParameter.getConfigs()", func(t *testing.T) {
		t.Run("Should: return nil", func(t *testing.T) {
			a := NewIntegerSwaggerParameter(nil)
			actual := a.getConfigs()
			assert.Nil(t, actual)
		})
	})
}
