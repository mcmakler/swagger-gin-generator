package parameters

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestObjectSwaggerParameter_ToString(t *testing.T) {
	t.Run("Test: ObjectSwaggerParameter.ToString()", func(t *testing.T) {
		t.Run("Should: return error "+errorNilObjectVariableName.Error(), func(t *testing.T) {
			a := NewObjectSwaggerParameter(nil, nil, false)
			_, actual := a.ToString()
			expected := errorNilObjectVariableName
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return error "+errorNilObjectVariableName.Error(), func(t *testing.T) {
			config := map[string]interface{}{}
			a := NewObjectSwaggerParameter(config, nil, false)
			_, actual := a.ToString()
			expected := errorNilObjectVariableName
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return error "+errorNilObjectVariableName.Error(), func(t *testing.T) {
			config := map[string]interface{}{
				"nameOfVariable": "name",
			}
			properties := map[string]SwaggParameter{
				"objectParam": NewObjectSwaggerParameter(nil, nil, false),
			}
			a := NewObjectSwaggerParameter(config, properties, false)
			_, actual := a.ToString()
			expected := errorNilObjectVariableName
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return with full config", func(t *testing.T) {
			config := map[string]interface{}{
				"nameOfVariable": "name",
				"required":       []string{"required"},
			}
			booleanParameter := NewBoolSwaggerParameter(nil)
			properties := map[string]SwaggParameter{
				"booleanParameter": booleanParameter,
			}
			a := NewObjectSwaggerParameter(config, properties, false)
			actual, err := a.ToString()
			assert.NoError(t, err)

			str, _ := booleanParameter.ToString()
			expected := "\nname:" + typeString + objectType +
				requiredString + requiredIndentStr + "required" +
				propertiesStr +
				propertyIndentStr + "booleanParameter:" +
				strings.Replace(str, "\n", propertyIndentStr, -1)
			assert.Equal(t, expected, actual)
		})
	})
}

func TestObjectSwaggerParameter_IsObject(t *testing.T) {
	t.Run("Test: ObjectSwaggerParameter.IsObject()", func(t *testing.T) {
		t.Run("Should: return true", func(t *testing.T) {
			config := map[string]interface{}{
				"nameOfVariable": "name",
			}
			a := NewObjectSwaggerParameter(config, nil, false)
			actual := a.IsObject()
			assert.True(t, actual)
		})
	})
}

func TestObjectSwaggerParameter_getConfigs(t *testing.T) {
	t.Run("Test: ObjectSwaggerParameter.getConfigs()", func(t *testing.T) {
		t.Run("Should: return config with nameOfVariable", func(t *testing.T) {
			expected := map[string]interface{}{
				"nameOfVariable": "name",
			}
			a := NewObjectSwaggerParameter(expected, nil, false)
			actual := a.getConfigs()
			assert.Equal(t, expected, actual)
		})
	})
}
