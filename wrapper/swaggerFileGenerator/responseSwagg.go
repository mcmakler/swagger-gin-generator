package swaggerFileGenerator

import (
	"errors"
	"strconv"
)

const (
	linkOnSchemaString = "\n  schema:\n    $ref: '#/definitions/"
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
	if r.linkOnSchema != "" {
		res += linkOnSchemaString + r.linkOnSchema + "'"
	}
	return res, nil
}

func NewResponseSwagg(code int, descr, schema string) ResponseSwagg {
	return &responseSwagg{
		code:         code,
		description:  descr,
		linkOnSchema: schema,
	}
}