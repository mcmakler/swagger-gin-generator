package swaggerFileGenerator

import (
	"errors"
	"fmt"
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
		res += strings.Replace(str, "\n", requestsIndentString, -1)
	}
	return res, nil
}

func ginPathToSwaggerPath(path string) string {
	fmt.Println(path)
	split := strings.Split(path, "/")
	path = ""
	for _, val := range split {
		fmt.Println(val)
		path += ""
		if len(path) > 1 {
			if string(val[1]) == ":" {
				val = "{" + val[2:] + "}"
			}
		}
		path += val
	}
	return path
}

func NewPathSwagger(path string, requests []RequestSwagger) PathSwagger {
	return &pathSwagger{
		path:     path,
		requests: requests,
	}
}
