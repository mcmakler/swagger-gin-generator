package parameters

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestObjectSwaggerParameter_ToString(t *testing.T) {
	t.Run("Test: ObjectSwaggParameter.ToString()", func(t *testing.T) {
		t.Run("Should: return error empty Object name", func(t *testing.T) {
			a := &objectSwaggerParameter{
				params:     nil,
				properties: nil,
			}
			_, expectedError := a.ToString()
			assert.Equal(t, expectedError, errorNilObjectName)
		})

		t.Run("Should: return string with empty params", func(t *testing.T) {
			params := map[string]interface{}{
				"test": 1,
			}
			a := &objectSwaggerParameter{
				params:     params,
				properties: nil,
			}
			_, expectedError := a.ToString()
			assert.Equal(t, expectedError, errorNilObjectName)
		})

		t.Run("Should: return string with empty params", func(t *testing.T) {
			params := map[string]interface{}{
				"name": "name",
			}
			properties := map[string]SwaggParameter{
				"objectParam": &objectSwaggerParameter{
					params:     nil,
					properties: nil,
				},
			}
			a := &objectSwaggerParameter{
				params:     params,
				properties: properties,
			}
			_, expectedError := a.ToString()
			assert.Equal(t, expectedError, errorNilObjectName)
		})

		t.Run("Should: return string with empty params", func(t *testing.T) {
			params := map[string]interface{}{
				"name": "name",
			}
			a := &objectSwaggerParameter{
				params:     params,
				properties: nil,
			}
			expected := "\nname:" + typeString + objectType
			actual, _ := a.ToString()
			assert.Equal(t, expected, actual)
		})
		t.Run("Should: return string with all params", func(t *testing.T) {
			params := map[string]interface{}{
				"name": "name",
				"required": []string{"req1", "req2"},
			}
			a := &objectSwaggerParameter{
				params:     params,
				properties: nil,
			}
			expected := "\nname:" + typeString + objectType +
				requiredIndentStr + "req1" + requiredIndentStr + "req2"
			actual, _ := a.ToString()
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return string with all params", func(t *testing.T) {
			params := map[string]interface{}{
				"name": "name",
				"required": []string{"req1", "req2"},
			}
			properties := map[string]SwaggParameter{
				"boolParam":    &boolSwaggParameter{params: nil},
				"stringParam":  &stringSwaggParameter{params: nil},
				"integerParam": &integerSwaggParameter{params: nil},
				"numberParam":  &numberSwaggParameter{params: nil},
				"arrayParam": &arraySwaggParameter{
					params: nil,
					items:  &boolSwaggParameter{params: nil},
				},
				"objectParam": &objectSwaggerParameter{
					params:     params,
					properties: nil,
				},
			}
			a := &objectSwaggerParameter{
				params:     params,
				properties: properties,
			}
			expected := "\nname:" + typeString + objectType +
				requiredIndentStr + "req1" + requiredIndentStr + "req2" +
				propertiesStr
			str, _ := properties["boolParam"].ToString()
			boolStrExpexted := propertyIndentStr + "boolParam:" + strings.Replace(str, "\n", propertyIndentStr, -1)
			str, _ = properties["stringParam"].ToString()
			stringStrExpected := propertyIndentStr + "stringParam:" + strings.Replace(str, "\n", propertyIndentStr, -1)
			str, _ = properties["integerParam"].ToString()
			integerStrExpected := propertyIndentStr + "integerParam:" + strings.Replace(str, "\n", propertyIndentStr, -1)
			str, _ = properties["numberParam"].ToString()
			numberStrExpected := propertyIndentStr + "numberParam:" + strings.Replace(str, "\n", propertyIndentStr, -1)
			str, _ = properties["arrayParam"].ToString()
			arrayStrExpected := propertyIndentStr + "arrayParam:" + strings.Replace(str, "\n", propertyIndentStr, -1)
			str, _ = properties["objectParam"].ToString()
			objectStrExpected := propertyIndentStr + "objectParam:" + strings.Replace(str, "\n", propertyIndentStr, -1)
			actual, _ := a.ToString()
			assert.True(t, strings.Contains(actual, expected))
			assert.True(t, strings.Contains(actual, boolStrExpexted))
			assert.True(t, strings.Contains(actual, stringStrExpected))
			assert.True(t, strings.Contains(actual, integerStrExpected))
			assert.True(t, strings.Contains(actual, numberStrExpected))
			assert.True(t, strings.Contains(actual, arrayStrExpected))
			assert.True(t, strings.Contains(actual, objectStrExpected))
		})
	})
}
