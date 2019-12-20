# swagger-gin-generator

# About

Swagger-gin-generation is a library developed to make generation of swagger description file easier. 

It represents wrapper for gin router which allows to define swagger configuration during router initialization.

# Example

Here is [example](https://github.com/mcmakler/swagger-gin-generator/tree/master/example) of library using.

# Problems to solve

- Require hand-changing of structure to the other in request/response swagger definition
- Require hand-adding of new responses in swagger definition

# Usage

## Download

In order to get library use

``` shell script
go get github.com/mcmakler/swagger-gin-generator
```

## Initialize wrapper

For wrapper use you need to initialize gin router and use it:

``` go
import "github.com/mcmakler/swagger-gin-generator/wrapper"

router := gin.New()
wr := wrapper.NewSwaggerRouterWrapper(
    wrapper.NewMainConfig(
        "1.0",
        "Example",
        "Usage example"
        "example.com",
        "/basePath"),
    router)
```

Here is function wrapper.NewMainConfig set basic swagger parameters:

``` go
wrapper.NewMainConfig(
    "1.0",            //Version of project
    "Example",        //Name of project
    "Usage example"   //Description,
    "example.com",    //host
    "/basePath"       //basePath
    )
```

The .yaml code generated in this part is:

``` yaml
swagger: '2.0'
info:
  title: Example
  version: '1.0'
  description: Usage example
host: example.com
basePath: /basePath
```

The next step is the setup authorization.

## Setup authorization

There are 6 types of authorization, which can be used together. 
The first parameter in each method is a swagger title, which needs to be unique. 
Other parameters depends on security type.

``` go
wr.NewBasicSecurityDefinition("BasicSecurityTitle")
```

This method setups basic security with the swagger title from parameter.

``` go
wr.NewApiKeySecurityDefinition(wrapper.SecurityBearer, wrapper.SecurityName, true)
```

The method setups new ApiKey authorization with a swagger title wrapper.SecurityBearer and request parameter name wrapper.SecurityName.
If the third parameter is true, it will be in header, otherwise - in query.

``` go
authorizationUrl := "http://authorization.com"
tokenUrl := "http://token.com"
wr.NewOauth2ImplicitSecurityDefinition("O2ImpTitle", authorizationUrl)
wr.NewOauth2PasswordSecurityDefinition("O2PasTitle", tokenUrl)
wr.NewOauth2ApplicationSecurityDefinition("O2DefTitle", tokenUrl)
wr.NewOauth2AccessCodeSecurityDefinition("O2IAccTitle", authorizationUrl, tokenUrl)
```

These four methods represent the Oath2 security definition. The first parameter is swagger title. 
Other parameters are token URL or/and authorization URL and depend on Oauth2 authentication type.

Security commands will generate the next .yaml code:

``` yaml
securityDefinitions:
  BasicSecurityTitle:
    type: basic
  Bearer:
    type: apiKey
    name: Authorization
    in: header
  O2ImpTitle:
    type: oauth2
    flow: implicit
    authorizationUrl: http://authorization.com
  O2PasTitle:
    type: oauth2
    flow: password
    tokenUrl: http://token.com
  O2DefTitle:
    type: oauth2
    flow: application
    tokenUrl: http://token.com
  O2IAccTitle:
    type: oauth2
    flow: accessCode
    authorizationUrl: http://authorization.com
    tokenUrl: http://token.com
```

The other functionality of wr is the same, as for group.

## Make group

A group is an analog of gin.group, but has more parameters on initialization and response description.

The group creates in next way:

``` go
wrGroup := wr.Group("/url", "tag")
```

"/url" is the string with URL for gin. 
"tag" is the swagger tag, which will be used for each response and subgroup of given group, can be empty

You can use middlewares in a group and create subgroups:

``` go
wrGroup.Use(someMiddleware, oneMoreMiddleware)
wrGroupSubgroup := wrGroup.Group("/subgroupurl", "subgroupTag")
```

As in gin router, you can add GET, POST, DELETE, HEAD, OPTIONS, PATCH and PUT request.
These functions have more parameters than in gin. Here is the GET request example.

## GET-example

First of all, the code:

``` go
wrGroup.GET(
    "/getpath",
    wrapper.NewRequestConfig(
        "description",                  //description
        "operationId",                  //operationId
        "summary",                      //summary
        []string[                       //Array of security titles
            "BasicSecurityTitle",
            wrapper.SecurityBearer,
        ], 
        []string{wrapper.TypesJson},    //Accept
        []string{wrapper.TypesBson},    //Produce
        []string{"getRequestTag"},      //Swagger tag
    ),
    parameterArray,
    responseMap,
    handlerFunction, oneMoreHandlerFunction,
    )
```

The first parameter is GET path, the last parameters - gin handlers.

NewRequestConfig is a function for setting swagger request configuration.
Description, operationId and summary are strings, operationId needs to be unique.
An array of security titles is a string array which contains swagger titles for security definitions which are used in this request.
All the parameters can be empty strings or nil.

The parameterArray is an array of wrapper.Parameter. Can be nil.

The responseMap is a map of code-wrapper.Response.

The result of GET method generation is (parameters and responses are omitted): 

``` yaml
paths:
  /url/getpath:
    get:
      security:
      - BasicSecurityTitle: []
      - Bearer: []
      description: description
      consumes:
      - json
      produces: 
      - bson
      tags: 
        - getRequestTag
        - tag               #this tag is added by group definition
      operationId: operationId
      summary: summary
      parameters:
        ...
      responses:
        ...
```

## wrapper.Parameter

wrapper.Parameter is the representation of response parameter. 
It has two fields: swagger config and an exemplar of an object, which is a parameter representation.
To create a parameter call the function NewParameter:

``` go
param := wrapper.NewParameter(
        wrapper.NewRequiredParameterConfig(wrapper.InHeader, "myParam"),
        "",
    )
```

The first parameter of this function is a swagger configuration generator and the second is exemplar of an object.

There are 6 functions for configuration generating:

``` go
wrapper.NewRequiredParameterConfig(in, name)
wrapper.NewBasicParameterConfig(in, name, description, required)
wrapper.NewArrayParameterConfig(in, name, description, required, minItems, maxItems, uniqueItems)
wrapper.NewIntegerParameterConfig(in, name, description, required, defaultValue, min, max, multipleOf, exclusiveMin, exclusiveMax)
wrapper.NewNumberParameterConfig(in, name, description, required, defaultValue, min, max, multipleOf, exclusiveMin, exclusiveMax)
wrapper.NewStringParameterConfig(in, name, description, required, minLength, maxLength, pattern, enum)
```

NewRequiredParameterConfig has two necessary basic parameters: in and name.
NewBasicParameterConfig has two additional parameters: description and required.

Other four function are used for Array, Integer, Number and String parameter definition in
cases, when it is necessary to set additional parameters, such as minItems for arrays or max for Numbers.

Description can be set as "", enum can be set as nil.

The setting of parameter type is going on automatically depending on object exemplar.
You can use your own types, if all the fields are PUBLIC:

``` go
type MySubType struct {
	IntParam int
}
type MyType struct {
	StringParam    string     `binding:"required"`
	MySubTypeParam MySubType
}

....

myParam := wrapper.NewParameter(
             wrapper.NewRequiredParameterConfig(wrapper.InBody,"name"),
             &MyType{
                StringParam:    "string",
                MySubTypeParam: MySubType{IntParam: 10},
            },
         )
```

This code defined structure in definitions and add ref to GET:

``` yaml
/url/getpath:
  get:
    ...
    parameters:
      - in: body
        name: name
        schema:
          $ref: '#/definitions/MyType'
...
definitions:
  MyType:
    type: object
    required: 
      - StringParam
    properties:
      StringParam:
        type: string
      MySubTypeParam:
        type: object
        properties:
          IntParam:
            type: integer
```

## wrapper.Response

wrapper.Response is a structure, which is represented by response description and expected response object.
It generates with wrapper.NewResponse:

``` go
wrapper.NewResponse("description", &MyType)
```

The second parameter can be nil, in that case response will have only description.

In request responses are represented by map where key is http.status:

``` go
responseMap := map[int]wrapper.Response{
    http.StatusOK:                  wrapper.NewResponse("ok", &MyType{}),
    http.StatusInternalServerError: wrapper.NewResponse("failure", nil),
}
```

In GET definition in .yaml this code will define:

``` yaml
/url/getpath:
  get:
    ...
    responses:
      '200':
        description: ok
        schema:
          $ref: '#/definitions/MyType'
      '500':
        description: failure
```

## Path

Path can be used as group child, when some types of requests are sent to one url:
``` go
path := wrGroup.Path("/pathUrl")
path.DELETE(
    NewRequestConfig(...),
    parameterArrayDelete,
    responseMapDelete,
    handlersDelete...
    )
path.HEAD(
    NewRequestConfig(...),
    parameterArrayHead,
    responseMapHead,
    handlersHead...
    )
``` 

The difference with group is that you don't write a response url.

## Generation

There are three methods for generation.

``` go
err := wr.GenerateFiles(pathToFiles)
```

This method generates swagger.yaml and swagger.json files in the folder pathToFiles.

``` go
err := wr.GenerateBasePath("/swaggerPath")
```

Method set gin handler for path "/swaggerPath, where swagger will be represented.
Also use path "/swaggerPath/swagger.json", where json code is shown.

These methods can be omitted in production.

# Author
- [Nikita Kharitonov](https://github.com/DreamAndDrum)
