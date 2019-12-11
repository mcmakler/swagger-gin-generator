package swaggerFileGenerator

import (
	"errors"
	"strings"
)

const (
	requestsIndentString = "\n  "
)

var (
	errorIncorrectPath = errors.New("ERROR_INCORRECT_PATH")
	errorNullRequests  = errors.New("ERROR_EMPTY_REQUESTS_ARRAY")
)

type PathSwagger interface {
	ToString() (string, error)
}

type pathSwagger struct {
	path     string
	requests []RequestSwagger
}

func (p *pathSwagger) ToString() (string, error) {
	if p.path == "" {
		return "", errorIncorrectPath
	}
	if p.requests == nil {
		return "", errorNullRequests
	}
	res := "\n" + ginPathToSwaggerPath(p.path) + ":"
	for _, request := range p.requests {
		str, err := request.ToString()
		if err != nil {
			return "", err
		}
		res += strings.ReplaceAll(str, "\n", requestsIndentString)
	}
	return res, nil
}

func ginPathToSwaggerPath(path string) string {
	split := strings.Split(path, "/")
	path = ""
	for _, val := range split[1:] {
		if len(val) > 1 {
			if string(val[0]) == ":" {
				val = "{" + val[1:] + "}"
			}
		}
		path += "/" + val
	}
	return path
}

func NewPathSwagger(path string, requests []RequestSwagger) PathSwagger {
	return &pathSwagger{
		path:     path,
		requests: requests,
	}
}
