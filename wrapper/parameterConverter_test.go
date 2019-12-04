package wrapper

import (
	"github.com/stretchr/testify/assert"
	"github.com/mcmakler/swagger-gin-generator/wrapper/swaggerFileGenerator/parameters"
	"testing"
	"time"
)

type TestStructEmpty struct {
}

type TestStructTime struct {
	T time.Time
}

type TestStructBool struct {
	B bool
}

type TestStructArr struct {
	A []bool
}

type TestStructSubstr struct {
	Substr TestStructBool
}

type TestStructFull struct {
	B      bool
	S      string
	I      int
	F      float64
	A      []bool
	Substr *TestStructBool
}

func TestConvertObjectToSwaggParameter(t *testing.T) {
	t.Run("Test: utils.ConvertObjectToSwaggParameter", func(t *testing.T) {
		t.Run("Should: return empty swag object", func(t *testing.T) {
			params := map[string]interface{}{
				"name": "TestStructEmpty",
			}
			expected := parameters.NewObjectSwaggerParameter(params, make(map[string]parameters.SwaggParameter), false)
			assert.Equal(t, expected, ConvertObjectToSwaggParameter(nil, &TestStructEmpty{}, false))
		})

		t.Run("Should: return swag object with bool param", func(t *testing.T) {
			params := map[string]interface{}{
				"name": "TestStructBool",
			}
			expected := parameters.NewObjectSwaggerParameter(params, map[string]parameters.SwaggParameter{
				"B": parameters.NewBoolSwagParameter(nil),
			}, false)
			assert.Equal(t, expected, ConvertObjectToSwaggParameter(nil, &TestStructBool{
				B: true,
			}, false))
		})

		t.Run("Should: return swag object with arr param", func(t *testing.T) {
			params := map[string]interface{}{
				"name": "TestStructArr",
			}
			expected := parameters.NewObjectSwaggerParameter(params, map[string]parameters.SwaggParameter{
				"A": parameters.NewArraySwaggParameter(nil, parameters.NewBoolSwagParameter(nil)),
			}, false)
			assert.Equal(t, expected, ConvertObjectToSwaggParameter(nil, &TestStructArr{
				A: []bool{true},
			}, false))
		})

		t.Run("Should: return time object", func(t *testing.T) {
			params := map[string]interface{}{
				"name": "TestStructTime",
			}
			expected := parameters.NewObjectSwaggerParameter(params, map[string]parameters.SwaggParameter{
				"T": parameters.NewStringSwagParameter(nil),
			}, false)

			assert.Equal(t, expected, ConvertObjectToSwaggParameter(nil, &TestStructTime{}, false))
		})

		t.Run("Should: return swag object with arr param", func(t *testing.T) {
			params := map[string]interface{}{
				"name": "TestStructSubstr",
			}
			paramsBool := map[string]interface{}{
				"name": "TestStructBool",
			}
			expected := parameters.NewObjectSwaggerParameter(params, map[string]parameters.SwaggParameter{
				"Substr": parameters.NewObjectSwaggerParameter(paramsBool, map[string]parameters.SwaggParameter{
					"B": parameters.NewBoolSwagParameter(nil),
				}, true),
			}, false)
			assert.Equal(t, expected, ConvertObjectToSwaggParameter(nil, &TestStructSubstr{
				Substr: TestStructBool{B: false},
			}, false))
		})

		t.Run("Should: return swag object with bool param", func(t *testing.T) {
			params := map[string]interface{}{
				"name": "TestStructFull",
			}
			paramsBool := map[string]interface{}{
				"name": "TestStructBool",
			}
			expected := parameters.NewObjectSwaggerParameter(params, map[string]parameters.SwaggParameter{
				"B": parameters.NewBoolSwagParameter(nil),
				"S": parameters.NewStringSwagParameter(nil),
				"I": parameters.NewIntegerSwagParameter(nil),
				"F": parameters.NewNumberSwagParameter(nil),
				"A": parameters.NewArraySwaggParameter(nil, parameters.NewBoolSwagParameter(nil)),
				"Substr": parameters.NewObjectSwaggerParameter(paramsBool, map[string]parameters.SwaggParameter{
					"B": parameters.NewBoolSwagParameter(nil),
				}, true),
			}, false)
			assert.Equal(t, expected, ConvertObjectToSwaggParameter(nil, &TestStructFull{
				B:      false,
				S:      "",
				I:      0,
				F:      0,
				A:      []bool{true},
				Substr: &TestStructBool{B: false},
			}, false))
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

func TestNewParameter(t *testing.T) {
	t.Run("Test: NewParameter", func(t *testing.T) {
		expected := &parameter{
			listOfparameters: nil,
			object:           nil,
		}
		actual := NewParameter(nil, nil)
		assert.Equal(t, expected, actual)
	})
}
