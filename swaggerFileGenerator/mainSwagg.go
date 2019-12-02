package swaggerFileGenerator

import (
	"errors"
	"strings"
	"swagger-gin-generator/swaggerFileGenerator/parameters"
)

const (
	swaggerString         = "swagger: '2.0'"
	infoString            = "\ninfo:"
	infoTitleString       = "\n  title: "
	infoDescriptionString = "\n  description: "
	infoVersionString     = "\n  version: "
	basePathString        = "\nbasePath: "
	pathsString           = "\npaths:"
	definitionsString     = "\ndefinitions:"

	mainIndentString = "\n  "
)

type MainSwagg interface {
	ToString() (string, error)
}

type mainSwagg struct {
	params map[string]interface{}

	paths       []PathSwagger
	definitions []parameters.SwaggParameter
	//TODO: security
}

var errorEmptyPaths = errors.New("EMPTY_PATHS")

func (m *mainSwagg) ToString() (string, error) {
	if m.paths == nil {
		return "", errorEmptyPaths
	}
	res := swaggerString
	infoExists := false
	if val, ok := m.params["description"]; ok {
		res += infoString
		res += infoDescriptionString + val.(string)
		infoExists = true
	}
	if val, ok := m.params["title"]; ok {
		if !infoExists {
			res += infoString
			infoExists = true
		}
		res += infoTitleString + val.(string)
	}
	if val, ok := m.params["version"]; ok {
		if !infoExists {
			res += infoString
			infoExists = true
		}
		res += infoVersionString + val.(string)
	}
	if val, ok := m.params["basePath"]; !ok {
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
			str, err := def.ToString()
			if err != nil {
				return "", err
			}
			res += strings.Replace(str, "\n", mainIndentString, -1)
		}
	}
	return res, nil
}

func NewMainSwagg(params map[string]interface{}, paths []PathSwagger, def []parameters.SwaggParameter) MainSwagg {
	return &mainSwagg{
		params:      params,
		paths:       paths,
		definitions: def,
	}
}
