package swaggerFileGenerator

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestPathSwagger_ToString(t *testing.T) {
	t.Run("Test: PathSwagger.ToString()", func(t *testing.T) {
		t.Run("Should: return error "+errorIncorrectPath.Error(), func(t *testing.T) {
			a := &pathSwagger{
				path:     "",
				requests: nil,
			}
			_, error := a.ToString()
			assert.Equal(t, error, errorIncorrectPath)
		})

		t.Run("Should: return error "+errorNullRequests.Error(), func(t *testing.T) {
			a := &pathSwagger{
				path:     "path",
				requests: nil,
			}
			_, error := a.ToString()
			assert.Equal(t, error, errorNullRequests)
		})

		t.Run("Should: return error "+errorEmptyTypeRequest.Error(), func(t *testing.T) {
			a := &pathSwagger{
				path: "path",
				requests: []RequestSwagg{
					NewRequestSwagg(nil, nil, nil),
				},
			}
			_, error := a.ToString()
			assert.Equal(t, error, errorEmptyTypeRequest)
		})

		t.Run("Should: pass", func(t *testing.T) {
			swaggParams := map[string]interface{}{
				"typeRequest": "GET",
			}
			req := []RequestSwagg{
				NewRequestSwagg(swaggParams, nil, nil),
			}
			a := &pathSwagger{
				path:     "path",
				requests: req,
			}
			str, _ := req[0].ToString()
			expected := "\npath:" + strings.Replace(str, "\n", requestsIndentString, -1)
			actual, _ := a.ToString()
			assert.Equal(t, expected, actual)
		})
	})
}

func TestNewPathSwagger(t *testing.T) {
	t.Run("Should: return new path swagger", func(t *testing.T) {
		swaggParams := map[string]interface{}{
			"typeRequest": "GET",
		}
		req := []RequestSwagg{
			NewRequestSwagg(swaggParams, nil, nil),
		}
		expected := &pathSwagger{
			path:     "path",
			requests: req,
		}
		actual := NewPathSwagger("path", req)
		assert.Equal(t, expected, actual)
	})
}
