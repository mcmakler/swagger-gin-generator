package example

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunExample(t *testing.T) {
	t.Run("Running of example", func(t *testing.T) {
		err := RunExample()
		assert.NoError(t, err)
	})
}
