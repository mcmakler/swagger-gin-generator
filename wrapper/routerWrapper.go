package wrapper

import (
	"SwaggerGin/swaggerFileGenerator"
	"SwaggerGin/swaggerFileGenerator/parameters"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"strings"
)

type SwaggRouterWrapper interface {
	Group(path, tag string) SwaggGroupWrapper
	Generate(filepath string) error
}

type swaggWrapper struct {
	title       string
	description string
	version     string
	bPath       string
	paths       []swaggerFileGenerator.PathSwagger
	definitions []parameters.SwaggParameter

	groups []SwaggGroupWrapper

	router *gin.Engine
}

func NewSwaggerRouterWrapper(title, description, version, bPath string, r *gin.Engine) SwaggRouterWrapper {
	return &swaggWrapper{
		title:       title,
		description: description,
		version:     version,
		bPath:       bPath,
		paths:       []swaggerFileGenerator.PathSwagger{},
		definitions: []parameters.SwaggParameter{},
		groups:      []SwaggGroupWrapper{},
		router:      r,
	}
}

func (s *swaggWrapper) Group(path, tag string) SwaggGroupWrapper {
	group := s.router.Group(path)
	return NewSwaggGroupWrapper(path, tag, group)
}

func (s *swaggWrapper) Generate(filepath string) error {
	for _, val := range s.groups {
		for _, path := range val.Generate() {
			s.paths = append(s.paths, path)
		}
		for _, def := range val.Definitions() {
			s.definitions = append(s.definitions, def)
		}
	}
	s.definitions = sliceUniqMap(s.definitions)
	mainSwagg := swaggerFileGenerator.NewMainSwagg(
		s.title,
		s.description,
		s.version,
		s.bPath,
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
	keys := make(map[parameters.SwaggParameter]bool)
	var list []parameters.SwaggParameter
	for _, entry := range s {
		if _, value := keys[entry]; !value {
			keys[entry] = true
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
