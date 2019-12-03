package wrapper

import (
	"github.com/gin-gonic/gin"
	"swagger-gin-generator/swaggerFileGenerator"
	"swagger-gin-generator/swaggerFileGenerator/parameters"
)

type SwaggGroupWrapper interface {
	Use(middleware ...gin.HandlerFunc)
	Path(string) SwaggPathWrapper
	generate() []swaggerFileGenerator.PathSwagger
	getDefinitions() []parameters.SwaggParameter
}

type swaggGroupWrapper struct {
	path          string
	tag           string
	swaggWrappers []SwaggPathWrapper
	definitions   []parameters.SwaggParameter

	group *gin.RouterGroup
}

func (s *swaggGroupWrapper) Use(middlware ...gin.HandlerFunc) {
	s.group.Use(middlware...)
}

func (s *swaggGroupWrapper) Path(path string) SwaggPathWrapper {
	res := newSwaggPathWrapper(s.path+path, s.tag, s.group)
	s.swaggWrappers = append(s.swaggWrappers, res)
	return res
}

func newSwaggGroupWrapper(path, tag string, group *gin.RouterGroup) SwaggGroupWrapper {
	return &swaggGroupWrapper{
		path:          path,
		tag:           tag,
		swaggWrappers: []SwaggPathWrapper{},
		definitions:   []parameters.SwaggParameter{},
		group:         group,
	}
}

func (s *swaggGroupWrapper) generate() []swaggerFileGenerator.PathSwagger {
	var res []swaggerFileGenerator.PathSwagger
	for _, val := range s.swaggWrappers {
		for _, def := range val.getDefinitions() {
			s.definitions = append(s.definitions, def)
		}
		res = append(res, val.generate())
	}
	return res
}

func (s *swaggGroupWrapper) getDefinitions() []parameters.SwaggParameter {
	return s.definitions
}
