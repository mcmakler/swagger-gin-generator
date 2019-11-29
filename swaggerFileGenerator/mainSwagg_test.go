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
				params:      nil,
				paths:       nil,
				definitions: nil,
			}
			_, error := a.ToString()
			assert.Equal(t, error, errorEmptyPaths)
		})

		t.Run("Should: return error "+errorIncorrectPath.Error(), func(t *testing.T) {
			a := &mainSwagg{
				params: nil,
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
			path := &pathSwagger{
				path: "path",
				requests: []RequestSwagg{
					NewRequestSwagg(map[string]interface{}{
						"typeRequest": "GET",
					}, nil, nil),
				},
			}
			a := &mainSwagg{
				params: nil,
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

		t.Run("Should: return ok", func(t *testing.T) {
			path := &pathSwagger{
				path: "path",
				requests: []RequestSwagg{
					NewRequestSwagg(map[string]interface{}{
						"typeRequest": "GET",
					}, nil, nil),
				},
			}
			a := &mainSwagg{
				params: nil,
				paths: []PathSwagger{
					path,
				},
				definitions: nil,
			}
			str, _ := path.ToString()
			expected := swaggerString +
				basePathString + "/" +
				pathsString + strings.Replace(str, "\n", mainIndentString, -1)
			actual, _ := a.ToString()
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return ok", func(t *testing.T) {
			path := &pathSwagger{
				path: "path",
				requests: []RequestSwagg{
					NewRequestSwagg(map[string]interface{}{
						"typeRequest": "GET",
					}, nil, nil),
				},
			}
			a := &mainSwagg{
				params: map[string]interface{}{
					"description": "description",
				},
				paths: []PathSwagger{
					path,
				},
				definitions: nil,
			}
			str, _ := path.ToString()
			expected := swaggerString + infoString +
				infoDescriptionString + "description" +
				basePathString + "/" +
				pathsString + strings.Replace(str, "\n", mainIndentString, -1)
			actual, _ := a.ToString()
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return ok", func(t *testing.T) {
			path := &pathSwagger{
				path: "path",
				requests: []RequestSwagg{
					NewRequestSwagg(map[string]interface{}{
						"typeRequest": "GET",
					}, nil, nil),
				},
			}
			a := &mainSwagg{
				params: map[string]interface{}{
					"title": "title",
				},
				paths: []PathSwagger{
					path,
				},
				definitions: nil,
			}
			str, _ := path.ToString()
			expected := swaggerString + infoString +
				infoTitleString + "title" +
				basePathString + "/" +
				pathsString + strings.Replace(str, "\n", mainIndentString, -1)
			actual, _ := a.ToString()
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return ok", func(t *testing.T) {
			path := &pathSwagger{
				path: "path",
				requests: []RequestSwagg{
					NewRequestSwagg(map[string]interface{}{
						"typeRequest": "GET",
					}, nil, nil),
				},
			}
			a := &mainSwagg{
				params: map[string]interface{}{
					"version": "version",
				},
				paths: []PathSwagger{
					path,
				},
				definitions: nil,
			}
			str, _ := path.ToString()
			expected := swaggerString + infoString +
				infoVersionString + "version" +
				basePathString + "/" +
				pathsString + strings.Replace(str, "\n", mainIndentString, -1)
			actual, _ := a.ToString()
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return ok", func(t *testing.T) {
			path := &pathSwagger{
				path: "path",
				requests: []RequestSwagg{
					NewRequestSwagg(map[string]interface{}{
						"typeRequest": "GET",
					}, nil, nil),
				},
			}
			a := &mainSwagg{
				params: map[string]interface{}{
					"basePath": "basePath",
				},
				paths: []PathSwagger{
					path,
				},
				definitions: nil,
			}
			str, _ := path.ToString()
			expected := swaggerString +
				basePathString + "basePath" +
				pathsString + strings.Replace(str, "\n", mainIndentString, -1)
			actual, _ := a.ToString()
			assert.Equal(t, expected, actual)
		})



		t.Run("Should: return ok", func(t *testing.T) {
			path := &pathSwagger{
				path: "path",
				requests: []RequestSwagg{
					NewRequestSwagg(map[string]interface{}{
						"typeRequest": "GET",
					}, nil, nil),
				},
			}
			def := parameters.NewBoolSwagParameter(nil)
			a := &mainSwagg{
				params: map[string]interface{}{
					"description": "description",
					"title": "title",
					"version": "version",
					"basePath": "basePath",
				},
				paths: []PathSwagger{
					path,
				},
				definitions: []parameters.SwaggParameter{
					def,
				},
			}
			strPath, _ := path.ToString()
			strDef, _ := def.ToString()
			expected := swaggerString + infoString +
				infoDescriptionString + "description" +
				infoTitleString + "title" +
				infoVersionString + "version" +
				basePathString + "basePath" +
				pathsString + strings.Replace(strPath, "\n", mainIndentString, -1) +
				definitionsString + strings.Replace(strDef, "\n", mainIndentString, -1)
			actual, _ := a.ToString()
			assert.Equal(t, expected, actual)
		})
	})
}
