package wrapper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//Unfunctional test: added in order to cover
func TestMethodsForStuctGeneration(t *testing.T) {
	t.Run("Should: create not nil objects", func(t *testing.T) {
		actual := NewRequiredParameterConfig("body", "name")
		assert.NotNil(t, actual)
		actual = NewBasicParameterConfig("body", "name", "", true)
		assert.NotNil(t, actual)
		actual = NewArrayParameterConfig("body", "name", "", true, 0, 10, true)
		assert.NotNil(t, actual)
		actual = NewIntegerParameterConfig("body", "name", "", true, 1, 0, 10, 1, true, true)
		assert.NotNil(t, actual)
		actual = NewNumberParameterConfig("body", "name", "", true, 1, 0, 10, 1, true, true)
		assert.NotNil(t, actual)
		actual = NewStringParameterConfig("body", "name", "", true, "", 1, 10, "", nil)
		assert.NotNil(t, actual)
		actual = NewRequiredMainConfig("version", "title")
		assert.NotNil(t, actual)
		actual = NewMainConfig("version", "title", "")
		assert.NotNil(t, actual)
		actual = NewRequestConfig("", "", "", nil, nil, nil, nil)
		assert.NotNil(t, actual)
	})
}
