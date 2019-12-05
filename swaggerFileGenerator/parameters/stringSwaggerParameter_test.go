package parameters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringSwaggerParameter_ToString(t *testing.T) {
	t.Run("Test: StringSwaggerParameter.ToString()", func(t *testing.T) {
		t.Run("Should: return with empty config", func(t *testing.T) {
			a := NewStringSwaggerParameter(nil)
			actual, err := a.ToString()
			assert.NoError(t, err)

			expected := typeString + stringType
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return with full config", func(t *testing.T) {
			config := map[string]interface{}{
				"format":    "date",
				"minLength": 79,
				"maxLength": 239,
				"pattern":   "pattern",
				"enum":      []string{"EIN", "ZWEI", "DREI"},
			}
			a := NewStringSwaggerParameter(config)
			actual, err := a.ToString()
			assert.NoError(t, err)

			expected := typeString + stringType +
				formatString + "date" +
				minLengthString + "79" +
				maxLengthString + "239" +
				patternString + "pattern" +
				enumString + enumNewString + "EIN" + enumNewString + "ZWEI" + enumNewString + "DREI"
			assert.Equal(t, expected, actual)
		})
	})
}

func TestStringSwaggerParameter_IsObject(t *testing.T) {
	t.Run("Test: StringSwaggerParameter.IsObject()", func(t *testing.T) {
		t.Run("Should: return false", func(t *testing.T) {
			a := NewStringSwaggerParameter(nil)
			actual := a.IsObject()
			assert.False(t, actual)
		})
	})
}

func TestStringSwaggerParameter_getConfigs(t *testing.T) {
	t.Run("Test: StringSwaggerParameter.getConfigs()", func(t *testing.T) {
		t.Run("Should: return nil", func(t *testing.T) {
			a := NewStringSwaggerParameter(nil)
			actual := a.getConfigs()
			assert.Nil(t, actual)
		})
	})
}
