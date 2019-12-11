package swaggerFileGenerator

import (
	"github.com/mcmakler/swagger-gin-generator/swaggerFileGenerator/parameters"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestResponseSwagg_ToString(t *testing.T) {
	t.Run("Test: ResponseSwagger.ToString()", func(t *testing.T) {
		t.Run("Should: return error"+errorWrongCode.Error(), func(t *testing.T) {
			a := NewResponseSwagger(-1, "", "", nil)
			_, actual := a.ToString()
			expected := errorWrongCode
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return error"+errorEmptyDescription.Error(), func(t *testing.T) {
			a := NewResponseSwagger(200, "", "", nil)
			_, actual := a.ToString()
			expected := errorEmptyDescription
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return error"+parameters.ErrorNillItemsParameter.Error(), func(t *testing.T) {
			parameter := parameters.NewArraySwaggerParameter(nil, nil)
			a := NewResponseSwagger(200, "description", "", parameter)
			_, actual := a.ToString()
			expected := parameters.ErrorNillItemsParameter
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return responseSwagger with nill param", func(t *testing.T) {
			a := NewResponseSwagger(200, "description", "linkonschema", nil)
			actual, err := a.ToString()
			assert.NoError(t, err)

			expected := "\n'200':" +
				descriptionString + "description" +
				linkOnSchemaString + refString + "linkonschema'"
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return full responseSwagger", func(t *testing.T) {
			parameter := parameters.NewBoolSwaggerParameter(nil)
			a := NewResponseSwagger(200, "description", "", parameter)
			actual, err := a.ToString()
			assert.NoError(t, err)

			parameterString, _ := parameter.ToString()
			expected := "\n'200':" +
				descriptionString + "description" +
				linkOnSchemaString + strings.ReplaceAll(parameterString, "\n", parametersIndentString)
			assert.Equal(t, expected, actual)
		})
	})
}
