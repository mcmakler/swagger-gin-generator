package wrapper

import (
	yaml "github.com/ghodss/yaml"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/mcmakler/swagger-gin-generator/structures"
	"github.com/mcmakler/swagger-gin-generator/swaggerFileGenerator"
	"github.com/mcmakler/swagger-gin-generator/swaggerFileGenerator/parameters"
	"html/template"
	"io"
	"os"
	"strings"
)

//TODO: consts for all parameters

//TODO: params in comments

const (
	InBody     = "body"
	InPath     = "path"
	InQuery    = "query"
	InHeader   = "header"
	InFromData = "fromData"

	SecurityBearer = "Bearer"

	TypesJson = "json"
	TypesBson = "bson"

	filenameString     = "swagger.yaml"
	filenameStringJson = "swagger.json"

	templateUrl = "./swagger.json"
)

type SwaggerRouterWrapper interface {
	Group(path, tag string) SwaggerGroupWrapper
	Use(middleware ...gin.HandlerFunc)
	GenerateFiles(filepath string) error
	GenerateBasePath(pathUrl string) error
	generate() (string, error)
	NewBasicSecurityDefinition(title string)
	NewApiKeySecurityDefinition(title, name string, inHeader bool)
	NewOauth2ImplicitSecurityDefinition(title, authorizationUrl string)
	NewOauth2PasswordSecurityDefinition(title, tokenURL string)
	NewOauth2ApplicationSecurityDefinition(title, tokenURL string)
	NewOauth2AccessCodeSecurityDefinition(title, authorizationUrl, tokenURL string)
	GET(path string, config structures.Config, parameters []Parameter, requests map[int]Response, handlerFunc ...gin.HandlerFunc)
	POST(path string, config structures.Config, parameters []Parameter, requests map[int]Response, handlerFunc ...gin.HandlerFunc)
	DELETE(path string, config structures.Config, parameters []Parameter, requests map[int]Response, handlerFunc ...gin.HandlerFunc)
	HEAD(path string, config structures.Config, parameters []Parameter, requests map[int]Response, handlerFunc ...gin.HandlerFunc)
	OPTIONS(path string, config structures.Config, parameters []Parameter, requests map[int]Response, handlerFunc ...gin.HandlerFunc)
	PATCH(path string, config structures.Config, parameters []Parameter, requests map[int]Response, handlerFunc ...gin.HandlerFunc)
	PUT(path string, config structures.Config, parameters []Parameter, requests map[int]Response, handlerFunc ...gin.HandlerFunc)
}

type swaggerWrapper struct {
	configs map[string]interface{}

	security    []swaggerFileGenerator.SecurityDefinitionSwagger
	paths       []swaggerFileGenerator.PathSwagger
	definitions []parameters.SwaggParameter

	mainGroup SwaggerGroupWrapper
}

func NewSwaggerRouterWrapper(config structures.Config, r *gin.Engine) SwaggerRouterWrapper {
	return &swaggerWrapper{
		configs:     config.ToMap(),
		security:    []swaggerFileGenerator.SecurityDefinitionSwagger{},
		paths:       []swaggerFileGenerator.PathSwagger{},
		definitions: []parameters.SwaggParameter{},
		mainGroup:   newSwaggerGroupWrapper("", "", r.Group("")),
	}
}

func (s *swaggerWrapper) Use(middlware ...gin.HandlerFunc) {
	s.mainGroup.Use(middlware...)
}

func (s *swaggerWrapper) Group(path, tag string) SwaggerGroupWrapper {
	return s.mainGroup.Group(path, tag)
}

func (s *swaggerWrapper) GenerateFiles(filepath string) error {
	str, err := s.generate()
	if err != nil {
		return err
	}
	err = writeStringToFile(filepath+filenameString, str)
	if err != nil {
		return err
	}
	jsonBytes, err := yaml.YAMLToJSON([]byte(str))
	jsonStr := string(jsonBytes)
	if err != nil {
		return err
	}
	err = writeStringToFile(filepath+filenameStringJson, jsonStr)
	if err != nil {
		return err
	}
	return nil
}

