package parameters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegerSwaggParameter_ToString(t *testing.T) {
	t.Run("Test: IntegerSwaggParameter.ToString()", func(t *testing.T) {
		t.Run("Should: return string with empty params", func(t *testing.T) {
			a := &integerSwaggParameter{
				params: nil,
			}
			expected := typeString + integerType
			actual, _ := a.ToString()
			assert.Equal(t, expected, actual)
		})
		t.Run("Should: return string with empty params", func(t *testing.T) {
			params := map[string]interface{}{
				"test": 1,
			}
			a := &integerSwaggParameter{
				params: params,
			}
			expected := typeString + integerType
			actual, _ := a.ToString()
			assert.Equal(t, expected, actual)
		})
		t.Run("Should: return string with all params", func(t *testing.T) {
			params := map[string]interface{}{
				"in": "in",
				"name": "name",
				"required": true,
				"defaultValue": 10,
				"minimumValue": -2000000000,
				"exclusiveMinimumValue": true,
				"maximumValue": 2000000000,
				"exclusiveMaximumValue": false,
				"multipleOf": 4,
				"allowEmptyValue": false,
				"description": "description",
			}
			a := &integerSwaggParameter{
				params: params,
			}
			expected := typeString + integerType +
				inString + "in" +
				nameString + "name" +
				requiredString + "true" +
				defaultValueString + "10"  +
				minimumValueString + "-2000000000"  +
				exclusiveMinimumValueString + "true" +
				maximumValueString + "2000000000" +
				exclusiveMaximumValueString + "false" +
				multipleOfString + "4" +
				allowEmptyValueString + "false" +
				descriptionString + "description"
			actual, _ := a.ToString()
			assert.Equal(t, expected, actual)
		})
	})
}