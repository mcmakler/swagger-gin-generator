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
		description string,
		consumes []string,
		produces []string,
		tags []string,
		summary string,
		parameters []utils.Parameter, //TODO: think
		requests map[int]Request,
		handlerFunc ...gin.HandlerFunc,
	)
	Post(
		description string,
		consumes []string,
		produces []string,
		tags []string,
		summary string,
		parameters []utils.Parameter, //TODO: think
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
	definition string
	object     interface{}
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
	descriptionP string,
	consumesP []string,
	producesP []string,
	tagsP []string,
	summaryP string,
	parametersP []utils.Parameter,
	requests map[int]Request,
	handlerFuncP ...gin.HandlerFunc,
) {
	s.group.GET(s.path, handlerFuncP...)

	var responses []swaggerFileGenerator.ResponseSwagg

	for key, val := range requests {
		respSwag := swaggerFileGenerator.NewResponseSwagg(key, val.definition, reflect.TypeOf(val.object).Name())
		responses = append(responses, respSwag)
		s.definitions = append(s.definitions, utils.ConvertObjectToSwaggParameter(nil, val.object))
	}

	var paramsSwagg []parameters.SwaggParameter
	for _, val := range parametersP {
		paramsSwagg = append(paramsSwagg, val.GetSwagParameter())
	}

	if s.tag != "" {
		if tagsP == nil {
			tagsP = []string{}
		}
		tagsP = append(tagsP, s.tag)
	}

	s.requests = append(s.requests, swaggerFileGenerator.NewRequestSwagg(
		map[string]interface{}{
			"typeRequest": "get",
			"description": descriptionP,
			"summary":     summaryP,
			"consumes":    consumesP,
			"produces":    producesP,
			"tags":        tagsP,
		},
		paramsSwagg,
		responses,
	))
}

func (s *swaggPathWrapper) Post(
	descriptionP string,
	consumesP []string,
	producesP []string,
	tagsP []string,
	summaryP string,
	parametersP []utils.Parameter,
	requests map[int]Request,
	handlerFuncP ...gin.HandlerFunc,
) {
	s.group.POST(s.path, handlerFuncP...)

	var responses []swaggerFileGenerator.ResponseSwagg

	for key, val := range requests {
		respSwag := swaggerFileGenerator.NewResponseSwagg(key, val.definition, reflect.TypeOf(val.object).Name())
		responses = append(responses, respSwag)
		s.definitions = append(s.definitions, utils.ConvertObjectToSwaggParameter(nil, val.object))
	}

	var paramsSwagg []parameters.SwaggParameter
	for _, val := range parametersP {
		paramsSwagg = append(paramsSwagg, val.GetSwagParameter())
	}

	if s.tag != "" {
		if tagsP == nil {
			tagsP = []string{}
		}
		tagsP = append(tagsP, s.tag)
	}

	s.requests = append(s.requests, swaggerFileGenerator.NewRequestSwagg(
		map[string]interface{}{
			"typeRequest": "post",
			"description": descriptionP,
			"summary":     summaryP,
			"consumes":    consumesP,
			"produces":    producesP,
			"tags":        tagsP,
		},
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
