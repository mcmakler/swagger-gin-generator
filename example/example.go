package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mcmakler/swagger-gin-generator/wrapper"
	"net/http"
)

type MySubType struct {
	IntParam int
}
type MyType struct {
	StringParam    string `binding:"required"`
	MySubTypeParam MySubType
}

func someMiddleware(c *gin.Context)         {}
func oneMoreMiddleware(c *gin.Context)      {}
func handlerFunction(c *gin.Context)        {}
func oneMoreHandlerFunction(c *gin.Context) {}

func main() {
	err := RunExample()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Files generated")
}

func RunExample() error {
	router := gin.New()
	wr := wrapper.NewSwaggerRouterWrapper(
		wrapper.NewMainConfig(
			"1.0",
			"Example",
			"Usage example",
			"example.com",
			"/basePath"),
		router)

	//Shall use all the types of authorization
	wr.NewBasicSecurityDefinition("BasicSecurityTitle")
	wr.NewApiKeySecurityDefinition(
		wrapper.SecurityBearer,
		wrapper.SecurityName,
		true)

	authorizationUrl := "http://authorization.com"
	tokenUrl := "http://token.com"
	wr.NewOauth2ImplicitSecurityDefinition(
		"O2ImpTitle",
		authorizationUrl)
	wr.NewOauth2PasswordSecurityDefinition(
		"O2PasTitle",
		tokenUrl)
	wr.NewOauth2ApplicationSecurityDefinition(
		"O2DefTitle",
		tokenUrl)
	wr.NewOauth2AccessCodeSecurityDefinition(
		"O2IAccTitle",
		authorizationUrl,
		tokenUrl)

	//Creating of group
	wrGroup := wr.Group("/url", "tag")

	//Creating of parameter array
	param := wrapper.NewParameter(
		wrapper.NewRequiredParameterConfig(wrapper.InHeader, "myParam"),
		"",
	)
	myParam := wrapper.NewParameter(
		wrapper.NewRequiredParameterConfig(wrapper.InBody, "name"),
		&MyType{
			StringParam:    "string",
			MySubTypeParam: MySubType{IntParam: 10},
		},
	)
	parameterArray := []wrapper.Parameter{param, myParam}

	//Creating of response map
	responseMap := map[int]wrapper.Response{
		http.StatusOK:                  wrapper.NewResponse("ok", &MyType{}),
		http.StatusInternalServerError: wrapper.NewResponse("failure", nil),
	}

	//GET example
	wrGroup.GET(
		"/getpath",
		wrapper.NewRequestConfig(
			"description", //description
			"operationId", //operationId
			"summary",     //summary
			[]string{ //Array of security titles
				"BasicSecurityTitle",
				wrapper.SecurityBearer,
			},
			[]string{wrapper.TypesJson}, //Accept
			[]string{wrapper.TypesBson}, //Produce
			[]string{"getRequestTag"},   //Swagger tag
		),
		parameterArray,
		responseMap,
		handlerFunction, oneMoreHandlerFunction,
	)

	//Using middleware and creating subgroup
	wrGroup.Use(someMiddleware, oneMoreMiddleware)
	wrGroupSubgroup := wrGroup.Group("/subgroupurl", "subgroupTag")

	//Empty POST subgroup
	wrGroupSubgroup.POST(
		"/postpath",
		nil,
		nil,
		responseMap,
		handlerFunction, oneMoreHandlerFunction,
	)

	//Path example
	path := wrGroup.Path("/pathUrl")
	path.DELETE(
		nil,
		nil,
		map[int]wrapper.Response{
			http.StatusOK: wrapper.NewResponse("Ok", nil),
		},
		handlerFunction, oneMoreHandlerFunction,
	)
	path.HEAD(
		nil,
		nil,
		map[int]wrapper.Response{
			http.StatusOK: wrapper.NewResponse("Ok", nil),
		},
		handlerFunction, oneMoreHandlerFunction,
	)

	err := wr.GenerateFiles("")
	return err
}
