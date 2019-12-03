package wrapper

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"swagger-gin-generator/swaggerFileGenerator/parameters"
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
			map[string]interface{}{
				"description": "getDescription",
				"consumes":    []string{"getConsume"},
				"produces":    []string{"getProduce"},
				"tags":        []string{"tagget"},
				"summary":     "getSummary",
			},
			[]Parameter{
				NewParameter(map[string]interface{}{
					"name":        "name",
					"in":          "in",
					"description": "boolGetParameter",
				}, true),
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
		expectedDefinitions := []parameters.SwaggParameter{
			ConvertObjectToSwaggParameter(nil, true, false),
		}
		assert.Equal(t, expectedDefinitions, sgw.getDefinitions())
	})
}
