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
			_, expectedError := a.ToString(false)
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
			_, expectedError := a.ToString(false)
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
			_, expectedError := a.ToString(false)
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
			expected := "\nname:" + typeDeficeString + objectType
			actual, _ := a.ToString(false)
			assert.Equal(t, expected, actual)
		})
		t.Run("Should: return string with all params", func(t *testing.T) {
			params := map[string]interface{}{
				"name":     "name",
				"required": []string{"req1", "req2"},
			}
			a := &objectSwaggerParameter{
				params:     params,
				properties: nil,
			}
			expected := "\nname:" + typeString + objectType +
				requiredIndentStr + "req1" + requiredIndentStr + "req2"
			actual, _ := a.ToString(true)
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return string with all params", func(t *testing.T) {
			params := map[string]interface{}{
				"name": "name",
				"in":   "in",
			}
			objParams := map[string]interface{}{
				"name":     "name",
				"in":       "in",
				"required": []string{"req1", "req2"},
			}
			properties := map[string]SwaggParameter{
				"boolParam":    &boolSwaggParameter{params: params},
				"stringParam":  &stringSwaggParameter{params: params},
				"integerParam": &integerSwaggParameter{params: params},
				"numberParam":  &numberSwaggParameter{params: params},
				"arrayParam": &arraySwaggParameter{
					params: params,
					items:  &boolSwaggParameter{params: nil},
				},
				"objectParam": &objectSwaggerParameter{
					params:     objParams,
					properties: nil,
				},
			}
			a := &objectSwaggerParameter{
				params:     objParams,
				properties: properties,
			}
			expected := "\nname:" + typeDeficeString + objectType +
				requiredIndentStr + "req1" + requiredIndentStr + "req2" +
				propertiesStr
			str, _ := properties["boolParam"].ToString(false)
			boolStrExpexted := propertyIndentStr + "boolParam:" + strings.Replace(str, "\n", propertyIndentStr, -1)
			str, _ = properties["stringParam"].ToString(false)
			stringStrExpected := propertyIndentStr + "stringParam:" + strings.Replace(str, "\n", propertyIndentStr, -1)
			str, _ = properties["integerParam"].ToString(false)
			integerStrExpected := propertyIndentStr + "integerParam:" + strings.Replace(str, "\n", propertyIndentStr, -1)
			str, _ = properties["numberParam"].ToString(false)
			numberStrExpected := propertyIndentStr + "numberParam:" + strings.Replace(str, "\n", propertyIndentStr, -1)
			str, _ = properties["arrayParam"].ToString(false)
			arrayStrExpected := propertyIndentStr + "arrayParam:" + strings.Replace(str, "\n", propertyIndentStr, -1)
			str, _ = properties["objectParam"].ToString(false)
			objectStrExpected := propertyIndentStr + "objectParam:" + strings.Replace(str, "\n", propertyIndentStr, -1)
			actual, _ := a.ToString(false)
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

func TestNewObjectSwagParameter(t *testing.T) {
	t.Run("Test: NewStringSwagParameter", func(t *testing.T) {
		params := map[string]interface{}{
			"name": "name",
		}
		expect := &objectSwaggerParameter{
			params:     params,
			properties: nil,
		}
		actual := NewObjectSwaggerParameter(params, nil, false)
		assert.Equal(t, expect, actual)
	})
}
