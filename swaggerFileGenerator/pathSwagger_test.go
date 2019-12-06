package swaggerFileGenerator

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestPathSwagger_ToString(t *testing.T) {
	t.Run("Test: PathSwagger.ToString()", func(t *testing.T) {
		t.Run("Should: return error "+errorIncorrectPath.Error(), func(t *testing.T) {
			a := NewPathSwagger("", nil)
			_, actual := a.ToString()
			expected := errorIncorrectPath
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return error "+errorNullRequests.Error(), func(t *testing.T) {
			a := NewPathSwagger("path", nil)
			_, actual := a.ToString()
			expected := errorNullRequests
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return error "+errorEmptyTypeRequest.Error(), func(t *testing.T) {
			requests := []RequestSwagger{
				NewRequestSwagger(nil, nil, nil),
			}
			a := NewPathSwagger("path", requests)
			_, actual := a.ToString()
			expected := errorEmptyTypeRequest
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: pass", func(t *testing.T) {
			swaggerParameters := map[string]interface{}{
				"typeRequest": "GET",
			}
			responseSwagger := NewResponseSwagger(200, "description", "", nil)
			requests := []RequestSwagger{
				NewRequestSwagger(swaggerParameters, nil, []ResponseSwagger{responseSwagger}),
			}
			a := NewPathSwagger("path", requests)
			actual, err := a.ToString()
			assert.NoError(t, err)

			str, _ := requests[0].ToString()
			expected := "\npath:" + strings.Replace(str, "\n", requestsIndentString, -1)
			assert.Equal(t, expected, actual)
		})
	})
}
