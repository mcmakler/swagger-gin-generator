package swaggerFileGenerator

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"swagger-gin-generator/swaggerFileGenerator/parameters"
	"testing"
)

func TestMainSwagg_ToString(t *testing.T) {
	t.Run("Test: MainSwagg.ToString()", func(t *testing.T) {
		t.Run("Should: return error "+errorEmptyPaths.Error(), func(t *testing.T) {
			a := &mainSwagg{
				configs:     nil,
				paths:       nil,
				definitions: nil,
			}
			_, error := a.ToString()
			assert.Equal(t, error, errorEmptyPaths)
		})

		t.Run("Should: return error "+errorIncorrectPath.Error(), func(t *testing.T) {
			a := &mainSwagg{
				configs: map[string]interface{}{
					"title":   "title",
					"version": "version",
				},
				paths: []PathSwagger{
					&pathSwagger{
						path:     "",
						requests: nil,
					},
				},
				definitions: nil,
			}
			_, error := a.ToString()
			assert.Equal(t, error, errorIncorrectPath)
		})

		t.Run("Should: return error "+parameters.ErrorNillItemsParameter.Error(), func(t *testing.T) {
			responseSwagg1 := NewResponseSwagg(200, "descr", "")
			path := &pathSwagger{
				path: "path",
				requests: []RequestSwagg{
					NewRequestSwagg(map[string]interface{}{
						"typeRequest": "GET",
					}, nil, []ResponseSwagg{responseSwagg1}),
				},
			}
			a := &mainSwagg{
				configs: map[string]interface{}{
					"title":   "title",
					"version": "version",
				},
				paths: []PathSwagger{
					path,
				},
				definitions: []parameters.SwaggParameter{
					parameters.NewArraySwaggParameter(nil, nil),
				},
			}
			_, error := a.ToString()
			assert.Equal(t, error, parameters.ErrorNillItemsParameter)
		})

		t.Run("Should: return empty title", func(t *testing.T) {
			responseSwagg1 := NewResponseSwagg(200, "descr", "")
			path := &pathSwagger{
				path: "path",
				requests: []RequestSwagg{
					NewRequestSwagg(map[string]interface{}{
						"typeRequest": "GET",
					}, nil, []ResponseSwagg{responseSwagg1}),
				},
			}
			a := &mainSwagg{
				configs: map[string]interface{}{
					"version": "version",
				},
				paths: []PathSwagger{
					path,
				},
				definitions: nil,
			}
			expected := errorEmptyTitle
			_, actual := a.ToString()
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return empty version", func(t *testing.T) {
			responseSwagg1 := NewResponseSwagg(200, "descr", "")
			path := &pathSwagger{
				path: "path",
				requests: []RequestSwagg{
					NewRequestSwagg(map[string]interface{}{
						"typeRequest": "GET",
					}, nil, []ResponseSwagg{responseSwagg1}),
				},
			}
			a := &mainSwagg{
				configs: map[string]interface{}{
					"title": "title",
				},
				paths: []PathSwagger{
					path,
				},
				definitions: nil,
			}
			expected := errorEmptyVersion
			_, actual := a.ToString()
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return empty version", func(t *testing.T) {
			responseSwagg1 := NewResponseSwagg(200, "descr", "")
			path := &pathSwagger{
				path: "path",
				requests: []RequestSwagg{
					NewRequestSwagg(map[string]interface{}{
						"typeRequest": "GET",
					}, nil, []ResponseSwagg{responseSwagg1}),
				},
			}
			a := NewMainSwagg(
				map[string]interface{}{
					"title":   "title",
					"version": "version",
				},
				[]SecurityDefinitionSwagg{
					NewOauth2AccessCodeSecurityDefinition("", "", ""),
				},
				[]PathSwagger{
					path,
				},
				nil,
			)
			expected := errorEmptySecurityTitle
			_, actual := a.ToString()
			assert.Equal(t, expected, actual)
		})



		t.Run("Should: return empty version", func(t *testing.T) {
			basicSD := NewBasicSecurityDefinition("title")
			responseSwagg1 := NewResponseSwagg(200, "descr", "")
			path := &pathSwagger{
				path: "path",
				requests: []RequestSwagg{
					NewRequestSwagg(map[string]interface{}{
						"typeRequest": "GET",
					}, nil, []ResponseSwagg{responseSwagg1}),
				},
			}
			a := NewMainSwagg(
				map[string]interface{}{
					"title":   "title",
					"version": "version",
				},
				[]SecurityDefinitionSwagg{
					basicSD,
				},
				[]PathSwagger{
					path,
				},
				nil,
			)
			strSecurity, _ := basicSD.ToString()
			strPath, _ := path.ToString()
			expected := swaggerString + infoString +
				infoTitleString + "title" +
				infoVersionString + "version" +
				securityDefinitionString + strings.Replace(strSecurity, "\n", mainIndentString, -1) +
				basePathString + "/" +
				pathsString + strings.Replace(strPath, "\n", mainIndentString, -1)
			actual, error := a.ToString()
			assert.NoError(t, error)
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return ok", func(t *testing.T) {
			responseSwagg1 := NewResponseSwagg(200, "descr", "")
			path := &pathSwagger{
				path: "path",
				requests: []RequestSwagg{
					NewRequestSwagg(map[string]interface{}{
						"typeRequest": "GET",
					}, nil, []ResponseSwagg{responseSwagg1}),
				},
			}
			a := &mainSwagg{
				configs: map[string]interface{}{
					"title":   "title",
					"version": "version",
				},
				paths: []PathSwagger{
					path,
				},
				definitions: nil,
			}
			str, _ := path.ToString()
			expected := swaggerString + infoString +
				infoTitleString + "title" +
				infoVersionString + "version" +
				basePathString + "/" +
				pathsString + strings.Replace(str, "\n", mainIndentString, -1)
			actual, _ := a.ToString()
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return ok", func(t *testing.T) {
			responseSwagg1 := NewResponseSwagg(200, "descr", "")
			path := &pathSwagger{
				path: "path",
				requests: []RequestSwagg{
					NewRequestSwagg(map[string]interface{}{
						"typeRequest": "GET",
					}, nil, []ResponseSwagg{responseSwagg1}),
				},
			}
			a := &mainSwagg{
				configs: map[string]interface{}{
					"title":    "title",
					"version":  "version",
					"basePath": "basePath",
				},
				paths: []PathSwagger{
					path,
				},
				definitions: nil,
			}
			str, _ := path.ToString()
			expected := swaggerString + infoString +
				infoTitleString + "title" +
				infoVersionString + "version" +
				basePathString + "basePath" +
				pathsString + strings.Replace(str, "\n", mainIndentString, -1)
			actual, _ := a.ToString()
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return ok", func(t *testing.T) {
			responseSwagg1 := NewResponseSwagg(200, "descr", "")
			path := &pathSwagger{
				path: "path",
				requests: []RequestSwagg{
					NewRequestSwagg(map[string]interface{}{
						"typeRequest": "GET",
					}, nil, []ResponseSwagg{responseSwagg1}),
				},
			}
			def := parameters.NewBoolSwagParameter(nil)
			a := &mainSwagg{
				configs: map[string]interface{}{
					"description": "description",
					"title":       "title",
					"version":     "version",
					"basePath":    "basePath",
				},
				paths: []PathSwagger{
					path,
				},
				definitions: []parameters.SwaggParameter{
					def,
				},
			}
			strPath, _ := path.ToString()
			strDef, _ := def.ToString(true)
			expected := swaggerString + infoString +
				infoTitleString + "title" +
				infoVersionString + "version" +
				infoDescriptionString + "description" +
				basePathString + "basePath" +
				pathsString + strings.Replace(strPath, "\n", mainIndentString, -1) +
				definitionsString + strings.Replace(strDef, "\n", mainIndentString, -1)
			actual, _ := a.ToString()
			assert.Equal(t, expected, actual)
		})
	})
}

func TestNewMainSwagg(t *testing.T) {
	t.Run("Test: NewMainSwag", func(t *testing.T) {
		path := &pathSwagger{
			path: "path",
			requests: []RequestSwagg{
				NewRequestSwagg(map[string]interface{}{
					"typeRequest": "GET",
				}, nil, nil),
			},
		}
		expected := &mainSwagg{
			configs: nil,
			paths: []PathSwagger{
				path,
			},
			definitions: nil,
		}
		actual := NewMainSwagg(nil, nil, []PathSwagger{
			path,
		}, nil)
		assert.Equal(t, expected, actual)
	})
}
