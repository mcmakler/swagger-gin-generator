package swaggerFileGenerator

import (
	"errors"
	"github.com/mcmakler/swagger-gin-generator/swaggerFileGenerator/parameters"
	"strconv"
	"strings"
)

const (
	linkOnSchemaString = "\n  schema:"
	refString = "\n    $ref: '#/definitions/"
)

var (
	errorWrongCode        = errors.New("ERROR_WRONG_HTTP_CODE")
	errorEmptyDescription = errors.New("ERROR_DESCRIPTION_IS_EMPTY")
)

type ResponseSwagger interface {
	ToString() (string, error)
}

type responseSwagger struct {
	code         int
	description  string
	linkOnSchema string
	parameter    parameters.SwaggParameter
}

func (r *responseSwagger) ToString() (string, error) {
	if r.code < 0 { //TODO: check correct code
		return "", errorWrongCode
	}
	res := "\n'" + strconv.FormatInt(int64(r.code), 10) + "':"
	if r.description != "" {
		res += descriptionString + r.description
	} else {
		return "", errorEmptyDescription
	}
	if r.parameter != nil {
		res += linkOnSchemaString
		str, err := r.parameter.ToString()
		if err != nil {
			return "", err
		}
		res += strings.Replace(str, "\n", parametersIndentString, -1)
	} else if r.linkOnSchema != "" {
		res += linkOnSchemaString + refString +  r.linkOnSchema + "'"
	}
	return res, nil
}

func NewResponseSwagger(code int, descr, schema string, parameter parameters.SwaggParameter) ResponseSwagger {
	return &responseSwagger{
		code:         code,
		description:  descr,
		linkOnSchema: schema,
		parameter:    parameter,
	}
}
