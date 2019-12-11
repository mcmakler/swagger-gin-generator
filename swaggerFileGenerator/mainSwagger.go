package swaggerFileGenerator

import (
	"errors"
	"github.com/mcmakler/swagger-gin-generator/swaggerFileGenerator/parameters"
	"strings"
)

const (
	swaggerString         = "swagger: '2.0'"
	infoString            = "\ninfo:"
	infoTitleString       = "\n  title: "
	infoDescriptionString = "\n  description: "
	infoVersionString     = "\n  version: '"

	securityDefinitionString = "\nsecurityDefinitions:"

	hostString = "\nhost: "
	basePathString = "\nbasePath: "
	pathsString    = "\npaths:"

	definitionsString = "\ndefinitions:"

	mainIndentString = "\n  "
)

var (
	errorEmptyTitle   = errors.New("ERROR_EMPTY_TITLE")
	errorEmptyVersion = errors.New("ERROR_EMPTY_VERSION")
)

type MainSwagger interface {
	ToString() (string, error)
}

type mainSwagger struct {
	config              map[string]interface{}
	securityDefinitions []SecurityDefinitionSwagger
	paths               []PathSwagger
	definitions         []parameters.SwaggParameter
}

var errorEmptyPaths = errors.New("EMPTY_PATHS")

func (m *mainSwagger) ToString() (string, error) {
	if m.paths == nil {
		return "", errorEmptyPaths
	}
	res := swaggerString + infoString
	if val, ok := m.config["title"]; ok {
		res += infoTitleString + val.(string)
	} else {
		return "", errorEmptyTitle
	}
	if val, ok := m.config["version"]; ok {
		res += infoVersionString + val.(string) + "'"
	} else {
		return "", errorEmptyVersion
	}
	if val, ok := m.config["description"]; ok {
		res += infoDescriptionString + val.(string)
	}
	if val, ok := m.config["host"]; ok {
		res += hostString + val.(string)
	}
	if val, ok := m.config["basePath"]; !ok {
		res += basePathString + "/"
	} else {
		res += basePathString + val.(string)
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
		for _, definition := range m.definitions {
			str, err := definition.ToString()
			if err != nil {
				return "", err
			}
			res += strings.Replace(str, "\n", mainIndentString, -1)
		}
	}
	return res, nil
}

func NewMainSwagger(config map[string]interface{}, securityDefinitions []SecurityDefinitionSwagger, paths []PathSwagger, definitions []parameters.SwaggParameter) MainSwagger {
	return &mainSwagger{
		config:              config,
		securityDefinitions: securityDefinitions,
		paths:               paths,
		definitions:         definitions,
	}
}
