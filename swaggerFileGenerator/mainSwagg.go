package swaggerFileGenerator

import (
	"SwaggerGin/swaggerFileGenerator/parameters"
	"errors"
	"strings"
)

const (
	swaggerString         = "swagger: '2.0'"
	infoString            = "\ninfo:"
	infoTitleString       = "\n  title:"
	infoDescriptionString = "\n  title:"
	infoVersionString     = "\n  version:"
	basePathString        = "\nbasePath:"
	pathsString           = "\npaths:"
	definitionsString     = "\ndefinitions:"

	mainIndentString = "\n  "

	errorEmptyPaths = "EMPTY_PATHS"
)

type MainSwagg interface {
	ToString() (string, error)
}

type mainSwagg struct {
	title       string
	description string
	version     string
	basePath    string
	paths       []PathSwagger
	definitions []parameters.SwaggParameter
	//TODO: security
}

func (m *mainSwagg) ToString() (string, error) {
	if m.paths == nil {
		return "", errors.New(errorEmptyPaths)
	}
	res := swaggerString
	res += infoString
	res += infoTitleString + m.description
	res += infoDescriptionString + m.title
	res += infoVersionString + m.version
	if m.basePath == "" {
		res += basePathString + "/"
	} else {
		res += basePathString + m.basePath
	}
	res += pathsString
	for _, path := range m.paths {
		str, err := path.ToString()
		if err != nil {
			return "", err
		}
		res += strings.Replace(str, "\n", mainIndentString, -1)
	}
	res += definitionsString
	for _, def := range m.definitions {
		str, err := def.ToString()
		if err != nil {
			return "", err
		}
		res += strings.Replace(str, "\n", mainIndentString, -1)
	}
	return res, nil
}

func NewMainSwagg(title, descr, vers, bPath string, paths []PathSwagger, def []parameters.SwaggParameter) MainSwagg {
	return &mainSwagg{
		title:       title,
		description: descr,
		version:     vers,
		basePath:    bPath,
		paths:       paths,
		definitions: def,
	}
}
