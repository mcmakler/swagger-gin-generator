package wrapper

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSwaggGroupWrapper(t *testing.T) {
	t.Run("Test: newSwaggGroupWrapper", func(t *testing.T) {
		g := gin.Default()
		group := g.Group("path")
		sgw := newSwaggGroupWrapper("path", "tag", group)
		sgw.Use(func(c *gin.Context) {})
		spw := sgw.Path("path")
		spw.GET(
			NewRequestConfig("description", "operationid", "summary", nil, []string{"consume"}, []string{"produce"}, []string{"tag"}),
			[]Parameter{
				NewParameter(
					NewBasicParameterConfig("in", "name", "boolGetParameter", true),
					true),
			},
			map[int]Request{
				200: {
					description: "getReqDef",
					object:      true,
				},
			},
			func(c *gin.Context) {})
		a := sgw.generate()
		var err error
		for _, val := range a {
			_, err = val.ToString()
			if err != nil {
				break
			}
		}
		assert.NoError(t, err)
	})
}
