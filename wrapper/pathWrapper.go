package wrapper

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"swagger-gin-generator/swaggerFileGenerator"
	"swagger-gin-generator/swaggerFileGenerator/parameters"
	"swagger-gin-generator/utils"
)

type SwaggPathWrapper interface {
	Get(
		swaggerParameters map[string]interface{},
		parameters []utils.Parameter,
		requests map[int]Request,
		handlerFunc ...gin.HandlerFunc,
	)
	Post(
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

func NewSwaggPathWrapper(path, tag string, group *gin.RouterGroup) SwaggPathWrapper {
	return &swaggPathWrapper{
		path:        path,
		tag:         tag,
		requests:    []swaggerFileGenerator.RequestSwagg{},
		definitions: []parameters.SwaggParameter{},
		group:       group,
	}
}

func (s *swaggPathWrapper) Get(
	swaggerParameters map[string]interface{},
	parametersP []utils.Parameter,
	requests map[int]Request,
	handlerFuncP ...gin.HandlerFunc,
) {
	s.group.GET(s.path, handlerFuncP...)

	var responses []swaggerFileGenerator.ResponseSwagg

	for key, val := range requests {
		if val.object != nil {
			respSwag := swaggerFileGenerator.NewResponseSwagg(key, val.description, reflect.TypeOf(val.object).Name())
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

	if swaggerParameters == nil {
		swaggerParameters = map[string]interface{}{}
	}

	if s.tag != "" {
		if _, ok := swaggerParameters["tagsP"]; !ok {
			swaggerParameters["tagsP"] = []string{}
		}
		swaggerParameters["tagsP"] = append(swaggerParameters["tagsP"].([]string), s.tag)
	}

	swaggerParameters["typeRequest"] = "get"
	s.requests = append(s.requests, swaggerFileGenerator.NewRequestSwagg(
		swaggerParameters,
		paramsSwagg,
		responses,
	))
}

func (s *swaggPathWrapper) Post(
	swaggerParameters map[string]interface{},
	parametersP []utils.Parameter,
	requests map[int]Request,
	handlerFuncP ...gin.HandlerFunc,
) {
	s.group.POST(s.path, handlerFuncP...)

	var responses []swaggerFileGenerator.ResponseSwagg

	for key, val := range requests {
		if val.object != nil {
			respSwag := swaggerFileGenerator.NewResponseSwagg(key, val.description, reflect.TypeOf(val.object).Name())
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

	if swaggerParameters == nil {
		swaggerParameters = map[string]interface{}{}
	}

	if s.tag != "" {
		if _, ok := swaggerParameters["tagsP"]; !ok {
			swaggerParameters["tagsP"] = []string{}
		}
		swaggerParameters["tagsP"] = append(swaggerParameters["tagsP"].([]string), s.tag)
	}

	swaggerParameters["typeRequest"] = "post"
	s.requests = append(s.requests, swaggerFileGenerator.NewRequestSwagg(
		swaggerParameters,
		paramsSwagg,
		responses,
	))
}

func (s *swaggPathWrapper) generate() swaggerFileGenerator.PathSwagger {
	res := swaggerFileGenerator.NewPathSwagger(s.path, s.requests)
	return res
}

func (s *swaggPathWrapper) getDefinitions() []parameters.SwaggParameter {
	return s.definitions
}
