package wrapper

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testPathObject struct {
}

func TestSwaggerPathWrapper_Generate(t *testing.T) {
	t.Run("Test: SwaggerPathWrapper", func(t *testing.T) {
		requestConfig := NewRequestConfig("description", "operationid", "summary", nil, []string{"consume"}, []string{"produce"}, []string{"tag"})
		parameters := []Parameter{
			NewParameter(
				NewBasicParameterConfig("body", "name", "boolGetParameter", true),
				true),
			NewParameter(
				NewBasicParameterConfig("body", "name", "boolGetParameter", true),
				&testPathObject{}),
		}
		boolRequests := map[int]Response{
			200: NewResponse("description", true),
		}
		emptyRequests := map[int]Response{
			200: NewResponse("description", nil),
		}
		emptyFunc := func(c *gin.Context) {}
		t.Run("Should: get path generate without errors", func(t *testing.T) {
			g := gin.Default()
			gr := g.Group("/")
			spw := newSwaggerPathWrapper(
				"path",
				"path",
				"tag",
				gr)
			spw.GET(requestConfig, parameters, boolRequests, emptyFunc)
			a := spw.generate()
			_, err := a.ToString()
			assert.NoError(t, err)
		})

		t.Run("Should: post path generate without errors", func(t *testing.T) {
			g := gin.Default()
			gr := g.Group("/")
			spw := newSwaggerPathWrapper(
				"path",
				"path",
				"tag",
				gr)
			spw.POST(requestConfig, parameters, boolRequests, emptyFunc)
			a := spw.generate()
			_, err := a.ToString()
			assert.NoError(t, err)
		})

		t.Run("Should: other types test", func(t *testing.T) {
			g := gin.Default()
			gr := g.Group("/")
			spw := newSwaggerPathWrapper(
				"path",
				"path",
				"tag",
				gr)
			spw.DELETE(requestConfig, parameters, emptyRequests, emptyFunc)
			spw.HEAD(requestConfig, parameters, emptyRequests, emptyFunc)
			spw.OPTIONS(requestConfig, parameters, emptyRequests, emptyFunc)
			spw.PATCH(requestConfig, parameters, emptyRequests, emptyFunc)
			spw.PUT(requestConfig, parameters, emptyRequests, emptyFunc)
			a := spw.generate()
			_, err := a.ToString()
			assert.NoError(t, err)
		})
	})
}
