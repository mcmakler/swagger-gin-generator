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
	errorWrongCode        = errors.New("WRONG_HTTP_CODE")
	errorEmptyDescription = errors.New("DESCRIPTION_IS_EMPTY")
)

type ResponseSwagg interface {
	ToString() (string, error)
}

type responseSwagg struct {
	code         int
	description  string
	linkOnSchema string
	parameter    parameters.SwaggParameter
}

func (r *responseSwagg) ToString() (string, error) {
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
		str, err := r.parameter.ToString(true)
		if err != nil {
			return "", err
		}
		res += strings.Replace(str, "\n", parametersIndentString, -1)
	} else if r.linkOnSchema != "" {
		res += linkOnSchemaString + refString +  r.linkOnSchema + "'"
	}
	return res, nil
}

func NewResponseSwagg(code int, descr, schema string, parameter parameters.SwaggParameter) ResponseSwagg {
	return &responseSwagg{
		code:         code,
		description:  descr,
		linkOnSchema: schema,
		parameter:    parameter,
	}
}
