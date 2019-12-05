package swaggerFileGenerator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResponseSwagg_ToString(t *testing.T) {
	t.Run("Test: ResponseSwagg.ToString()", func(t *testing.T) {
		t.Run("Should: return wrong http code error", func(t *testing.T) {
			a := &responseSwagg{
				code:         -1,
				description:  "de",
				linkOnSchema: "",
			}
			_, error := a.ToString()
			assert.Equal(t, errorWrongCode, error)
		})

		t.Run("Should: return empty response swag", func(t *testing.T) {
			a := &responseSwagg{
				code:         200,
				description:  "",
				linkOnSchema: "",
			}
			expected := errorEmptyDescription
			_, actual := a.ToString()
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return response swag", func(t *testing.T) {
			a := &responseSwagg{
				code:         200,
				description:  "description",
				linkOnSchema: "linkonschema",
			}
			expected := "\n'200':" +
				descriptionString + "description" +
				linkOnSchemaString + refString + "linkonschema'"
			actual, _ := a.ToString()
			assert.Equal(t, expected, actual)
		})
	})
}

func TestNewResponseSwagg(t *testing.T) {
	t.Run("Should: return response swag", func(t *testing.T) {
		expected := &responseSwagg{
			code:         200,
			description:  "description",
			linkOnSchema: "linkonschema",
		}
		actual := NewResponseSwagg(200, "description", "linkonschema", nil)
		assert.Equal(t, expected, actual)
	})
}
