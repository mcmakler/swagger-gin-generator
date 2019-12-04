package swaggerFileGenerator

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"strings"
	"swagger-gin-generator/wrapper/swaggerFileGenerator/parameters"
	"testing"
)

func TestRequestSwagg_ToString(t *testing.T) {
	t.Run("Test: RequestSwagg.ToString()", func(t *testing.T) {
		t.Run("Should: return error EMPTY_TYPE_OF_REQUEST", func(t *testing.T) {
			a := &requestSwagg{
				configs:    nil,
				parameters: nil,
				responses:  nil,
			}
			_, error := a.ToString()
			assert.Equal(t, error, errorEmptyTypeRequest)
		})

		t.Run("Should: return error EMPTY_TYPE_OF_REQUEST", func(t *testing.T) {
			swaggParams := map[string]interface{}{
				"test": 1,
			}
			a := &requestSwagg{
				configs:    swaggParams,
				parameters: nil,
				responses:  nil,
			}
			_, error := a.ToString()
			assert.Equal(t, error, errorEmptyTypeRequest)
		})

		t.Run("Should: return empty request", func(t *testing.T) {
			swaggParams := map[string]interface{}{
				"typeRequest": "GET",
			}
			a := &requestSwagg{
				configs:    swaggParams,
				parameters: nil,
				responses:  nil,
			}
			expected := errorEmptyResponses
			_, actual := a.ToString()
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return error NIL_OBJECT_NAME", func(t *testing.T) {
			swaggParams := map[string]interface{}{
				"typeRequest": "GET",
			}
			swagP1 := parameters.NewObjectSwaggerParameter(nil, nil, false)
			responseSwagg1 := NewResponseSwagg(200, "descr", "")
			a := &requestSwagg{
				configs:    swaggParams,
				parameters: []parameters.SwaggParameter{swagP1},
				responses:  []ResponseSwagg{responseSwagg1},
			}
			_, err := a.ToString()
			assert.Equal(t, errors.New("NIL_OBJECT_NAME"), err)
		})

		t.Run("Should: return error WRONG_HTTP_CODE", func(t *testing.T) {
			swaggParams := map[string]interface{}{
				"typeRequest": "GET",
			}
			responseSwagg1 := NewResponseSwagg(-1, "", "")
			a := &requestSwagg{
				configs:    swaggParams,
				parameters: nil,
				responses:  []ResponseSwagg{responseSwagg1},
			}
			_, err := a.ToString()
			assert.Equal(t, errors.New("WRONG_HTTP_CODE"), err)
		})

		t.Run("Should: return empty request", func(t *testing.T) {
			swaggParams := map[string]interface{}{
				"typeRequest": "GET",
				"description": "description",
				"consumes":    []string{"c1", "c2"},
				"produces":    []string{"p1", "p2"},
				"tags":        []string{"t1", "t2"},
				"summary":     "summary",
			}
			swagP1 := parameters.NewBoolSwagParameter(nil)
			swagP2 := parameters.NewStringSwagParameter(nil)
			responseSwagg1 := NewResponseSwagg(200, "descr", "")
			responseSwagg2 := NewResponseSwagg(300, "descr", "")
			a := &requestSwagg{
				configs:    swaggParams,
				parameters: []parameters.SwaggParameter{swagP1, swagP2},
				responses:  []ResponseSwagg{responseSwagg1, responseSwagg2},
			}
			expected := "\nGET:" +
				descriptionString + "description" +
				consumesString + consumesIndentString + "c1" + consumesIndentString + "c2" +
				producesString + producesIndentString + "p1" + producesIndentString + "p2" +
				tagsString + tagsIndentString + "t1" + tagsIndentString + "t2" +
				summaryString + "summary" +
				parametersString
			str, _ := swagP1.ToString(false)
			expected += strings.Replace(str, "\n", parametersIndentString, -1)
			str, _ = swagP2.ToString(false)
			expected += strings.Replace(str, "\n", parametersIndentString, -1) +
				responsesString
			str, _ = responseSwagg1.ToString()
			expected += strings.Replace(str, "\n", parametersIndentString, -1)
			str, _ = responseSwagg2.ToString()
			expected += strings.Replace(str, "\n", parametersIndentString, -1)
			actual, _ := a.ToString()
			assert.Equal(t, expected, actual)
		})
	})
}

func TestNewRequestSwagg(t *testing.T) {
	t.Run("Test: NewRequestSwagg", func(t *testing.T) {
		swaggParams := map[string]interface{}{
			"typeRequest": "GET",
			"description": "description",
			"consumes":    []string{"c1", "c2"},
			"produces":    []string{"p1", "p2"},
			"tags":        []string{"t1", "t2"},
			"summary":     "summary",
		}
		swagP1 := parameters.NewBoolSwagParameter(nil)
		swagP2 := parameters.NewStringSwagParameter(nil)
		responseSwagg1 := NewResponseSwagg(200, "descr", "")
		responseSwagg2 := NewResponseSwagg(300, "descr", "")
		expected := &requestSwagg{
			configs:    swaggParams,
			parameters: []parameters.SwaggParameter{swagP1, swagP2},
			responses:  []ResponseSwagg{responseSwagg1, responseSwagg2},
		}
		actual := NewRequestSwagg(
			swaggParams,
			[]parameters.SwaggParameter{swagP1, swagP2},
			[]ResponseSwagg{responseSwagg1, responseSwagg2},
		)
		assert.Equal(t, expected, actual)
	})
}
