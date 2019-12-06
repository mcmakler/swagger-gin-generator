package wrapper

import (
	"github.com/gin-gonic/gin"
	"github.com/mcmakler/swagger-gin-generator/structures"
	"github.com/mcmakler/swagger-gin-generator/swaggerFileGenerator"
	"github.com/mcmakler/swagger-gin-generator/swaggerFileGenerator/parameters"
	"reflect"
)

type SwaggPathWrapper interface {
	GET(
		swaggerParameters structures.Config,
		parameters []Parameter,
		requests map[int]Request,
		handlerFunc ...gin.HandlerFunc,
	)
	POST(
		swaggerParameters structures.Config,
		parameters []Parameter,
		requests map[int]Request,
		handlerFunc ...gin.HandlerFunc,
	)
	DELETE(
		swaggerParameters structures.Config,
		parameters []Parameter,
		requests map[int]Request,
		handlerFunc ...gin.HandlerFunc,
	)
	HEAD(
		swaggerParameters structures.Config,
		parameters []Parameter,
		requests map[int]Request,
		handlerFunc ...gin.HandlerFunc,
	)
	OPTIONS(
		swaggerParameters structures.Config,
		parameters []Parameter,
		requests map[int]Request,
		handlerFunc ...gin.HandlerFunc,
	)
	PATCH(
		swaggerParameters structures.Config,
		parameters []Parameter,
		requests map[int]Request,
		handlerFunc ...gin.HandlerFunc,
	)
	PUT(
		swaggerParameters structures.Config,
		parameters []Parameter,
		requests map[int]Request,
		handlerFunc ...gin.HandlerFunc,
	)

	generate() swaggerFileGenerator.PathSwagger
	getDefinitions() []parameters.SwaggParameter
	readRequests(requests map[int]Request) []swaggerFileGenerator.ResponseSwagger
}

type swaggPathWrapper struct {
	path        string
	tag         string
	requests    []swaggerFileGenerator.RequestSwagger
	definitions []parameters.SwaggParameter

	group *gin.RouterGroup
}

type Request struct {
	description string
	object      interface{}
}

func NewRequest(description string, object interface{}) Request {
	return Request{
		description: description,
		object:      object,
	}
}

func newSwaggPathWrapper(path, tag string, group *gin.RouterGroup) SwaggPathWrapper {
	return &swaggPathWrapper{
		path:        path,
		tag:         tag,
		requests:    []swaggerFileGenerator.RequestSwagger{},
		definitions: []parameters.SwaggParameter{},
		group:       group,
	}
}

func (s *swaggPathWrapper) GET(
	swaggerParameters structures.Config,
	parametersP []Parameter,
	requests map[int]Request,
	handlerFuncP ...gin.HandlerFunc,
) {
	s.group.GET(s.path, handlerFuncP...)

	s.addRequest(swaggerParameters, parametersP, requests, "get")
}

func (s *swaggPathWrapper) POST(
	swaggerParameters structures.Config,
	parametersP []Parameter,
	requests map[int]Request,
	handlerFuncP ...gin.HandlerFunc,
) {
	s.group.POST(s.path, handlerFuncP...)

	s.addRequest(swaggerParameters, parametersP, requests, "post")
}

func (s *swaggPathWrapper) DELETE(
	swaggerParameters structures.Config,
	parametersP []Parameter,
	requests map[int]Request,
	handlerFuncP ...gin.HandlerFunc,
) {
	s.group.DELETE(s.path, handlerFuncP...)

	s.addRequest(swaggerParameters, parametersP, requests, "delete")
}

func (s *swaggPathWrapper) HEAD(
	swaggerParameters structures.Config,
	parametersP []Parameter,
	requests map[int]Request,
	handlerFuncP ...gin.HandlerFunc,
) {
	s.group.HEAD(s.path, handlerFuncP...)

	s.addRequest(swaggerParameters, parametersP, requests, "head")
}

func (s *swaggPathWrapper) OPTIONS(
	swaggerParameters structures.Config,
	parametersP []Parameter,
	requests map[int]Request,
	handlerFuncP ...gin.HandlerFunc,
) {
	s.group.OPTIONS(s.path, handlerFuncP...)

	s.addRequest(swaggerParameters, parametersP, requests, "options")
}

func (s *swaggPathWrapper) PATCH(
	swaggerParameters structures.Config,
	parametersP []Parameter,
	requests map[int]Request,
	handlerFuncP ...gin.HandlerFunc,
) {
	s.group.PATCH(s.path, handlerFuncP...)

	s.addRequest(swaggerParameters, parametersP, requests, "patch")
}

func (s *swaggPathWrapper) PUT(
	swaggerParameters structures.Config,
	parametersP []Parameter,
	requests map[int]Request,
	handlerFuncP ...gin.HandlerFunc,
) {
	s.group.PUT(s.path, handlerFuncP...)

	s.addRequest(swaggerParameters, parametersP, requests, "put")
}

func (s *swaggPathWrapper) generate() swaggerFileGenerator.PathSwagger {
	res := swaggerFileGenerator.NewPathSwagger(s.path, s.requests)
	return res
}

func (s *swaggPathWrapper) getDefinitions() []parameters.SwaggParameter {
	return s.definitions
}

func (s *swaggPathWrapper) addRequest(
	strConfigs structures.Config,
	parametersP []Parameter,
	requests map[int]Request,
	reqType string,
) {
	var paramsSwagg []parameters.SwaggParameter
	for _, val := range parametersP {
		if val.GetSwagParameter().IsObject() {
			paramsSwagg = append(paramsSwagg, parameters.NewSchemaSwaggParameter(val.GetSwagParameter()))
			s.definitions = append(s.definitions, val.GetSwagParameter())
		} else {
			paramsSwagg = append(paramsSwagg, val.GetSwagParameter())
		}
	}

	var configs map[string]interface{}
	if strConfigs != nil {
		configs = strConfigs.ToMap()
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
	s.requests = append(s.requests, swaggerFileGenerator.NewRequestSwagg(
		configs,
		paramsSwagg,
		s.readRequests(requests),
	))
}

func (s *swaggPathWrapper) readRequests(requests map[int]Request) []swaggerFileGenerator.ResponseSwagger {
	var responses []swaggerFileGenerator.ResponseSwagger
	//TODO: CHECK USUAL PARAMETER (STRING, BOOL, ...)

	for key, val := range requests {
		if val.object != nil {
			elemTypeName := reflect.TypeOf(val.object).Name()
			if reflect.ValueOf(val.object).Kind() == reflect.Ptr {
				elemTypeName = reflect.ValueOf(val.object).Elem().Type().Name()
			}
			element := ReturnNonStructureObject(nil, val.object)
			respSwag := swaggerFileGenerator.NewResponseSwagg(key, val.description, elemTypeName, element)
			responses = append(responses, respSwag)
			if element == nil {
				s.definitions = append(s.definitions, ConvertObjectToSwaggParameter(nil, val.object, false))
			}
		} else {
			respSwag := swaggerFileGenerator.NewResponseSwagg(key, val.description, "", nil)
			responses = append(responses, respSwag)
		}
	}

	return responses
}
