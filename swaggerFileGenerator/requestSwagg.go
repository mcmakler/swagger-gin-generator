package swaggerFileGenerator

import (
	"errors"
	"github.com/mcmakler/swagger-gin-generator/swaggerFileGenerator/parameters"
	"strings"
)

//TODO: accept

const (
	descriptionString      = "\n  description: "
	securityString         = "\n  security:"
	securityIndentString   = "\n  - "
	consumesString         = "\n  consumes:"
	consumesIndentString   = "\n  - "
	producesString         = "\n  produces: "
	producesIndentString   = "\n  - "
	tagsString             = "\n  tags: "
	tagsIndentString       = "\n    - "
	operationIdString      = "\n  operationId: "
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
	if val, ok := r.configs["security"]; ok && val.([]string) != nil {
		res += securityString
		for _, cons := range val.([]string) {
			res += securityIndentString + cons + ": []"
		}
	}
	if val, ok := r.configs["description"]; ok && val != "" {
		res += descriptionString + val.(string)
	}
	if val, ok := r.configs["consumes"]; ok && val.([]string) != nil {
		res += consumesString
		for _, cons := range val.([]string) {
			res += consumesIndentString + cons
		}
	}
	if val, ok := r.configs["produces"]; ok && val.([]string) != nil {
		res += producesString
		for _, prod := range val.([]string) {
			res += producesIndentString + prod
		}
	}
	if val, ok := r.configs["tags"]; ok && val.([]string) != nil {
		res += tagsString
		for _, tag := range val.([]string) {
			res += tagsIndentString + tag
		}
	}
	if val, ok := r.configs["operationId"]; ok && val != "" {
		res += operationIdString + val.(string)
	}
	if val, ok := r.configs["summary"]; ok && val != "" {
		res += summaryString + val.(string)
	}
	if r.parameters != nil {
		res += parametersString
		for _, param := range r.parameters {
			str, err := param.ToString()
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
