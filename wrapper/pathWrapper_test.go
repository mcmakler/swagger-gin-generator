package wrapper

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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
			a := spw.generate()
			_, err := a.ToString()
			assert.NoError(t, err)
		})

		t.Run("Should: post path generate without errors", func(t *testing.T) {
			g := gin.Default()
			gr := g.Group("/")
			spw := newSwaggPathWrapper(
				"path",
				"tag",
				gr)
			spw.POST(
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
