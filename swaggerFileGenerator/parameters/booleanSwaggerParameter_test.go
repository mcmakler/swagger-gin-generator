package parameters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoolSwaggerParameter_ToString(t *testing.T) {
	t.Run("Test: BoolSwaggerParameter.ToString()", func(t *testing.T) {
		t.Run("Should: return with empty config", func(t *testing.T) {
			a := NewBoolSwaggerParameter(nil)
			actual, err := a.ToString()
			assert.NoError(t, err)

			expected := typeString + boolType
			assert.Equal(t, expected, actual)
		})
	})
}

func TestBoolSwaggerParameter_IsObject(t *testing.T) {
	t.Run("Test: BoolSwaggerParameter.IsObject()", func(t *testing.T) {
		t.Run("Should: return false", func(t *testing.T) {
			a := NewBoolSwaggerParameter(nil)
			actual := a.IsObject()
			assert.False(t, actual)
		})
	})
}
func TestBoolSwaggerParameter_getConfigs(t *testing.T) {
	t.Run("Test: BoolSwaggerParameter.getConfigs()", func(t *testing.T) {
		t.Run("Should: return nil", func(t *testing.T) {
			a := NewBoolSwaggerParameter(nil)
			actual := a.getConfigs()
			assert.Nil(t, actual)
		})
	})
}
