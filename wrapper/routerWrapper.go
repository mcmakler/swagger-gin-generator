package wrapper

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"strings"
	"swagger-gin-generator/wrapper/swaggerFileGenerator"
	"swagger-gin-generator/wrapper/swaggerFileGenerator/parameters"
)

//TODO: caps-free params
//TODO: group subgroups
//TODO: group get, post, ...

const (
	filenameString = "swagger.yaml"
)

type SwaggRouterWrapper interface {
	Group(path, tag string) SwaggGroupWrapper
	Use(middleware ...gin.HandlerFunc)
	Generate(filepath string) error
	NewBasicSecurityDefinition(title string)
	NewApiKeySecurityDefinition(title, name string, inHeader bool)
	NewOauth2ImplicitSecurityDefinition(title, authorizationUrl string)
	NewOauth2PasswordSecurityDefinition(title, tokenURL string)
	NewOauth2ApplicationSecurityDefinition(title, tokenURL string)
	NewOauth2AccessCodeSecurityDefinition(title, authorizationUrl, tokenURL string)
}

type swaggWrapper struct {
	configs map[string]interface{}

	security    []swaggerFileGenerator.SecurityDefinitionSwagg
	paths       []swaggerFileGenerator.PathSwagger
	definitions []parameters.SwaggParameter

	groups []SwaggGroupWrapper

	router *gin.Engine
}

func NewSwaggerRouterWrapper(configs map[string]interface{}, r *gin.Engine) SwaggRouterWrapper {
	return &swaggWrapper{
		configs:     configs,
		security:    []swaggerFileGenerator.SecurityDefinitionSwagg{},
		paths:       []swaggerFileGenerator.PathSwagger{},
		definitions: []parameters.SwaggParameter{},
		groups:      []SwaggGroupWrapper{},
		router:      r,
	}
}

func (s *swaggWrapper) Use(middlware ...gin.HandlerFunc) {
	s.router.Use(middlware...)
}

func (s *swaggWrapper) Group(path, tag string) SwaggGroupWrapper {
	group := s.router.Group(path)
	res := newSwaggGroupWrapper(path, tag, group)
	s.groups = append(s.groups, res)
	return res
}

func (s *swaggWrapper) Generate(filepath string) error {
	for _, val := range s.groups {
		for _, path := range val.generate() {
			s.paths = append(s.paths, path)
		}
		for _, def := range val.getDefinitions() {
			s.definitions = append(s.definitions, def)
		}
	}
	s.definitions = sliceUniqMap(s.definitions)
	mainSwagg := swaggerFileGenerator.NewMainSwagg(
		s.configs,
		s.security, //TODO
		s.paths,
		s.definitions)
	str, err := mainSwagg.ToString()
	if err != nil {
		return err
	}
	err = writeStringToFile(filepath+filenameString, str)
	if err != nil {
		return err
	}
	return nil
}

func (s *swaggWrapper) NewBasicSecurityDefinition(title string) {
	s.security = append(s.security, swaggerFileGenerator.NewBasicSecurityDefinition(title))
}

func (s *swaggWrapper) NewApiKeySecurityDefinition(title, name string, inHeader bool) {
	s.security = append(s.security, swaggerFileGenerator.NewApiKeySecurityDefinition(title, name, inHeader))
}

func (s *swaggWrapper) NewOauth2ImplicitSecurityDefinition(title, authorizationUrl string) {
	s.security = append(s.security, swaggerFileGenerator.NewOauth2ImplicitSecurityDefinition(title, authorizationUrl))
}

func (s *swaggWrapper) NewOauth2PasswordSecurityDefinition(title, tokenURL string) {
	s.security = append(s.security, swaggerFileGenerator.NewOauth2PasswordSecurityDefinition(title, tokenURL))
}

func (s *swaggWrapper) NewOauth2ApplicationSecurityDefinition(title, tokenURL string) {
	s.security = append(s.security, swaggerFileGenerator.NewOauth2ApplicationSecurityDefinition(title, tokenURL))
}

func (s *swaggWrapper) NewOauth2AccessCodeSecurityDefinition(title, authorizationUrl, tokenURL string) {
	s.security = append(s.security, swaggerFileGenerator.NewOauth2AccessCodeSecurityDefinition(title, authorizationUrl, tokenURL))
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