func (s *swaggerWrapper) GenerateBasePath(pathUrl string) error {
	str, err := s.generate()
	if err != nil {
		return err
	}
	jsonBytes, err := yaml.YAMLToJSON([]byte(str))
	jsonStr := string(jsonBytes)
	if err != nil {
		return err
	}
	templateJson := template.Must(template.New("json").Parse(jsonStr))
	s.mainGroup.GET(pathUrl[:strings.LastIndex(pathUrl, "/")] + filenameStringJson, nil, nil,
		map[int]Response{
			200: NewResponse("ok", nil),
		},
		func(c *gin.Context) {
			c.Render(200, render.HTML{
				Template: templateJson,
			})
		})
	templateDoc := template.Must(template.New("body").Parse(indexTemplate))
	s.mainGroup.GET(pathUrl, nil, nil,
		map[int]Response{
			200: NewResponse("ok", nil),
		},
		func(c *gin.Context) {
			c.Render(200, render.HTML{
				Template: templateDoc,
				Data: map[string]string{
					"templateUrl": templateUrl,
				},
			})
		})
	return nil
}

func (s *swaggerWrapper) generate() (string, error) {
	for _, path := range s.mainGroup.generate() {
		s.paths = append(s.paths, path)
	}
	for _, def := range s.mainGroup.getDefinitions() {
		s.definitions = append(s.definitions, def)
	}
	s.definitions = sliceUniqMap(s.definitions)
	mainSwagg := swaggerFileGenerator.NewMainSwagger(
		s.configs,
		s.security,
		s.paths,
		s.definitions)
	return mainSwagg.ToString()
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

func (s *swaggerWrapper) GET(path string, config structures.Config, parameters []Parameter, requests map[int]Response, handlerFunc ...gin.HandlerFunc) {
	s.mainGroup.GET(path, config, parameters, requests, handlerFunc...)
}

func (s *swaggerWrapper) POST(path string, config structures.Config, parameters []Parameter, requests map[int]Response, handlerFunc ...gin.HandlerFunc) {
	s.mainGroup.POST(path, config, parameters, requests, handlerFunc...)
}

func (s *swaggerWrapper) DELETE(path string, config structures.Config, parameters []Parameter, requests map[int]Response, handlerFunc ...gin.HandlerFunc) {
	s.mainGroup.DELETE(path, config, parameters, requests, handlerFunc...)
}

func (s *swaggerWrapper) HEAD(path string, config structures.Config, parameters []Parameter, requests map[int]Response, handlerFunc ...gin.HandlerFunc) {
	s.mainGroup.HEAD(path, config, parameters, requests, handlerFunc...)
}

func (s *swaggerWrapper) OPTIONS(path string, config structures.Config, parameters []Parameter, requests map[int]Response, handlerFunc ...gin.HandlerFunc) {
	s.mainGroup.OPTIONS(path, config, parameters, requests, handlerFunc...)
}

func (s *swaggerWrapper) PATCH(path string, config structures.Config, parameters []Parameter, requests map[int]Response, handlerFunc ...gin.HandlerFunc) {
	s.mainGroup.PATCH(path, config, parameters, requests, handlerFunc...)
}

func (s *swaggerWrapper) PUT(path string, config structures.Config, parameters []Parameter, requests map[int]Response, handlerFunc ...gin.HandlerFunc) {
	s.mainGroup.PUT(path, config, parameters, requests, handlerFunc...)
}

const indexTemplate = `<!-- HTML for static distribution bundle build -->
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Swagger UI</title>
  <link href="https://fonts.googleapis.com/css?family=Open+Sans:400,700|Source+Code+Pro:300,600|Titillium+Web:400,600,700" rel="stylesheet">
  <link rel="stylesheet" type="text/css" href="./swagger-ui.css" >
  <link rel="icon" type="image/png" href="./favicon-32x32.png" sizes="32x32" />
  <link rel="icon" type="image/png" href="./favicon-16x16.png" sizes="16x16" />
  <style>
    html
    {
        box-sizing: border-box;
        overflow: -moz-scrollbars-vertical;
        overflow-y: scroll;
    }
    *,
    *:before,
    *:after
    {
        box-sizing: inherit;
    }

    body {
      margin:0;
      background: #fafafa;
    }
  </style>
</head>

<body>

<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" style="position:absolute;width:0;height:0">
  <defs>
    <symbol viewBox="0 0 20 20" id="unlocked">
          <path d="M15.8 8H14V5.6C14 2.703 12.665 1 10 1 7.334 1 6 2.703 6 5.6V6h2v-.801C8 3.754 8.797 3 10 3c1.203 0 2 .754 2 2.199V8H4c-.553 0-1 .646-1 1.199V17c0 .549.428 1.139.951 1.307l1.197.387C5.672 18.861 6.55 19 7.1 19h5.8c.549 0 1.428-.139 1.951-.307l1.196-.387c.524-.167.953-.757.953-1.306V9.199C17 8.646 16.352 8 15.8 8z"></path>
    </symbol>

    <symbol viewBox="0 0 20 20" id="locked">
      <path d="M15.8 8H14V5.6C14 2.703 12.665 1 10 1 7.334 1 6 2.703 6 5.6V8H4c-.553 0-1 .646-1 1.199V17c0 .549.428 1.139.951 1.307l1.197.387C5.672 18.861 6.55 19 7.1 19h5.8c.549 0 1.428-.139 1.951-.307l1.196-.387c.524-.167.953-.757.953-1.306V9.199C17 8.646 16.352 8 15.8 8zM12 8H8V5.199C8 3.754 8.797 3 10 3c1.203 0 2 .754 2 2.199V8z"/>
    </symbol>

    <symbol viewBox="0 0 20 20" id="close">
      <path d="M14.348 14.849c-.469.469-1.229.469-1.697 0L10 11.819l-2.651 3.029c-.469.469-1.229.469-1.697 0-.469-.469-.469-1.229 0-1.697l2.758-3.15-2.759-3.152c-.469-.469-.469-1.228 0-1.697.469-.469 1.228-.469 1.697 0L10 8.183l2.651-3.031c.469-.469 1.228-.469 1.697 0 .469.469.469 1.229 0 1.697l-2.758 3.152 2.758 3.15c.469.469.469 1.229 0 1.698z"/>
    </symbol>

    <symbol viewBox="0 0 20 20" id="large-arrow">
      <path d="M13.25 10L6.109 2.58c-.268-.27-.268-.707 0-.979.268-.27.701-.27.969 0l7.83 7.908c.268.271.268.709 0 .979l-7.83 7.908c-.268.271-.701.27-.969 0-.268-.269-.268-.707 0-.979L13.25 10z"/>
    </symbol>

    <symbol viewBox="0 0 20 20" id="large-arrow-down">
      <path d="M17.418 6.109c.272-.268.709-.268.979 0s.271.701 0 .969l-7.908 7.83c-.27.268-.707.268-.979 0l-7.908-7.83c-.27-.268-.27-.701 0-.969.271-.268.709-.268.979 0L10 13.25l7.418-7.141z"/>
    </symbol>


    <symbol viewBox="0 0 24 24" id="jump-to">
      <path d="M19 7v4H5.83l3.58-3.59L8 6l-6 6 6 6 1.41-1.41L5.83 13H21V7z"/>
    </symbol>

    <symbol viewBox="0 0 24 24" id="expand">
      <path d="M10 18h4v-2h-4v2zM3 6v2h18V6H3zm3 7h12v-2H6v2z"/>
    </symbol>

  </defs>
</svg>

<div id="swagger-ui"></div>

<script src="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/3.23.11/swagger-ui-bundle.js"> </script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/3.23.11/swagger-ui-standalone-preset.js"> </script>
<script>
window.onload = function() {
  // Build a system
  const ui = SwaggerUIBundle({
    url: "{{.templateUrl}}",
    dom_id: '#swagger-ui',
    validatorUrl: null,
    presets: [
      SwaggerUIBundle.presets.apis,
      SwaggerUIStandalonePreset
    ],
    plugins: [
      SwaggerUIBundle.plugins.DownloadUrl
    ],
    layout: "StandaloneLayout"
  })

  window.ui = ui
}
</script>
</body>

</html>
`
