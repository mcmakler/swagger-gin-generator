package swaggerFileGenerator

import (
	"errors"
	"strings"
)

const (
	requestsIndentString = "\n  "
)

var (
	errorIncorrectPath = errors.New("INCORRECT_PATH")
	errorNullRequests  = errors.New("EMPTY_REQUESTS_ARRAY")
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
		return "", errorIncorrectPath
	}
	if p.requests == nil {
		return "", errorNullRequests
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
