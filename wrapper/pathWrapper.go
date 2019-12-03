package wrapper

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"swagger-gin-generator/swaggerFileGenerator"
	"swagger-gin-generator/swaggerFileGenerator/parameters"
	"swagger-gin-generator/utils"
)

type SwaggPathWrapper interface {
	GET(
		swaggerParameters map[string]interface{},
		parameters []utils.Parameter,
		requests map[int]Request,
		handlerFunc ...gin.HandlerFunc,
	)
	POST(
		swaggerParameters map[string]interface{},
		parameters []utils.Parameter,
		requests map[int]Request,
		handlerFunc ...gin.HandlerFunc,
	)
	DELETE(
		swaggerParameters map[string]interface{},
		parameters []utils.Parameter,
		requests map[int]Request,
		handlerFunc ...gin.HandlerFunc,
	)
	HEAD(
		swaggerParameters map[string]interface{},
		parameters []utils.Parameter,
		requests map[int]Request,
		handlerFunc ...gin.HandlerFunc,
	)
	OPTIONS(
		swaggerParameters map[string]interface{},
		parameters []utils.Parameter,
		requests map[int]Request,
		handlerFunc ...gin.HandlerFunc,
	)
	PATCH(
		swaggerParameters map[string]interface{},
		parameters []utils.Parameter,
		requests map[int]Request,
		handlerFunc ...gin.HandlerFunc,
	)
	PUT(
		swaggerParameters map[string]interface{},
		parameters []utils.Parameter,
		requests map[int]Request,
		handlerFunc ...gin.HandlerFunc,
	)

	generate() swaggerFileGenerator.PathSwagger
	getDefinitions() []parameters.SwaggParameter
}

type swaggPathWrapper struct {
	path        string
	tag         string
	requests    []swaggerFileGenerator.RequestSwagg
	definitions []parameters.SwaggParameter

	group *gin.RouterGroup
}

type Request struct {
	description string
	object      interface{}
}

func newSwaggPathWrapper(path, tag string, group *gin.RouterGroup) SwaggPathWrapper {
	return &swaggPathWrapper{
		path:        path,
		tag:         tag,
		requests:    []swaggerFileGenerator.RequestSwagg{},
		definitions: []parameters.SwaggParameter{},
		group:       group,
	}
}

func (s *swaggPathWrapper) GET(
	swaggerParameters map[string]interface{},
	parametersP []utils.Parameter,
	requests map[int]Request,
	handlerFuncP ...gin.HandlerFunc,
) {
	s.group.GET(s.path, handlerFuncP...)

	s.addRequest(swaggerParameters, parametersP, requests, "get")
}

func (s *swaggPathWrapper) POST(
	swaggerParameters map[string]interface{},
	parametersP []utils.Parameter,
	requests map[int]Request,
	handlerFuncP ...gin.HandlerFunc,
) {
	s.group.POST(s.path, handlerFuncP...)

	s.addRequest(swaggerParameters, parametersP, requests, "post")
}

func (s *swaggPathWrapper) DELETE(
	swaggerParameters map[string]interface{},
	parametersP []utils.Parameter,
	requests map[int]Request,
	handlerFuncP ...gin.HandlerFunc,
) {
	s.group.DELETE(s.path, handlerFuncP...)

	s.addRequest(swaggerParameters, parametersP, requests, "delete")
}

func (s *swaggPathWrapper) HEAD(
	swaggerParameters map[string]interface{},
	parametersP []utils.Parameter,
	requests map[int]Request,
	handlerFuncP ...gin.HandlerFunc,
) {
	s.group.HEAD(s.path, handlerFuncP...)

	s.addRequest(swaggerParameters, parametersP, requests, "head")
}

func (s *swaggPathWrapper) OPTIONS(
	swaggerParameters map[string]interface{},
	parametersP []utils.Parameter,
	requests map[int]Request,
	handlerFuncP ...gin.HandlerFunc,
) {
	s.group.OPTIONS(s.path, handlerFuncP...)

	s.addRequest(swaggerParameters, parametersP, requests, "options")
}

func (s *swaggPathWrapper) PATCH(
	swaggerParameters map[string]interface{},
	parametersP []utils.Parameter,
	requests map[int]Request,
	handlerFuncP ...gin.HandlerFunc,
) {
	s.group.PATCH(s.path, handlerFuncP...)

	s.addRequest(swaggerParameters, parametersP, requests, "patch")
}

func (s *swaggPathWrapper) PUT(
	swaggerParameters map[string]interface{},
	parametersP []utils.Parameter,
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
	configs map[string]interface{},
	parametersP []utils.Parameter,
	requests map[int]Request,
	reqType string,
) {
	var responses []swaggerFileGenerator.ResponseSwagg

	for key, val := range requests {
		if val.object != nil {
			var respSwag swaggerFileGenerator.ResponseSwagg
			if reflect.ValueOf(val.object).Kind() == reflect.Ptr {
				respSwag = swaggerFileGenerator.NewResponseSwagg(key, val.description, reflect.ValueOf(val.object).Elem().Type().Name())
			} else {
				respSwag = swaggerFileGenerator.NewResponseSwagg(key, val.description, reflect.TypeOf(val.object).Name())
			}
			responses = append(responses, respSwag)
			s.definitions = append(s.definitions, utils.ConvertObjectToSwaggParameter(nil, val.object, false))
		} else {
			respSwag := swaggerFileGenerator.NewResponseSwagg(key, val.description, "")
			responses = append(responses, respSwag)
		}
	}

	var paramsSwagg []parameters.SwaggParameter
	for _, val := range parametersP {
		paramsSwagg = append(paramsSwagg, val.GetSwagParameter())
	}

	if configs == nil {
		configs = map[string]interface{}{}
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
		responses,
	))
}
