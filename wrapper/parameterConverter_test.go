package wrapper

import (
	"github.com/mcmakler/swagger-gin-generator/swaggerFileGenerator/parameters"
	"github.com/stretchr/testify/assert"
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

func TestConvertObjectToSwaggerParameter(t *testing.T) {
	t.Run("Test: utils.ConvertObjectToSwaggerParameter", func(t *testing.T) {
		t.Run("Should: return SwaggerParameter", func(t *testing.T) {
			params := map[string]interface{}{
				"nameOfVariable": "TestStructEmpty",
				"required":       nil,
			}
			expected := parameters.NewObjectSwaggerParameter(params, make(map[string]parameters.SwaggParameter), false)
			actual := ConvertObjectToSwaggerParameter(nil, &TestStructEmpty{}, false)
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return SwaggerParameter", func(t *testing.T) {
			params := map[string]interface{}{
				"nameOfVariable": "TestStructBool",
				"required":       nil,
			}
			expected := parameters.NewObjectSwaggerParameter(params, map[string]parameters.SwaggParameter{
				"B": parameters.NewBoolSwaggerParameter(nil),
			}, false)
			actual := ConvertObjectToSwaggerParameter(nil, &TestStructBool{
				B: true,
			}, false)
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return SwaggerParameter", func(t *testing.T) {
			params := map[string]interface{}{
				"nameOfVariable": "TestStructArr",
				"required":       nil,
			}
			expected := parameters.NewObjectSwaggerParameter(params, map[string]parameters.SwaggParameter{
				"A": parameters.NewArraySwaggerParameter(nil, parameters.NewBoolSwaggerParameter(nil)),
			}, false)
			actual := ConvertObjectToSwaggerParameter(nil, &TestStructArr{
				A: []bool{true},
			}, false)
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return SwaggerParameter", func(t *testing.T) {
			params := map[string]interface{}{
				"nameOfVariable": "TestStructTime",
				"required":       nil,
			}
			expected := parameters.NewObjectSwaggerParameter(params, map[string]parameters.SwaggParameter{
				"T": parameters.NewStringSwaggerParameter(nil),
			}, false)
			actual := ConvertObjectToSwaggerParameter(nil, &TestStructTime{}, false)
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return SwaggerParameter", func(t *testing.T) {
			params := map[string]interface{}{
				"nameOfVariable": "TestStructSubstr",
				"required":       nil,
			}
			paramsBool := map[string]interface{}{
				"nameOfVariable": "TestStructBool",
				"required":       nil,
			}
			expected := parameters.NewObjectSwaggerParameter(params, map[string]parameters.SwaggParameter{
				"Substr": parameters.NewObjectSwaggerParameter(paramsBool, map[string]parameters.SwaggParameter{
					"B": parameters.NewBoolSwaggerParameter(nil),
				}, true),
			}, false)
			actual := ConvertObjectToSwaggerParameter(nil, &TestStructSubstr{
				Substr: TestStructBool{B: false},
			}, false)
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return SwaggerParameter", func(t *testing.T) {
			params := map[string]interface{}{
				"nameOfVariable": "TestStructFull",
				"required":       nil,
			}
			paramsBool := map[string]interface{}{
				"nameOfVariable": "TestStructBool",
				"required":       nil,
			}
			expected := parameters.NewObjectSwaggerParameter(params, map[string]parameters.SwaggParameter{
				"B": parameters.NewBoolSwaggerParameter(nil),
				"S": parameters.NewStringSwaggerParameter(nil),
				"I": parameters.NewIntegerSwaggerParameter(nil),
				"F": parameters.NewNumberSwaggerParameter(nil),
				"A": parameters.NewArraySwaggerParameter(nil, parameters.NewBoolSwaggerParameter(nil)),
				"Substr": parameters.NewObjectSwaggerParameter(paramsBool, map[string]parameters.SwaggParameter{
					"B": parameters.NewBoolSwaggerParameter(nil),
				}, true),
			}, false)
			actual := ConvertObjectToSwaggerParameter(nil, &TestStructFull{
				B:      false,
				S:      "",
				I:      0,
				F:      0,
				A:      []bool{true},
				Substr: &TestStructBool{B: false},
			}, false)
			assert.Equal(t, expected, actual)
		})
	})
}

func TestParameter_GetSwaggerParameter(t *testing.T) {
	t.Run("Test: Parameter.getSwaggerParameter()", func(t *testing.T) {
		t.Run("Should: return new bool swag parameter", func(t *testing.T) {
			p := NewParameter(nil, true)
			actual := p.getSwaggerParameter(true)
			expected := parameters.NewSchemaSwaggParameter(parameters.NewBoolSwaggerParameter(nil))
			assert.Equal(t, expected, actual)
		})
	})
}
