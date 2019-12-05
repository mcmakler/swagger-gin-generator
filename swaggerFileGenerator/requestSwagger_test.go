package swaggerFileGenerator

import (
	"errors"
	"github.com/mcmakler/swagger-gin-generator/swaggerFileGenerator/parameters"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestRequestSwagger_ToString(t *testing.T) {
	t.Run("Test: RequestSwagger.ToString()", func(t *testing.T) {
		t.Run("Should: return error "+errorEmptyTypeRequest.Error(), func(t *testing.T) {
			a := NewRequestSwagg(nil, nil, nil)
			_, actual := a.ToString()
			expected := errorEmptyTypeRequest
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return error "+errorEmptyTypeRequest.Error(), func(t *testing.T) {
			config := map[string]interface{}{
				"test": 1,
			}
			a := NewRequestSwagg(config, nil, nil)
			_, actual := a.ToString()
			expected := errorEmptyTypeRequest
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return error"+errorEmptyResponses.Error(), func(t *testing.T) {
			config := map[string]interface{}{
				"typeRequest": "GET",
			}
			a := NewRequestSwagg(config, nil, nil)
			_, actual := a.ToString()
			expected := errorEmptyResponses
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return error NIL_OBJECT_VARIABLE_NAME", func(t *testing.T) {
			config := map[string]interface{}{
				"typeRequest": "GET",
			}
			swaggerParameter := parameters.NewObjectSwaggerParameter(nil, nil, false)
			responseSwagger := NewResponseSwagg(200, "description", "", nil)
			a := NewRequestSwagg(config, []parameters.SwaggParameter{swaggerParameter}, []ResponseSwagg{responseSwagger})
			_, actual := a.ToString()
			expected := errors.New("NIL_OBJECT_VARIABLE_NAME")
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return error WRONG_HTTP_CODE", func(t *testing.T) {
			config := map[string]interface{}{
				"typeRequest": "GET",
			}
			responseSwagger := NewResponseSwagg(-1, "", "", nil)
			a := NewRequestSwagg(config, nil, []ResponseSwagg{responseSwagger})

			_, actual := a.ToString()
			expected := errors.New("WRONG_HTTP_CODE")
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
			responseSwagger := NewResponseSwagg(200, "description", "", nil)
			a := NewRequestSwagg(config, []parameters.SwaggParameter{swaggerParameter}, []ResponseSwagg{responseSwagger})
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
