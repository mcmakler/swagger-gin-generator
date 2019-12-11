package parameters

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestArraySwaggerParameter_ToString(t *testing.T) {
	t.Run("Test: ArraySwaggerParameter.ToString()", func(t *testing.T) {
		t.Run("Should: return error "+ErrorNillItemsParameter.Error(), func(t *testing.T) {
			a := NewArraySwaggerParameter(nil, nil)
			_, actual := a.ToString()
			expected := ErrorNillItemsParameter
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return error "+ErrorNillItemsParameter.Error(), func(t *testing.T) {
			a := NewArraySwaggerParameter(nil, NewArraySwaggerParameter(nil, nil))
			_, actual := a.ToString()
			expected := ErrorNillItemsParameter
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return with full config", func(t *testing.T) {
			config := map[string]interface{}{
				"minItems":    79,
				"maxItems":    239,
				"uniqueItems": false,
			}
			items := NewBoolSwaggerParameter(nil)
			a := NewArraySwaggerParameter(config, items)
			actual, err := a.ToString()
			assert.NoError(t, err)

			str, _ := items.ToString()
			expected := typeString + arrayType +
				minItemsString + "79" +
				maxItemsString + "239" +
				uniqueItemsString + "false" +
				itemsString + strings.ReplaceAll(str, "\n", "\n  ")
			assert.Equal(t, expected, actual)
		})
	})
}

func TestArraySwaggerParameter_IsObject(t *testing.T) {
	t.Run("Test: ArraySwaggerParameter.IsObject()", func(t *testing.T) {
		t.Run("Should: return false", func(t *testing.T) {
			items := NewBoolSwaggerParameter(nil)
			a := NewArraySwaggerParameter(nil, items)
			actual := a.IsObject()
			assert.False(t, actual)
		})
	})
}

func TestArraySwaggerParameter_getConfigs(t *testing.T) {
	t.Run("Test: ArraySwaggerParameter.getConfigs()", func(t *testing.T) {
		t.Run("Should: return nil", func(t *testing.T) {
			items := NewBoolSwaggerParameter(nil)
			a := NewArraySwaggerParameter(nil, items)
			actual := a.getConfigs()
			assert.Nil(t, actual)
		})
	})
}
