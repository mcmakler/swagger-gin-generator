package wrapper

import (
	"github.com/gin-gonic/gin"
	"github.com/mcmakler/swagger-gin-generator/swaggerFileGenerator"
	"github.com/mcmakler/swagger-gin-generator/swaggerFileGenerator/parameters"
)

type SwaggerGroupWrapper interface {
	Use(middleware ...gin.HandlerFunc)
	Path(string) SwaggerPathWrapper
	generate() []swaggerFileGenerator.PathSwagger
	getDefinitions() []parameters.SwaggParameter
}

type swaggerGroupWrapper struct {
	path            string
	tag             string
	swaggerWrappers []SwaggerPathWrapper
	definitions     []parameters.SwaggParameter

	group *gin.RouterGroup
}

func (s *swaggerGroupWrapper) Use(middlware ...gin.HandlerFunc) {
	s.group.Use(middlware...)
}

func (s *swaggerGroupWrapper) Path(path string) SwaggerPathWrapper {
	res := newSwaggerPathWrapper(s.path+path, s.tag, s.group)
	s.swaggerWrappers = append(s.swaggerWrappers, res)
	return res
}

func newSwaggerGroupWrapper(path, tag string, group *gin.RouterGroup) SwaggerGroupWrapper {
	return &swaggerGroupWrapper{
		path:            path,
		tag:             tag,
		swaggerWrappers: []SwaggerPathWrapper{},
		definitions:     []parameters.SwaggParameter{},
		group:           group,
	}
}

func (s *swaggerGroupWrapper) generate() []swaggerFileGenerator.PathSwagger {
	var res []swaggerFileGenerator.PathSwagger
	for _, val := range s.swaggerWrappers {
		for _, def := range val.getDefinitions() {
			s.definitions = append(s.definitions, def)
		}
		res = append(res, val.generate())
	}
	return res
}

func (s *swaggerGroupWrapper) getDefinitions() []parameters.SwaggParameter {
	return s.definitions
}
