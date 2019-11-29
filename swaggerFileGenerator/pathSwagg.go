package swaggerFileGenerator

import (
	"errors"
	"strings"
)

const (
	requestsIndentString = "\n  "

	errorIncorrectPath = "INCORRECT_PATH"
	errorNullRequests  = "EMPTY_REQUESTS_ARRAY"
)

type PathSwagger interface {
	ToString() (string, error)
}

type pathSwagger struct {
	path     string
	requests []RequestSwagg
}

func (p *pathSwagger) ToString() (string, error) {
	if p.path == "" {
		return "", errors.New(errorIncorrectPath)
	}
	if p.requests == nil {
		return "", errors.New(errorNullRequests)
	}
	res := "\n" + p.path + ":"
	for _, request := range p.requests {
		str, err := request.ToString()
		if err != nil {
			return "", err
		}
		res += strings.Replace(str, "\n", requestsIndentString, -1)
	}
	return res, nil
}

func NewPathSwagger(path string, req []RequestSwagg) PathSwagger {
	return &pathSwagger{
		path:     path,
		requests: req,
	}
}
