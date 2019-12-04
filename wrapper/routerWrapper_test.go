package wrapper

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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
		basicSecurity := "Basic"
		srw.NewBasicSecurityDefinition(basicSecurity)
		apiSecurity := "ApiKey"
		srw.NewApiKeySecurityDefinition(apiSecurity, "Bearer", true)
		oauth2AccessCodeSecurity := "Oauth2Access"
		srw.NewOauth2AccessCodeSecurityDefinition(oauth2AccessCodeSecurity, "http://auth.com", "http://token.com")
		oauth2ImplicitSecurity := "Oauth2Implicit"
		srw.NewOauth2ImplicitSecurityDefinition(oauth2ImplicitSecurity, "http://auth.com")
		oauth2PasswordSecurity := "Oauth2Password"
		srw.NewOauth2PasswordSecurityDefinition(oauth2PasswordSecurity, "http://token.com")
		oauth2AppSecurity := "Oauth2App"
		srw.NewOauth2ApplicationSecurityDefinition(oauth2AppSecurity, "http://token.com")
		emptyGroup := srw.Group("", "health")
		emptyPath := emptyGroup.Path("/health")
		emptyPath.GET(
			nil,
			nil,
			map[int]Request{
				200: {
					description: "getReqDef",
				},
			}, )

		firstGroup := srw.Group("/group1", "firstGroup")
		path1 := firstGroup.Path("/path1")
		path1.GET(
			map[string]interface{}{
				"description": "getDescription",
				"security":    []string{basicSecurity, apiSecurity, oauth2AccessCodeSecurity},
				"consumes":    []string{"getConsume"},
				"produces":    []string{"getProduce"},
				"summary":     "getSummary",
			},
			[]Parameter{
				NewParameter(map[string]interface{}{
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
		path1.POST(
			map[string]interface{}{
				"description": "getDescription",
				"security":    []string{oauth2ImplicitSecurity},
				"consumes":    []string{"getConsume"},
				"produces":    []string{"getProduce"},
				"tags":        []string{"tagget"},
				"summary":     "getSummary",
			},
			[]Parameter{
				NewParameter(map[string]interface{}{
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
		path2.GET(
			nil,
			nil,
			map[int]Request{
				200: {
					description: "getReqDef",
					object:      &testResponse{},
				},
			},
			func(c *gin.Context) {})
		path2.POST(
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
		paht22.DELETE(
			map[string]interface{}{
				"description": "getDescription",
				"security":    []string{oauth2PasswordSecurity, oauth2AppSecurity},
				"consumes":    []string{"getConsume"},
				"produces":    []string{"getProduce"},
				"tags":        []string{"tagget"},
				"summary":     "getSummary",
			},
			[]Parameter{
				NewParameter(map[string]interface{}{
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
		paht22.PATCH(
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
		paht22.OPTIONS(
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
		paht22.HEAD(
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
		paht22.PUT(
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
		err := srw.Generate("")
		assert.NoError(t, err)
	})
}
