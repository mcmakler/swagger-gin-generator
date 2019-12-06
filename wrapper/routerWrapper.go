package wrapper

import (
	"github.com/gin-gonic/gin"
	"github.com/mcmakler/swagger-gin-generator/structures"
	"github.com/mcmakler/swagger-gin-generator/swaggerFileGenerator"
	"github.com/mcmakler/swagger-gin-generator/swaggerFileGenerator/parameters"
	"io"
	"os"
	"strings"
)

//TODO: bearer in consts;  json, bson
//TODO: consts for all parameters
//TODO: group subgroups
//TODO: group get, post, ...
//TODO: params: check case when []string is only one element/ []string{"el"} = "el"

//TODO: params in comments
//TODO: required in object

const (
	InBody     = "body"
	InPath     = "path"
	InQuery    = "query"
	InHeader   = "header"
	InFromData = "fromData"

	filenameString = "swagger.yaml"
)

type SwaggerRouterWrapper interface {
	Group(path, tag string) SwaggerGroupWrapper
	Use(middleware ...gin.HandlerFunc)
	Generate(filepath string) error
	NewBasicSecurityDefinition(title string)
	NewApiKeySecurityDefinition(title, name string, inHeader bool)
	NewOauth2ImplicitSecurityDefinition(title, authorizationUrl string)
	NewOauth2PasswordSecurityDefinition(title, tokenURL string)
	NewOauth2ApplicationSecurityDefinition(title, tokenURL string)
	NewOauth2AccessCodeSecurityDefinition(title, authorizationUrl, tokenURL string)
}

type swaggerWrapper struct {
	configs map[string]interface{}

	security    []swaggerFileGenerator.SecurityDefinitionSwagger
	paths       []swaggerFileGenerator.PathSwagger
	definitions []parameters.SwaggParameter

	groups []SwaggerGroupWrapper

	router *gin.Engine
}

func NewSwaggerRouterWrapper(config structures.Config, r *gin.Engine) SwaggerRouterWrapper {
	return &swaggerWrapper{
		configs:     config.ToMap(),
		security:    []swaggerFileGenerator.SecurityDefinitionSwagger{},
		paths:       []swaggerFileGenerator.PathSwagger{},
		definitions: []parameters.SwaggParameter{},
		groups:      []SwaggerGroupWrapper{},
		router:      r,
	}
}

func (s *swaggerWrapper) Use(middlware ...gin.HandlerFunc) {
	s.router.Use(middlware...)
}

func (s *swaggerWrapper) Group(path, tag string) SwaggerGroupWrapper {
	group := s.router.Group(path)
	res := newSwaggerGroupWrapper(path, tag, group)
	s.groups = append(s.groups, res)
	return res
}

func (s *swaggerWrapper) Generate(filepath string) error {
	for _, val := range s.groups {
		for _, path := range val.generate() {
			s.paths = append(s.paths, path)
		}
		for _, def := range val.getDefinitions() {
			s.definitions = append(s.definitions, def)
		}
	}
	s.definitions = sliceUniqMap(s.definitions)
	mainSwagg := swaggerFileGenerator.NewMainSwagger(
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

func (s *swaggerWrapper) NewBasicSecurityDefinition(title string) {
	s.security = append(s.security, swaggerFileGenerator.NewBasicSecurityDefinition(title))
}

func (s *swaggerWrapper) NewApiKeySecurityDefinition(title, name string, inHeader bool) {
	s.security = append(s.security, swaggerFileGenerator.NewApiKeySecurityDefinition(title, name, inHeader))
}

func (s *swaggerWrapper) NewOauth2ImplicitSecurityDefinition(title, authorizationUrl string) {
	s.security = append(s.security, swaggerFileGenerator.NewOauth2ImplicitSecurityDefinition(title, authorizationUrl))
}

func (s *swaggerWrapper) NewOauth2PasswordSecurityDefinition(title, tokenURL string) {
	s.security = append(s.security, swaggerFileGenerator.NewOauth2PasswordSecurityDefinition(title, tokenURL))
}

func (s *swaggerWrapper) NewOauth2ApplicationSecurityDefinition(title, tokenURL string) {
	s.security = append(s.security, swaggerFileGenerator.NewOauth2ApplicationSecurityDefinition(title, tokenURL))
}

func (s *swaggerWrapper) NewOauth2AccessCodeSecurityDefinition(title, authorizationUrl, tokenURL string) {
	s.security = append(s.security, swaggerFileGenerator.NewOauth2AccessCodeSecurityDefinition(title, authorizationUrl, tokenURL))
}

func sliceUniqMap(s []parameters.SwaggParameter) []parameters.SwaggParameter {
	keys := make(map[string]bool)
	var list []parameters.SwaggParameter
	for _, entry := range s {
		//TODO: Maybe some better solution
		str, _ := entry.ToString()
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
