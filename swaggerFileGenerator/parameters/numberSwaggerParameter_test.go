package parameters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumberSwaggerParameter_ToString(t *testing.T) {
	t.Run("Test: NumberSwaggerParameter.ToString()", func(t *testing.T) {
		t.Run("Should: return with empty config", func(t *testing.T) {
			a := NewNumberSwaggerParameter(nil)
			actual, err := a.ToString()
			assert.NoError(t, err)

			expected := typeString + numberType
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return with full config", func(t *testing.T) {
			config := map[string]interface{}{
				"defaultValue":          23.9,
				"minimumValue":          -14203.0399394,
				"exclusiveMinimumValue": true,
				"maximumValue":          12321.44623412345,
				"exclusiveMaximumValue": false,
				"multipleOf":            7.9,
			}
			a := NewNumberSwaggerParameter(config)
			actual, err := a.ToString()
			assert.NoError(t, err)

			expected := typeString + numberType +
				defaultValueString + "23.9" +
				minimumValueString + "-14203.0399394" +
				exclusiveMinimumValueString + "true" +
				maximumValueString + "12321.44623412345" +
				exclusiveMaximumValueString + "false" +
				multipleOfString + "7.9"
			assert.Equal(t, expected, actual)
		})
	})
}

func TestNumberSwaggerParameter_IsObject(t *testing.T) {
	t.Run("Test: NumberSwaggerParameter.IsObject()", func(t *testing.T) {
		t.Run("Should: return false", func(t *testing.T) {
			a := NewNumberSwaggerParameter(nil)
			actual := a.IsObject()
			assert.False(t, actual)
		})
	})
}

func TestNumberSwaggerParameter_getConfigs(t *testing.T) {
	t.Run("Test: NumberSwaggerParameter.getConfigs()", func(t *testing.T) {
		t.Run("Should: return nil", func(t *testing.T) {
			a := NewNumberSwaggerParameter(nil)
			actual := a.getConfigs()
			assert.Nil(t, actual)
		})
	})
}
