package wrapper

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"strings"
	"swagger-gin-generator/swaggerFileGenerator"
	"swagger-gin-generator/swaggerFileGenerator/parameters"
)

type SwaggRouterWrapper interface {
	Group(path, tag string) SwaggGroupWrapper
	Generate(filepath string) error
}

type swaggWrapper struct {
	params      map[string]interface{}
	paths       []swaggerFileGenerator.PathSwagger
	definitions []parameters.SwaggParameter

	groups []SwaggGroupWrapper

	router *gin.Engine
}

func NewSwaggerRouterWrapper(params map[string]interface{}, r *gin.Engine) SwaggRouterWrapper {
	return &swaggWrapper{
		params:      params,
		paths:       []swaggerFileGenerator.PathSwagger{},
		definitions: []parameters.SwaggParameter{},
		groups:      []SwaggGroupWrapper{},
		router:      r,
	}
}

func (s *swaggWrapper) Group(path, tag string) SwaggGroupWrapper {
	group := s.router.Group(path)
	res := newSwaggGroupWrapper(path, tag, group)
	s.groups = append(s.groups, res)
	return res
}

func (s *swaggWrapper) Generate(filepath string) error {
	for _, val := range s.groups {
		for _, path := range val.Generate() {
			s.paths = append(s.paths, path)
		}
		for _, def := range val.getDefinitions() {
			s.definitions = append(s.definitions, def)
		}
	}
	s.definitions = sliceUniqMap(s.definitions)
	mainSwagg := swaggerFileGenerator.NewMainSwagg(
		s.params,
		s.paths,
		s.definitions)
	str, err := mainSwagg.ToString()
	if err != nil {
		return err
	}
	err = writeStringToFile(filepath, str)
	if err != nil {
		return err
	}
	return nil
}

func sliceUniqMap(s []parameters.SwaggParameter) []parameters.SwaggParameter {
	keys := make(map[string]bool)
	var list []parameters.SwaggParameter
	for _, entry := range s {
		//TODO: Maybe some better solution
		str, _ := entry.ToString(true)
		str = strings.Split(str, ":")[0]
		if _, value := keys[str]; !value {
			keys[str] = true
			list = append(list, entry)
		}
	}
	return list
}

func writeStringToFile(filepath, s string) error {
	fo, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer fo.Close()

	_, err = io.Copy(fo, strings.NewReader(s))
	if err != nil {
		return err
	}

	return nil
}
