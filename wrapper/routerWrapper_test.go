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
	B      bool   `binding:"required"`
	S      string `binding:"required"`
	I      int
	F      float64
	A      []bool
	Substr testStructBool
}

type testObjectStruct struct {
	B      bool
	S      string
	I      int
	F      float64 `binding:"required"`
	A      []bool  `binding:"required"`
	Substr testStructBool
}

func TestNewSwaggerRouterWrapper(t *testing.T) {
	t.Run("Test: NewSwaggerRouterWrapper", func(t *testing.T) {
		parameters := []Parameter{
			NewParameter(
				NewBasicParameterConfig("header", "name", "boolGetParameter", true),
				true),
			NewParameter(
				NewBasicParameterConfig("body", "object", "object Parameter", true),
				&testObjectStruct{}),
		}

		g := gin.Default()
		srw := NewSwaggerRouterWrapper(
			NewRequiredMainConfig("version", "title"),
			g)
		srw = NewSwaggerRouterWrapper(
			NewMainConfig("version", "title", "description", "example", "/"),
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
		srw.Use(func(c *gin.Context) {})
		emptyGroup := srw.Group("", "health")
		emptyPath := emptyGroup.Path("/health")
		emptyPath.GET(
			nil,
			nil,
			map[int]Response{
				200: {
					description: "getReqDef",
				},
			}, )

		firstGroup := srw.Group("/group1", "firstGroup")
		path1 := firstGroup.Path("/path1")
		path1.GET(
			NewRequestConfig("description", "", "summary", []string{basicSecurity, apiSecurity, oauth2AccessCodeSecurity}, []string{"consume"}, []string{"produce"}, []string{"tag"}),
			parameters,
			map[int]Response{
				200: {
					description: "getReqDef",
					object:      "",
				},
				400: {
					description: "description",
					object:      nil,
				},
			},
			func(c *gin.Context) {})
		path1.POST(
			NewRequestConfig("description", "", "summary", []string{oauth2ImplicitSecurity}, []string{"consume"}, []string{"produce"}, []string{"tag"}),
			parameters,
			map[int]Response{
				200: {
					description: "getReqDef",
					object:      &testResponse{},
				},
			},
			func(c *gin.Context) {})

		firstGroup.GET(
			"/path2",
			nil,
			nil,
			map[int]Response{
				200: {
					description: "getReqDef",
					object:      &testResponse{},
				},
			},
			func(c *gin.Context) {})
		firstGroup.POST(
			"/path2",
			nil,
			nil,
			map[int]Response{
				200: {
					description: "getReqDef",
					object:      &testResponse{},
				},
			},
			func(c *gin.Context) {})

		oneAndHalfGroup := srw.Group("/oneAdHalf", "oneAdHalfGroup")
		secondGroup := oneAndHalfGroup.Group("/group2", "secondGroup")
		paht22 := secondGroup.Path("/path2")
		paht22.DELETE(
			NewRequestConfig("description", "operationid3", "summary", []string{oauth2PasswordSecurity, oauth2AppSecurity}, []string{"consume"}, []string{"produce"}, []string{"tag"}),
			parameters,
			map[int]Response{
				200: {
					description: "getReqDef",
					object:      &testStructFull{},
				},
			},
			func(c *gin.Context) {})
		paht22.PATCH(
			NewRequestConfig("description", "operationid2", "summary", nil, []string{"consume"}, []string{"produce"}, []string{"tag"}),
			[]Parameter{
				NewParameter(
					NewIntegerParameterConfig(InBody, "name", "descr", true, 1, 0, 10, 2, true, true),
					10),
			},
			map[int]Response{
				200: {
					description: "getReqDef",
					object:      &testStructFull{},
				},
			},
			func(c *gin.Context) {})
		paht22.OPTIONS(
			NewRequestConfig("description", "operationid1", "summary", nil, []string{"consume"}, []string{"produce"}, []string{"tag"}),
			[]Parameter{
				NewParameter(
					NewNumberParameterConfig(InHeader, "name", "descr", true, 1.0, 0.0, 10.0, 2.0, true, true),
					10.0),
			},
			map[int]Response{
				200: {
					description: "getReqDef",
					object:      &testStructFull{},
				},
			},
			func(c *gin.Context) {})
		paht22.HEAD(
			NewRequestConfig("description", "", "summary", nil, []string{"consume"}, []string{"produce"}, []string{"tag"}),
			[]Parameter{
				NewParameter(
					NewStringParameterConfig(InBody, "name", "descr", true, "format", 0, 10, "dfdf", nil),
					""),
			},
			map[int]Response{
				200: {
					description: "getReqDef",
					object:      &testStructFull{},
				},
			},
			func(c *gin.Context) {})
		paht22.PUT(
			NewRequestConfig("description", "", "summary", nil, []string{"consume"}, []string{"produce"}, []string{"tag"}),
			[]Parameter{
				NewParameter(
					NewArrayParameterConfig("body", "name", "descr", true, 0, 10, true),
					[]string{}),
			},
			map[int]Response{
				200: {
					description: "getReqDef",
					object:      &testStructFull{},
				},
			},
			func(c *gin.Context) {})
		err := srw.GenerateBasePath("api/doc")
		assert.NoError(t, err)
		err = srw.GenerateFiles("")
		assert.NoError(t, err)
	})
}

func TestSwaggerRouterWrapper(t *testing.T) {
	t.Run("Test of SwaggerRouterWrapper.GET, POST, ...", func(t *testing.T) {
		emptyRequests := map[int]Response{
			200: NewResponse("description", nil),
		}
		emptyFunc := func(c *gin.Context) {}
		g := gin.Default()
		sgw := NewSwaggerRouterWrapper(
			NewMainConfig("version", "title", "description", "example", "/"),
			g)
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
