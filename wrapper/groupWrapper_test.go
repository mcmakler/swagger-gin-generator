package wrapper

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSwaggerGroupWrapper(t *testing.T) {
	t.Run("Test: newSwaggerGroupWrapper", func(t *testing.T) {
		g := gin.Default()
		group := g.Group("path")
		sgw := newSwaggerGroupWrapper("path", "tag", group)
		sgw.Use(func(c *gin.Context) {})
		spw := sgw.Path("path")
		spw.GET(
			NewRequestConfig("description", "operationid", "summary", nil, []string{"consume"}, []string{"produce"}, []string{"tag"}),
			[]Parameter{
				NewParameter(
					NewBasicParameterConfig("in", "name", "boolGetParameter", true),
					true),
			},
			map[int]Response{
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

func TestSwaggerGroupWrapper(t *testing.T) {
	t.Run("Test of SwaggerGroupWrapper.GET, POST, ...", func(t *testing.T) {
		emptyRequests := map[int]Response{
			200: NewRequest("description", nil),
		}
		emptyFunc := func(c *gin.Context) {}
		g := gin.Default()
		group := g.Group("path")
		sgw := newSwaggerGroupWrapper("path", "tag", group)
		sgw.GET(
			"get",
			NewRequestConfig("", "", "", nil, nil, nil, nil),
			nil,
			emptyRequests,
			emptyFunc)
		sgw.POST(
			"post",
			NewRequestConfig("", "", "", nil, nil, nil, nil),
			nil,
			emptyRequests,
			emptyFunc)
		sgw.DELETE(
			"delete",
			NewRequestConfig("", "", "", nil, nil, nil, nil),
			nil,
			emptyRequests,
			emptyFunc)
		sgw.HEAD(
			"head",
			NewRequestConfig("", "", "", nil, nil, nil, nil),
			nil,
			emptyRequests,
			emptyFunc)
		sgw.OPTIONS(
			"options",
			NewRequestConfig("", "", "", nil, nil, nil, nil),
			nil,
			emptyRequests,
			emptyFunc)
		sgw.PATCH(
			"patch",
			NewRequestConfig("", "", "", nil, nil, nil, nil),
			nil,
			emptyRequests,
			emptyFunc)
		sgw.PUT(
			"put",
			NewRequestConfig("", "", "", nil, nil, nil, nil),
			nil,
			emptyRequests,
			emptyFunc)
	})
}
