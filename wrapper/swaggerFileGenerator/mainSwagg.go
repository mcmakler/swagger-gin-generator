package swaggerFileGenerator

import (
	"errors"
	"strings"
	"github.com/mcmakler/swagger-gin-generator/wrapper/swaggerFileGenerator/parameters"
)

const (
	swaggerString         = "swagger: '2.0'"
	infoString            = "\ninfo:"
	infoTitleString       = "\n  title: "
	infoDescriptionString = "\n  description: "
	infoVersionString     = "\n  version: "

	securityDefinitionString = "\nsecurityDefinitions:"

	basePathString = "\nbasePath: "
	pathsString    = "\npaths:"

	definitionsString = "\ndefinitions:"

	mainIndentString = "\n  "
)

var (
	errorEmptyTitle   = errors.New("ERROR_EMPTY_TITLE")
	errorEmptyVersion = errors.New("ERROR_EMPTY_VERSION")
)

type MainSwagg interface {
	ToString() (string, error)
}

type mainSwagg struct {
	configs             map[string]interface{}
	securityDefinitions []SecurityDefinitionSwagg
	paths               []PathSwagger
	definitions         []parameters.SwaggParameter
	//TODO: security
}

var errorEmptyPaths = errors.New("EMPTY_PATHS")

func (m *mainSwagg) ToString() (string, error) {
	if m.paths == nil {
		return "", errorEmptyPaths
	}
	res := swaggerString + infoString
	if val, ok := m.configs["title"]; ok {
		res += infoTitleString + val.(string)
	} else {
		return "", errorEmptyTitle
	}
	if val, ok := m.configs["version"]; ok {
		res += infoVersionString + val.(string)
	} else {
		return "", errorEmptyVersion
	}
	if val, ok := m.configs["description"]; ok {
		res += infoDescriptionString + val.(string)
	}

	if m.securityDefinitions != nil && len(m.securityDefinitions) > 0 {
		res += securityDefinitionString
		for _, val := range m.securityDefinitions {
			str, err := val.ToString()
			if err != nil {
				return "", err
			}
			res += strings.Replace(str, "\n", mainIndentString, -1)
		}
	}

	if val, ok := m.configs["basePath"]; !ok {
		res += basePathString + "/"
	} else {
		res += basePathString + val.(string)
	}
	res += pathsString
	for _, path := range m.paths {
		str, err := path.ToString()
		if err != nil {
			return "", err
		}
		res += strings.Replace(str, "\n", mainIndentString, -1)
	}

	if m.definitions != nil {
		res += definitionsString
		for _, def := range m.definitions {
			str, err := def.ToString(true)
			if err != nil {
				return "", err
			}
			res += strings.Replace(str, "\n", mainIndentString, -1)
		}
	}
	return res, nil
}

func NewMainSwagg(params map[string]interface{}, securityDefs []SecurityDefinitionSwagg, paths []PathSwagger, def []parameters.SwaggParameter) MainSwagg {
	return &mainSwagg{
		configs:             params,
		securityDefinitions: securityDefs,
		paths:               paths,
		definitions:         def,
	}
}
