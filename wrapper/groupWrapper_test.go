package wrapper

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"swagger-gin-generator/swaggerFileGenerator/parameters"
	"swagger-gin-generator/utils"
	"testing"
)

func TestNewSwaggGroupWrapper(t *testing.T) {
	t.Run("Test: NewSwaggGroupWrapper", func(t *testing.T) {
		g := gin.Default()
		group := g.Group("path")
		sgw := NewSwaggGroupWrapper("path", "tag", group)
		sgw.Use(func(c *gin.Context) {})
		spw := sgw.Path("path")
		spw.Get(
			"getDescription",
			[]string{"getConsume"},
			[]string{"getProduce"},
			nil,
			"getSummary",
			[]utils.Parameter{
				utils.NewParameter(map[string]interface{}{
					"description": "boolGetParameter",
				}, true),
			},
			map[int]Request{
				200: {
					definition: "getReqDef",
					object:     true,
				},
			},
			func(c *gin.Context){})
		a := sgw.Generate()
		var err error
		for _, val := range a {
			_, err = val.ToString()
			if err != nil {
				break
			}
		}
		assert.NoError(t, err)
		expectedDefinitions := []parameters.SwaggParameter{
			utils.ConvertObjectToSwaggParameter(nil, true),
		}
		assert.Equal(t, expectedDefinitions, sgw.Definitions())
	})
}
