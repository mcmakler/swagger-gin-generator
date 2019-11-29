package wrapper

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"swagger-gin-generator/swaggerFileGenerator"
	"swagger-gin-generator/swaggerFileGenerator/parameters"
	"swagger-gin-generator/utils"
)

const (
	schemaPrefix = "#/definitions/"
)

type SwaggPathWrapper interface {
	Get(
		description string,
		consumes []string,
		produces []string,
		tags []string,
		summary string,
		parameters []map[string]string, //TODO: think
		requests map[int]Request,
		handlerFunc ...gin.HandlerFunc,
	)
	Post(
		description string,
		consumes []string,
		produces []string,
		tags []string,
		summary string,
		parameters []map[string]string, //TODO: think
		requests map[int]Request,
		handlerFunc ...gin.HandlerFunc,
	)
	Generate() *swaggerFileGenerator.PathSwagger
	Definitions() []parameters.SwaggParameter
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
	parametersP []map[string]string, //TODO: think
	responsesP map[int]Request,
	handlerFuncP ...gin.HandlerFunc,
) {
	s.group.GET(s.path, handlerFuncP...)

	var responses []swaggerFileGenerator.ResponseSwagg

	for key, val := range responsesP {
		respSwag := swaggerFileGenerator.NewResponseSwagg(key, val.definition, schemaPrefix+"/"+reflect.ValueOf(&val.object).Type().Name())
		responses = append(responses, respSwag)
		s.definitions = append(s.definitions, utils.ConvertObjectToSwaggParameter(nil, val.object))
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
		nil, //TODO: set parameters
		responses,
	))
}

func (s *swaggPathWrapper) Post(
	descriptionP string,
	consumesP []string,
	producesP []string,
	tagsP []string,
	summaryP string,
	parametersP []map[string]string, //TODO: think
	responsesP map[int]Request,
	handlerFuncP ...gin.HandlerFunc,
) {
	s.group.GET(s.path, handlerFuncP...)

	var responses []swaggerFileGenerator.ResponseSwagg

	for key, val := range responsesP {
		respSwag := swaggerFileGenerator.NewResponseSwagg(key, val.definition, schemaPrefix+"/"+reflect.ValueOf(&val.object).Type().Name())
		responses = append(responses, respSwag)
		s.definitions = append(s.definitions, utils.ConvertObjectToSwaggParameter(nil, val.object))
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
		nil, //TODO: set parameters
		responses,
	))
}

func (s *swaggPathWrapper) Generate() *swaggerFileGenerator.PathSwagger {
	res := swaggerFileGenerator.NewPathSwagger(s.path, s.requests)
	return &res
}

func (s *swaggPathWrapper) Definitions() []parameters.SwaggParameter {
	return s.definitions
}
