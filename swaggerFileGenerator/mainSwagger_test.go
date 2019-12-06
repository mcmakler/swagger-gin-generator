package swaggerFileGenerator

import (
	"github.com/mcmakler/swagger-gin-generator/swaggerFileGenerator/parameters"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestMainSwagger_ToString(t *testing.T) {
	t.Run("Test: MainSwagger.ToString()", func(t *testing.T) {
		responseSwagger := NewResponseSwagger(200, "descr", "", nil)
		requestSwagger := NewRequestSwagger(
			map[string]interface{}{"typeRequest": "GET"},
			nil,
			[]ResponseSwagger{responseSwagger},
		)
		pathSwagger := NewPathSwagger(
			"path",
			[]RequestSwagger{requestSwagger},
		)

		t.Run("Should: return error "+errorEmptyPaths.Error(), func(t *testing.T) {
			a := NewMainSwagger(nil, nil, nil, nil)
			expected := errorEmptyPaths
			_, actual := a.ToString()
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return error "+errorIncorrectPath.Error(), func(t *testing.T) {
			a := NewMainSwagger(
				map[string]interface{}{
					"title":   "title",
					"version": "version",
				},
				nil,
				[]PathSwagger{NewPathSwagger("", nil)},
				nil,
			)
			expected := errorIncorrectPath
			_, actual := a.ToString()
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return error "+parameters.ErrorNillItemsParameter.Error(), func(t *testing.T) {
			a := NewMainSwagger(
				map[string]interface{}{
					"title":   "title",
					"version": "version",
				},
				nil,
				[]PathSwagger{pathSwagger},
				[]parameters.SwaggParameter{
					parameters.NewArraySwaggerParameter(nil, nil),
				},
			)
			_, actual := a.ToString()
			expected := parameters.ErrorNillItemsParameter
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return empty title", func(t *testing.T) {
			a := NewMainSwagger(
				map[string]interface{}{"version": "version"},
				nil,
				[]PathSwagger{pathSwagger},
				nil,
			)
			_, actual := a.ToString()
			expected := errorEmptyTitle
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return empty version", func(t *testing.T) {
			a := NewMainSwagger(
				map[string]interface{}{"title": "title"},
				nil,
				[]PathSwagger{pathSwagger},
				nil,
			)
			_, actual := a.ToString()
			expected := errorEmptyVersion
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return empty version", func(t *testing.T) {
			a := NewMainSwagger(
				map[string]interface{}{
					"title":   "title",
					"version": "version",
				},
				[]SecurityDefinitionSwagger{
					NewOauth2AccessCodeSecurityDefinition("", "", ""),
				},
				[]PathSwagger{pathSwagger},
				nil,
			)
			_, actual := a.ToString()
			expected := errorEmptySecurityTitle
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return empty version", func(t *testing.T) {
			basicSD := NewBasicSecurityDefinition("title")
			a := NewMainSwagger(
				map[string]interface{}{
					"title":   "title",
					"version": "version",
				},
				[]SecurityDefinitionSwagger{basicSD},
				[]PathSwagger{pathSwagger},
				nil,
			)
			actual, error := a.ToString()
			assert.NoError(t, error)

			strSecurity, _ := basicSD.ToString()
			strPath, _ := pathSwagger.ToString()
			expected := swaggerString + infoString +
				infoTitleString + "title" +
				infoVersionString + "version'" +
				securityDefinitionString + strings.Replace(strSecurity, "\n", mainIndentString, -1) +
				basePathString + "/" +
				pathsString + strings.Replace(strPath, "\n", mainIndentString, -1)
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return ok", func(t *testing.T) {
			a := NewMainSwagger(
				map[string]interface{}{
					"title":   "title",
					"version": "version",
				},
				nil,
				[]PathSwagger{pathSwagger},
				nil,
			)
			actual, error := a.ToString()
			assert.NoError(t, error)

			str, _ := pathSwagger.ToString()
			expected := swaggerString + infoString +
				infoTitleString + "title" +
				infoVersionString + "version'" +
				basePathString + "/" +
				pathsString + strings.Replace(str, "\n", mainIndentString, -1)
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return ok", func(t *testing.T) {
			a := NewMainSwagger(
				map[string]interface{}{
					"title":    "title",
					"version":  "version",
					"basePath": "basePath",
				},
				nil,
				[]PathSwagger{pathSwagger},
				nil,
			)
			actual, error := a.ToString()
			assert.NoError(t, error)

			str, _ := pathSwagger.ToString()
			expected := swaggerString + infoString +
				infoTitleString + "title" +
				infoVersionString + "version'" +
				basePathString + "basePath" +
				pathsString + strings.Replace(str, "\n", mainIndentString, -1)
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return ok", func(t *testing.T) {
			def := parameters.NewBoolSwaggerParameter(nil)
			a := NewMainSwagger(
				map[string]interface{}{
					"description": "description",
					"title":       "title",
					"version":     "version",
					"basePath":    "basePath",
				},
				nil,
				[]PathSwagger{pathSwagger},
				[]parameters.SwaggParameter{def},
			)
			actual, error := a.ToString()
			assert.NoError(t, error)

			strPath, _ := pathSwagger.ToString()
			strDef, _ := def.ToString()
			expected := swaggerString + infoString +
				infoTitleString + "title" +
				infoVersionString + "version'" +
				infoDescriptionString + "description" +
				basePathString + "basePath" +
				pathsString + strings.Replace(strPath, "\n", mainIndentString, -1) +
				definitionsString + strings.Replace(strDef, "\n", mainIndentString, -1)
			assert.Equal(t, expected, actual)
		})
	})
}
