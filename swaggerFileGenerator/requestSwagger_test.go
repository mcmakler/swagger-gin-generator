package swaggerFileGenerator

import (
	"github.com/mcmakler/swagger-gin-generator/swaggerFileGenerator/parameters"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestRequestSwagger_ToString(t *testing.T) {
	t.Run("Test: RequestSwagger.ToString()", func(t *testing.T) {
		t.Run("Should: return error "+errorEmptyTypeRequest.Error(), func(t *testing.T) {
			a := NewRequestSwagger(nil, nil, nil)
			_, actual := a.ToString()
			expected := errorEmptyTypeRequest
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return error "+errorEmptyTypeRequest.Error(), func(t *testing.T) {
			config := map[string]interface{}{
				"test": 1,
			}
			a := NewRequestSwagger(config, nil, nil)
			_, actual := a.ToString()
			expected := errorEmptyTypeRequest
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return error"+errorEmptyResponses.Error(), func(t *testing.T) {
			config := map[string]interface{}{
				"typeRequest": "GET",
			}
			a := NewRequestSwagger(config, nil, nil)
			_, actual := a.ToString()
			expected := errorEmptyResponses
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return error "+parameters.ErrorNillItemsParameter.Error(), func(t *testing.T) {
			config := map[string]interface{}{
				"typeRequest": "GET",
			}
			swaggerParameter := parameters.NewArraySwaggerParameter(nil, nil)
			responseSwagger := NewResponseSwagger(200, "description", "", nil)
			a := NewRequestSwagger(config, []parameters.SwaggParameter{swaggerParameter}, []ResponseSwagger{responseSwagger})
			_, actual := a.ToString()
			expected := parameters.ErrorNillItemsParameter
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return error "+errorWrongCode.Error(), func(t *testing.T) {
			config := map[string]interface{}{
				"typeRequest": "GET",
			}
			responseSwagger := NewResponseSwagger(-1, "", "", nil)
			a := NewRequestSwagger(config, nil, []ResponseSwagger{responseSwagger})

			_, actual := a.ToString()
			expected := errorWrongCode
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return request with full config", func(t *testing.T) {
			config := map[string]interface{}{
				"typeRequest": "GET",
				"security":    []string{"security"},
				"description": "description",
				"consumes":    []string{"consumes"},
				"produces":    []string{"produces"},
				"tags":        []string{"tags"},
				"operationId": "operationId",
				"summary":     "summary",
			}
			swaggerParameter := parameters.NewBoolSwaggerParameter(nil)
			responseSwagger := NewResponseSwagger(200, "description", "", nil)
			a := NewRequestSwagger(config, []parameters.SwaggParameter{swaggerParameter}, []ResponseSwagger{responseSwagger})
			actual, err := a.ToString()
			assert.NoError(t, err)

			strSwaggerParameter, _ := swaggerParameter.ToString()
			strResponseSwagger, _ := responseSwagger.ToString()
			expected := "\nGET:" +
				securityString + securityIndentString + "security: []" +
				descriptionString + "description" +
				consumesString + consumesIndentString + "consumes" +
				producesString + producesIndentString + "produces" +
				tagsString + tagsIndentString + "tags" +
				operationIdString + "operationId" +
				summaryString + "summary"
			expected += parametersString + strings.Replace(strSwaggerParameter, "\n", parametersIndentString, -1)
			expected += responsesString + strings.Replace(strResponseSwagger, "\n", parametersIndentString, -1)
			assert.Equal(t, expected, actual)
		})
	})
}
