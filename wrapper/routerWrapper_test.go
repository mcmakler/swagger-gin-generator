package wrapper

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"swagger-gin-generator/utils"
	"testing"
)

type testStructBool struct {
	B bool
}

type testResponse struct {
	Text string
}

type testStructFull struct {
	B      bool
	S      string
	I      int
	F      float64
	A      []bool
	Substr testStructBool
}

func TestNewSwaggerRouterWrapper(t *testing.T) {
	t.Run("Test: NewSwaggerRouterWrapper", func(t *testing.T) {
		g := gin.Default()
		srw := NewSwaggerRouterWrapper(
			map[string]interface{}{
				"title":       "title",
				"description": "description",
				"version":     "version",
			},
			g)
		emptyGroup := srw.Group("", "health")
		emptyPath := emptyGroup.Path("/health")
		emptyPath.Get(
			nil,
			nil,
			map[int]Request{
				200: {
					description: "getReqDef",
				},
			}, )

		firstGroup := srw.Group("/group1", "firstGroup")
		path1 := firstGroup.Path("/path1")
		path1.Get(
			map[string]interface{}{
				"description": "getDescription",
				"consumes":    []string{"getConsume"},
				"produces":    []string{"getProduce"},
				"summary":     "getSummary",
			},
			[]utils.Parameter{
				utils.NewParameter(map[string]interface{}{
					"name":        "name",
					"description": "boolGetParameter",
					"in":          "header",
				}, true),
			},
			map[int]Request{
				200: {
					description: "getReqDef",
					object:      &testStructFull{},
				},
				400: {
					description: "description",
					object:      nil,
				},
			},
			func(c *gin.Context) {})
		path1.Post(
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
					"description": "boolGetParameter",
					"in":          "header",
				}, true),
			},
			map[int]Request{
				200: {
					description: "getReqDef",
					object:      &testResponse{},
				},
			},
			func(c *gin.Context) {})

		path2 := firstGroup.Path("/path2")
		path2.Get(
			nil,
			nil,
			map[int]Request{
				200: {
					description: "getReqDef",
					object:      &testResponse{},
				},
			},
			func(c *gin.Context) {})
		path2.Post(
			nil,
			nil,
			map[int]Request{
				200: {
					description: "getReqDef",
					object:      &testResponse{},
				},
			},
			func(c *gin.Context) {})
		secondGroup := srw.Group("/group2", "secondGroup")
		paht22 := secondGroup.Path("/path2")
		paht22.Get(
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
					"description": "boolGetParameter",
					"in":          "header",
				}, true),
			},
			map[int]Request{
				200: {
					description: "getReqDef",
					object:      &testStructFull{},
				},
			},
			func(c *gin.Context) {})
		err := srw.Generate("rez.txt")
		assert.NoError(t, err)
	})
}
