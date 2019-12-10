# swagger-gin-generator

# About

Swagger-gin-generation is a library developed to make generation of swagger description file more easy. 

It represents wrapper for gin router which allows to define swagger configuration during router initialization.

# Usage

## Download

To get library use

``` shell script
go get github.com/mcmakler/swagger-gin-generator
```

## Initialize wrapper

For wrapper using you need to initialize gin router and use it:

``` go
import "github.com/mcmakler/swagger-gin-generator/wrapper"

router := gin.New()
wr := wrapper.NewSwaggerRouterWrapper(
    wrapper.NewMainConfig(
        "1.0",
        "Example",
        "Usage example"),
    router)
```

Here function wrapper.NewMainConfig set basic swagger parameters:

``` go
wrapper.NewMainConfig(
    "1.0",            //Version of project
    "Example",        //Name of project
    "Usage example"   //Description
    )
```

The next step is setup authorization.

## Setup authorization

There are 6 types of authorization, which can be used together. 
First parameter in each method is swagger title, which need to be unique. 
Other params depends on security type.

``` go
wr.NewBasicSecurityDefinition("BasicSecurityTitle")
```

This method setup basic security with swagger title from parameter.

``` go
wr.NewApiKeySecurityDefinition(wrapper.SecurityBearer, wrapper.SecurityName, true)
```

Method setups new ApiKey authorization with swagger title wrapper.SecurityBearer and request parameter name wrapper.SecurityName.
If the third parameter is true, it will be in header, otherwise - in query.

``` go
authorizationUrl = "http://authorization.com"
tokenUrl = "http://token.com"
wr.NewOauth2ImplicitSecurityDefinition("O2ImpTitile", authorizationUrl)
wr.NewOauth2PasswordSecurityDefinition("O2PasTitile", tokenUrl)
wr.NewOauth2ApplicationSecurityDefinition("O2DefTitile", tokenUrl)
wr.NewOauth2AccessCodeSecurityDefinition("O2IAccTitile", authorizationUrl, tokenUrl)
```

These four methods represents  the Oath2 security definition. First parameter is swagger title. 
Other parameters are token URL or/and authorization URL and depends on Oauth2 authentication type.

The other functionality of wr is the same, as group.

## Make group

Group is analog of gin.group, but it has more parameters on initialization and response description.

Group creates in next way:

``` go
wrGroup := wr.Group("/url", "tag")
```

"/url" is the string with URL for gin. 
"tag" is the swagger tag, which will be used for each response and subgroup of given group, can be empty

You can use middlewares in group and create subgroups:

``` go
wrGroup.Use(someMiddleware, oneMoreMiddleware)
wrGroupSubgroup := wrGroup.Group("/subgroupurl", "subgroupTag")
```

As in gin router, you can add GET, POST, DELETE, HEAD, OPTIONS, PATCH and PUT request.
These functions has more parameters than in gin. Here is the GET request example.

## GET-example

First of all, the code:

``` go
wrGroup.GET(
    "/getpath",
    NewRequestConfig(
        "description",                  //description
        "operationId",                  //operationId
        "summary",                      //summary
        []string[                       //Array of security titles
            "BasicSecurityTitle",
            wrapper.SecurityBearer
        ], 
        []string{wrapper.TypesJson},    //Accept
        []string{wrapper.TypesBson},    //Produce
        []string{"getRequestTag"}       //Swagger tag
    ),
    parameterArray,
    responseMap,
    handlerFunction, oneMoreHandlerFunction
    )
```

The first parameter is GET path, the last parameters - gin handlers.

NewRequestConfig is function for setting swagger request configuration.
description, operationId and summary are strings, operationId need to be unique.
Array of security titles is string array which contains swagger titles for security definitions which are used in this request.
All the parameters can be empty strings or nil.

The parameterArray is arrays of wrapper.Parameter. Can be nil.

The responseMap is map of code-wrapper.Response.



## wrapper.Parameter

wrapper.Parameter is the representation of response parameter. 
It has two fields: swagger config and an exemplar of an object, which is parameter representation.
To create parameter call the function NewParameter:

``` go
param := wrapper.NewParameter(
    NewStringParameterConfig(...),
    ""
)
```

The fist parameter of this function is swagger configuration generator and the second is exemplar of an object.

There are 6 functions for configuration generating:

``` go
NewRequiredParameterConfig(in, name)
NewBasicParameterConfig(in, name, description, required)
NewArrayParameterConfig(in, name, description, required, minItems, maxItems, uniqueItems)
NewIntegerParameterConfig(in, name, description, required, defaultValue, min, max, multipleOf, exclusiveMin, exclusiveMax)
NewNumberParameterConfig(in, name, description, required, defaultValue, min, max, multipleOf, exclusiveMin, exclusiveMax)
NewStringParameterConfig(in, name, description, required, minLength, maxLength, pattern, enum)
```

NewRequiredParameterConfig has two basic parameters which are necessary: in and name.
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
    StringParam string
    MySubTypeParam MySubType
}

....

param := wrapper.NewParameter(
             NewRequiredParameterConfig("body","name"),
             &MyType{
                StringParam:    "string",
                MySubTypeParam: MySubType{IntParam: 10},
            }
         )
```

## Response

wrapper.Response is a structure, which is represented by response description and expected response object.
It generates with wrapper.NewResponse:

``` go
wrapper.NewResponse("description", &MyType)
```

The second parameter can be nil, in that case response will have only description.

In request responses are represented by map where key is http.status:

``` go
responsesMap := map[int]wrapper.Response{
    http.StatusOk:                  wrapper.NewResponse("ok", &MyType),
    http.StatusInternalServerError: wrapper.NewResponse("failure", nil),
}
```

## Path

Path can be used as group child, when there are some types of requests are send to one url:
``` go
path := wrGroup.Path("/pathUrl")
wrGroup.GET(
    NewRequestConfig(...),
    parameterArrayGet,
    responseMapGet,
    handlersGet...
    )
path.POST(
    NewRequestConfig(...),
    parameterArrayPost,
    responseMapPost,
    handlersPost...
    )
``` 

The difference with group is that you don't write a response url.

## Generation

There are three methods for generation.

``` go
wr.GenerateFiles(pathToFiles)
```

This method generates swagger.yaml and swagger.json files in the folder pathToFiles.

``` go
wr.GenerateBasePath("/swaggerPath")
```

Method set gin handler for path "/swaggerPath, where swagger will be represented.
Also use path "/swaggerPath/swagger.json", where json code is shown.

``` go
wr.GenerateWithoutSwagger()
```

This method generates no swagger and can be used on production.

# Author
- [Nikita Kharitonov](https://github.com/DreamAndDrum)
