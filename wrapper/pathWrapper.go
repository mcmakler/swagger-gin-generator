package wrapper

import (
	"github.com/gin-gonic/gin"
	"github.com/mcmakler/swagger-gin-generator/structures"
	"github.com/mcmakler/swagger-gin-generator/swaggerFileGenerator"
	"github.com/mcmakler/swagger-gin-generator/swaggerFileGenerator/parameters"
	"reflect"
)

type SwaggerPathWrapper interface {
	GET(config structures.Config, parameters []Parameter, requests map[int]Response, handlerFunc ...gin.HandlerFunc)
	POST(config structures.Config, parameters []Parameter, requests map[int]Response, handlerFunc ...gin.HandlerFunc)
	DELETE(config structures.Config, parameters []Parameter, requests map[int]Response, handlerFunc ...gin.HandlerFunc)
	HEAD(config structures.Config, parameters []Parameter, requests map[int]Response, handlerFunc ...gin.HandlerFunc)
	OPTIONS(config structures.Config, parameters []Parameter, requests map[int]Response, handlerFunc ...gin.HandlerFunc)
	PATCH(config structures.Config, parameters []Parameter, requests map[int]Response, handlerFunc ...gin.HandlerFunc)
	PUT(config structures.Config, parameters []Parameter, requests map[int]Response, handlerFunc ...gin.HandlerFunc)

	generate() swaggerFileGenerator.PathSwagger
	getDefinitions() []parameters.SwaggParameter
	readRequests(requests map[int]Response) []swaggerFileGenerator.ResponseSwagger
}

type swaggerPathWrapper struct {
	path        string
	tag         string
	requests    []swaggerFileGenerator.RequestSwagger
	definitions []parameters.SwaggParameter

	group gin.RouterGroup
}

type Response struct {
	description string
	object      interface{}
}

func NewResponse(description string, object interface{}) Response {
	return Response{
		description: description,
		object:      object,
	}
}

func newSwaggerPathWrapper(path, tag string, group gin.RouterGroup) SwaggerPathWrapper {
	return &swaggerPathWrapper{
		path:        path,
		tag:         tag,
		requests:    []swaggerFileGenerator.RequestSwagger{},
		definitions: []parameters.SwaggParameter{},
		group:       group,
	}
}

func (s *swaggerPathWrapper) GET(
	config structures.Config,
	inputParameters []Parameter,
	requests map[int]Response,
	handlerFuncP ...gin.HandlerFunc,
) {
	s.group.GET(s.path, handlerFuncP...)
	s.addRequest(config, inputParameters, requests, "get")
}

func (s *swaggerPathWrapper) POST(
	config structures.Config,
	inputParameters []Parameter,
	requests map[int]Response,
	handlerFuncP ...gin.HandlerFunc,
) {
	s.group.POST(s.path, handlerFuncP...)
	s.addRequest(config, inputParameters, requests, "post")
}

func (s *swaggerPathWrapper) DELETE(
	config structures.Config,
	inputParameters []Parameter,
	requests map[int]Response,
	handlerFuncP ...gin.HandlerFunc,
) {
	s.group.DELETE(s.path, handlerFuncP...)
	s.addRequest(config, inputParameters, requests, "delete")
}

func (s *swaggerPathWrapper) HEAD(
	config structures.Config,
	inputParameters []Parameter,
	requests map[int]Response,
	handlerFuncP ...gin.HandlerFunc,
) {
	s.group.HEAD(s.path, handlerFuncP...)
	s.addRequest(config, inputParameters, requests, "head")
}

func (s *swaggerPathWrapper) OPTIONS(
	config structures.Config,
	inputParameters []Parameter,
	requests map[int]Response,
	handlerFuncP ...gin.HandlerFunc,
) {
	s.group.OPTIONS(s.path, handlerFuncP...)
	s.addRequest(config, inputParameters, requests, "options")
}

func (s *swaggerPathWrapper) PATCH(
	config structures.Config,
	inputParameters []Parameter,
	requests map[int]Response,
	handlerFuncP ...gin.HandlerFunc,
) {
	s.group.PATCH(s.path, handlerFuncP...)
	s.addRequest(config, inputParameters, requests, "patch")
}

func (s *swaggerPathWrapper) PUT(
	config structures.Config,
	inputParameters []Parameter,
	requests map[int]Response,
	handlerFuncP ...gin.HandlerFunc,
) {
	s.group.PUT(s.path, handlerFuncP...)
	s.addRequest(config, inputParameters, requests, "put")
}

func (s *swaggerPathWrapper) generate() swaggerFileGenerator.PathSwagger {
	res := swaggerFileGenerator.NewPathSwagger(s.path, s.requests)
	return res
}

func (s *swaggerPathWrapper) getDefinitions() []parameters.SwaggParameter {
	return s.definitions
}

func (s *swaggerPathWrapper) addRequest(
	structConfigs structures.Config,
	inputParameters []Parameter,
	requests map[int]Response,
	reqType string,
) {
	var swaggerParameters []parameters.SwaggParameter
	for _, val := range inputParameters {
		if val.getSwaggerParameter(false).IsObject() {
			swaggerParameters = append(swaggerParameters, val.getSwaggerParameter(true))
			s.definitions = append(s.definitions, val.getSwaggerParameter(false))
		} else {
			swaggerParameters = append(swaggerParameters, val.getSwaggerParameter(true))
		}
	}

	var configs map[string]interface{}
	if structConfigs != nil {
		configs = structConfigs.ToMap()
	} else {
		configs = make(map[string]interface{})
	}

	if s.tag != "" {
		if _, ok := configs["tags"]; !ok {
			configs["tags"] = []string{}
		}
		configs["tags"] = append(configs["tags"].([]string), s.tag)
	}

	configs["typeRequest"] = reqType
	s.requests = append(s.requests, swaggerFileGenerator.NewRequestSwagger(
		configs,
		swaggerParameters,
		s.readRequests(requests),
	))
}

func (s *swaggerPathWrapper) readRequests(requests map[int]Response) []swaggerFileGenerator.ResponseSwagger {
	var responses []swaggerFileGenerator.ResponseSwagger
	//TODO: CHECK USUAL PARAMETER (STRING, BOOL, ...)

	for key, val := range requests {
		if val.object != nil {
			elemTypeName := reflect.TypeOf(val.object).Name()
			if reflect.ValueOf(val.object).Kind() == reflect.Ptr {
				elemTypeName = reflect.ValueOf(val.object).Elem().Type().Name()
			}
			element := ReturnNonStructureObject(nil, val.object)
			respSwag := swaggerFileGenerator.NewResponseSwagger(key, val.description, elemTypeName, element)
			responses = append(responses, respSwag)
			if element == nil {
				s.definitions = append(s.definitions, ConvertObjectToSwaggerParameter(nil, val.object, false))
			}
		} else {
			respSwag := swaggerFileGenerator.NewResponseSwagger(key, val.description, "", nil)
			responses = append(responses, respSwag)
		}
	}

	return responses
}
