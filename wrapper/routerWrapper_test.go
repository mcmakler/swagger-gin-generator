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
		firstGroup := srw.Group("/group1", "firstGroup")
		path1 := firstGroup.Path("/path1")
		path1.Get(
			"getDescription",
			[]string{"getConsume"},
			[]string{"getProduce"},
			nil,
			"getSummary",
			[]utils.Parameter{
				utils.NewParameter(map[string]interface{}{
					"description": "boolGetParameter",
					"in":          "header",
				}, true),
			},
			map[int]Request{
				200: {
					definition: "getReqDef",
					object: testStructFull{
						B:      false,
						S:      "",
						I:      0,
						F:      0,
						A:      nil,
						Substr: testStructBool{B: false},
					},
				},
			},
			func(c *gin.Context) {})
		path1.Post(
			"getDescription",
			[]string{"getConsume"},
			[]string{"getProduce"},
			nil,
			"getSummary",
			[]utils.Parameter{
				utils.NewParameter(map[string]interface{}{
					"description": "boolGetParameter",
					"in":          "header",
				}, true),
			},
			map[int]Request{
				200: {
					definition: "getReqDef",
					object:     testResponse{"sdsd"},
				},
			},
			func(c *gin.Context) {})

		path2 := firstGroup.Path("/path2")
		path2.Get(
			"getDescription",
			[]string{"getConsume"},
			[]string{"getProduce"},
			nil,
			"getSummary",
			[]utils.Parameter{
				utils.NewParameter(map[string]interface{}{
					"description": "boolGetParameter",
					"in":          "header",
				}, true),
			},
			map[int]Request{
				200: {
					definition: "getReqDef",
					object:     testResponse{"sdsd"},
				},
			},
			func(c *gin.Context) {})
		path2.Post(
			"getDescription",
			[]string{"getConsume"},
			[]string{"getProduce"},
			nil,
			"getSummary",
			[]utils.Parameter{
				utils.NewParameter(map[string]interface{}{
					"description": "boolGetParameter",
					"in":          "header",
				}, true),
			},
			map[int]Request{
				200: {
					definition: "getReqDef",
					object:     testResponse{"sdsd"},
				},
			},
			func(c *gin.Context) {})
		secondGroup := srw.Group("/group2", "secondGroup")
		paht22 := secondGroup.Path("/path2")
		paht22.Get(
			"getDescription",
			[]string{"getConsume"},
			[]string{"getProduce"},
			nil,
			"getSummary",
			[]utils.Parameter{
				utils.NewParameter(map[string]interface{}{
					"description": "boolGetParameter",
					"in":          "header",
				}, true),
			},
			map[int]Request{
				200: {
					definition: "getReqDef",
					object: testStructFull{
						B:      false,
						S:      "",
						I:      0,
						F:      0,
						A:      nil,
						Substr: testStructBool{B: false},
					},
				},
			},
			func(c *gin.Context) {})
		err := srw.Generate("rez.txt")
		assert.NoError(t, err)
	})
}
