package swaggerFileGenerator

import (
	"errors"
	"strings"
	"swagger-gin-generator/swaggerFileGenerator/parameters"
)

const (
	descriptionString      = "\n  description: "
	consumesString         = "\n  consumes:"
	consumesIndentString   = "\n  - "
	producesString         = "\n  produces: "
	producesIndentString   = "\n  - "
	tagsString             = "\n  tags: "
	tagsIndentString       = "\n    - "
	summaryString          = "\n  summary: "
	parametersString       = "\n  parameters:"
	parametersIndentString = "\n    "
	responsesString        = "\n  responses:"
	responsesIndentString  = "\n    "
)

type RequestSwagg interface {
	ToString() (string, error)
}

type requestSwagg struct {
	configs    map[string]interface{}
	parameters []parameters.SwaggParameter
	responses  []ResponseSwagg
}

var (
	errorEmptyTypeRequest = errors.New("EMPTY_TYPE_OF_REQUEST")
	errorEmptyResponses   = errors.New("EMPTY_RESPONSES")
)

//TODO: some checks it need not be empty
func (r *requestSwagg) ToString() (string, error) {
	if r.configs == nil {
		return "", errorEmptyTypeRequest
	}
	if _, ok := r.configs["typeRequest"]; !ok {
		return "", errorEmptyTypeRequest
	}
	if r.responses == nil || len(r.responses) == 0 {
		return "", errorEmptyResponses
	}
	res := "\n" + r.configs["typeRequest"].(string) + ":"
	if val, ok := r.configs["description"]; ok {
		res += descriptionString + val.(string)
	}
	if val, ok := r.configs["consumes"]; ok {
		res += consumesString
		for _, cons := range val.([]string) {
			res += consumesIndentString + cons
		}
	}
	if val, ok := r.configs["produces"]; ok {
		res += producesString
		for _, prod := range val.([]string) {
			res += producesIndentString + prod
		}
	}
	if val, ok := r.configs["tags"]; ok {
		res += tagsString
		for _, tag := range val.([]string) {
			res += tagsIndentString + tag
		}
	}
	if val, ok := r.configs["summary"]; ok {
		res += summaryString + val.(string)
	}
	if r.parameters != nil {
		res += parametersString
		for _, param := range r.parameters {
			str, err := param.ToString(false)
			if err != nil {
				return "", err
			}
			res += strings.Replace(str, "\n", parametersIndentString, -1)
		}
	}
	if r.responses != nil {
		res += responsesString
		for _, resp := range r.responses {
			str, err := resp.ToString()
			if err != nil {
				return "", err
			}
			res += strings.Replace(str, "\n", responsesIndentString, -1)
		}
	}
	return res, nil
}

func NewRequestSwagg(params map[string]interface{}, parameters []parameters.SwaggParameter, resp []ResponseSwagg) RequestSwagg {
	return &requestSwagg{
		configs:    params,
		parameters: parameters,
		responses:  resp,
	}
}
