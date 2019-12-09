package wrapper

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mcmakler/swagger-gin-generator/structures"
	"github.com/mcmakler/swagger-gin-generator/swaggerFileGenerator"
	"github.com/mcmakler/swagger-gin-generator/swaggerFileGenerator/parameters"
)

type SwaggerGroupWrapper interface {
	Use(middleware ...gin.HandlerFunc)
	Path(path string) SwaggerPathWrapper
	Group(path, tag string) SwaggerGroupWrapper
	generate() []swaggerFileGenerator.PathSwagger
	getDefinitions() []parameters.SwaggParameter
	GET(path string, config structures.Config, parameters []Parameter, requests map[int]Response, handlerFunc ...gin.HandlerFunc)
	POST(path string, config structures.Config, parameters []Parameter, requests map[int]Response, handlerFunc ...gin.HandlerFunc)
	DELETE(path string, config structures.Config, parameters []Parameter, requests map[int]Response, handlerFunc ...gin.HandlerFunc)
	HEAD(path string, config structures.Config, parameters []Parameter, requests map[int]Response, handlerFunc ...gin.HandlerFunc)
	OPTIONS(path string, config structures.Config, parameters []Parameter, requests map[int]Response, handlerFunc ...gin.HandlerFunc)
	PATCH(path string, config structures.Config, parameters []Parameter, requests map[int]Response, handlerFunc ...gin.HandlerFunc)
	PUT(path string, config structures.Config, parameters []Parameter, requests map[int]Response, handlerFunc ...gin.HandlerFunc)
}

type swaggerGroupWrapper struct {
	path        string
	tag         string
	paths       map[string]SwaggerPathWrapper
	definitions []parameters.SwaggParameter
	subgroups   []SwaggerGroupWrapper

	group *gin.RouterGroup
}

func (s *swaggerGroupWrapper) Use(middlware ...gin.HandlerFunc) {
	s.group.Use(middlware...)
}

func (s *swaggerGroupWrapper) Path(path string) SwaggerPathWrapper {
	res := newSwaggerPathWrapper(s.path + path, s.tag, s.group)
	s.paths[path] = res
	return res
}

func (s *swaggerGroupWrapper) Group(path, tag string) SwaggerGroupWrapper {
	group := s.group.Group(path)
	fmt.Println(s.group.BasePath())
	fmt.Println(s.group.Handlers)
	res := newSwaggerGroupWrapper(s.path+path, tag, group)
	s.subgroups = append(s.subgroups, res)
	return res
}

func newSwaggerGroupWrapper(path, tag string, group *gin.RouterGroup) SwaggerGroupWrapper {
	return &swaggerGroupWrapper{
		path:        path,
		tag:         tag,
		paths:       make(map[string]SwaggerPathWrapper),
		definitions: []parameters.SwaggParameter{},
		subgroups:   []SwaggerGroupWrapper{},
		group:       group,
	}
}

func (s *swaggerGroupWrapper) generate() []swaggerFileGenerator.PathSwagger {
	var res []swaggerFileGenerator.PathSwagger
	for _, val := range s.paths {
		for _, def := range val.getDefinitions() {
			s.definitions = append(s.definitions, def)
		}
		res = append(res, val.generate())
	}
	for _, val := range s.subgroups {
		for _, path := range val.generate() {
			res = append(res, path)
		}
		for _, def := range val.getDefinitions() {
			s.definitions = append(s.definitions, def)
		}
	}
	return res
}

func (s *swaggerGroupWrapper) getDefinitions() []parameters.SwaggParameter {
	return s.definitions
}

func (s *swaggerGroupWrapper) GET(path string, config structures.Config, parameters []Parameter, requests map[int]Response, handlerFunc ...gin.HandlerFunc) {
	if _, ok := s.paths[path]; !ok {
		s.Path(path)
	}
	s.paths[path].GET(config, parameters, requests, handlerFunc...)
}

func (s *swaggerGroupWrapper) POST(path string, config structures.Config, parameters []Parameter, requests map[int]Response, handlerFunc ...gin.HandlerFunc) {
	if _, ok := s.paths[path]; !ok {
		s.Path(path)
	}
	s.paths[path].POST(config, parameters, requests, handlerFunc...)
}

func (s *swaggerGroupWrapper) DELETE(path string, config structures.Config, parameters []Parameter, requests map[int]Response, handlerFunc ...gin.HandlerFunc) {
	if _, ok := s.paths[path]; !ok {
		s.Path(path)
	}
	s.paths[path].DELETE(config, parameters, requests, handlerFunc...)
}

func (s *swaggerGroupWrapper) HEAD(path string, config structures.Config, parameters []Parameter, requests map[int]Response, handlerFunc ...gin.HandlerFunc) {
	if _, ok := s.paths[path]; !ok {
		s.Path(path)
	}
	s.paths[path].HEAD(config, parameters, requests, handlerFunc...)
}

func (s *swaggerGroupWrapper) OPTIONS(path string, config structures.Config, parameters []Parameter, requests map[int]Response, handlerFunc ...gin.HandlerFunc) {
	if _, ok := s.paths[path]; !ok {
		s.Path(path)
	}
	s.paths[path].OPTIONS(config, parameters, requests, handlerFunc...)
}

func (s *swaggerGroupWrapper) PATCH(path string, config structures.Config, parameters []Parameter, requests map[int]Response, handlerFunc ...gin.HandlerFunc) {
	if _, ok := s.paths[path]; !ok {
		s.Path(path)
	}
	s.paths[path].PATCH(config, parameters, requests, handlerFunc...)
}

func (s *swaggerGroupWrapper) PUT(path string, config structures.Config, parameters []Parameter, requests map[int]Response, handlerFunc ...gin.HandlerFunc) {
	if _, ok := s.paths[path]; !ok {
		s.Path(path)
	}
	s.paths[path].PUT(config, parameters, requests, handlerFunc...)
}
