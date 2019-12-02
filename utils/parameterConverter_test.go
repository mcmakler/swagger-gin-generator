package utils

import (
	"github.com/stretchr/testify/assert"
	"swagger-gin-generator/swaggerFileGenerator/parameters"
	"testing"
)

type testStructEmpty struct {
}

type testStructBool struct {
	B bool
}

type testStructFull struct {
	b      bool
	s      string
	i      int
	f      float64
	a      []bool
	substr testStructBool
}

func TestConvertObjectToSwaggParameter(t *testing.T) {
	t.Run("Test: utils.ConvertObjectToSwaggParameter", func(t *testing.T) {
		t.Run("Should: return empty swag object", func(t *testing.T) {
			params := map[string]interface{}{
				"name": "testStructEmpty",
			}
			expected := parameters.NewObjectSwaggerParameter(params, make(map[string]parameters.SwaggParameter))
			assert.Equal(t, expected, ConvertObjectToSwaggParameter(nil, testStructEmpty{}))
		})

		t.Run("Should: return swag object with bool param", func(t *testing.T) {
			params := map[string]interface{}{
				"name": "testStructBool",
			}
			expected := parameters.NewObjectSwaggerParameter(params, map[string]parameters.SwaggParameter {
				"B": parameters.NewBoolSwagParameter(nil),
			})
			assert.Equal(t, expected, ConvertObjectToSwaggParameter(nil, testStructBool{
				B: true,
			}))
		})
	})
}

func TestParameter_GetSwagParameter(t *testing.T) {
	t.Run("Test: Parameter.GetSwagParameter()", func(t *testing.T) {
		t.Run("Should: return new bool swag parameter", func(t *testing.T) {
			p := &parameter{
				listOfparameters: nil,
				object:           true,
			}
			expected := parameters.NewBoolSwagParameter(nil)
			actual := p.GetSwagParameter()
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return new string swag parameter", func(t *testing.T) {
			p := &parameter{
				listOfparameters: nil,
				object:           "String",
			}
			expected := parameters.NewStringSwagParameter(nil)
			actual := p.GetSwagParameter()
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return new int swag parameter", func(t *testing.T) {
			p := &parameter{
				listOfparameters: nil,
				object:           10,
			}
			expected := parameters.NewIntegerSwagParameter(nil)
			actual := p.GetSwagParameter()
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return new int swag parameter", func(t *testing.T) {
			p := &parameter{
				listOfparameters: nil,
				object:           int64(10),
			}
			expected := parameters.NewIntegerSwagParameter(nil)
			actual := p.GetSwagParameter()
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return new number swag parameter", func(t *testing.T) {
			p := &parameter{
				listOfparameters: nil,
				object:           10.10,
			}
			expected := parameters.NewNumberSwagParameter(nil)
			actual := p.GetSwagParameter()
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return new number swag parameter", func(t *testing.T) {
			p := &parameter{
				listOfparameters: nil,
				object:           float64(10),
			}
			expected := parameters.NewNumberSwagParameter(nil)
			actual := p.GetSwagParameter()
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return new bool swag parameter", func(t *testing.T) {
			p := &parameter{
				listOfparameters: nil,
				object:           float64(10),
			}
			expected := parameters.NewNumberSwagParameter(nil)
			actual := p.GetSwagParameter()
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return new array of bool swag parameter", func(t *testing.T) {
			p := &parameter{
				listOfparameters: nil,
				object:           []bool{},
			}
			expected := parameters.NewArraySwaggParameter(nil, parameters.NewBoolSwagParameter(nil))
			actual := p.GetSwagParameter()
			assert.Equal(t, expected, actual)
		})
	})
}
