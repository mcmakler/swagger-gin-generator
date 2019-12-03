package wrapper

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"swagger-gin-generator/swaggerFileGenerator/parameters"
	"swagger-gin-generator/utils"
	"testing"
)

func TestSwaggPathWrapper_Generate(t *testing.T) {
	t.Run("Test: SwaggPathWrapper", func(t *testing.T) {
		t.Run("Should: get path generate without errors", func(t *testing.T) {
			g := gin.Default()
			gr := g.Group("/")
			spw := newSwaggPathWrapper(
				"path",
				"tag",
				gr)
			spw.GET(
				map[string]interface{}{
					"description": "getDescription",
					"consumes":    []string{"getConsume"},
					"produces":    []string{"getProduce"},
					"tags":        []string{"tagget"},
					"summary":     "getSummary",
				},
				[]utils.Parameter{
					utils.NewParameter(map[string]interface{}{
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
			a := spw.generate()
			_, err := a.ToString()
			assert.NoError(t, err)
			expectedDefinitions := []parameters.SwaggParameter{
				utils.ConvertObjectToSwaggParameter(nil, true, false),
			}
			assert.Equal(t, expectedDefinitions, spw.getDefinitions())
		})

		t.Run("Should: post path generate without errors", func(t *testing.T) {
			g := gin.Default()
			gr := g.Group("/")
			spw := newSwaggPathWrapper(
				"path",
				"tag",
				gr)
			spw.POST(
				map[string]interface{}{
					"description": "getDescription",
					"consumes":    []string{"getConsume"},
					"produces":    []string{"getProduce"},
					"tags":        []string{"tagget"},
					"summary":     "getSummary",
				},
				[]utils.Parameter{
					utils.NewParameter(map[string]interface{}{
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
			a := spw.generate()
			_, err := a.ToString()
			assert.NoError(t, err)
		})

		t.Run("Should: other types test", func(t *testing.T) {
			g := gin.Default()
			gr := g.Group("/")
			spw := newSwaggPathWrapper(
				"path",
				"tag",
				gr)
			spw.DELETE(
				nil,
				nil,
				map[int]Request{
					200: {
						description: "getReqDef",
					},
				},
				func(c *gin.Context) {})
			spw.HEAD(
				nil,
				nil,
				map[int]Request{
					200: {
						description: "getReqDef",
					},
				},
				func(c *gin.Context) {})
			spw.OPTIONS(
				nil,
				nil,
				map[int]Request{
					200: {
						description: "getReqDef",
					},
				},
				func(c *gin.Context) {})
			spw.PATCH(
				nil,
				nil,
				map[int]Request{
					200: {
						description: "getReqDef",
					},
				},
				func(c *gin.Context) {})
			spw.PUT(
				nil,
				nil,
				map[int]Request{
					200: {
						description: "getReqDef",
					},
				},
				func(c *gin.Context) {})
			a := spw.generate()
			_, err := a.ToString()
			assert.NoError(t, err)
		})
	})
}
