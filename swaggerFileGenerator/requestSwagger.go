package swaggerFileGenerator

import (
	"errors"
	"github.com/mcmakler/swagger-gin-generator/swaggerFileGenerator/parameters"
	"strings"
)

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

type RequestSwagger interface {
	ToString() (string, error)
}

type requestSwagger struct {
	config     map[string]interface{}
	parameters []parameters.SwaggParameter
	responses  []ResponseSwagger
}

var (
	errorEmptyTypeRequest = errors.New("EMPTY_TYPE_OF_REQUEST")
	errorEmptyResponses   = errors.New("EMPTY_RESPONSES")
)

func (r *requestSwagger) ToString() (string, error) {
	if r.config == nil {
		return "", errorEmptyTypeRequest
	}
	if _, ok := r.config["typeRequest"]; !ok {
		return "", errorEmptyTypeRequest
	}
	if r.responses == nil || len(r.responses) == 0 {
		return "", errorEmptyResponses
	}
	res := "\n" + r.config["typeRequest"].(string) + ":"
	if val, ok := r.config["security"]; ok && val.([]string) != nil {
		res += securityString
		for _, security := range val.([]string) {
			res += securityIndentString + security + ": []"
		}
	}
	if val, ok := r.config["description"]; ok && val != "" {
		res += descriptionString + val.(string)
	}
	if val, ok := r.config["consumes"]; ok && val.([]string) != nil {
		res += consumesString
		for _, consumes := range val.([]string) {
			res += consumesIndentString + consumes
		}
	}
	if val, ok := r.config["produces"]; ok && val.([]string) != nil {
		res += producesString
		for _, produce := range val.([]string) {
			res += producesIndentString + produce
		}
	}
	if val, ok := r.config["tags"]; ok && val.([]string) != nil {
		res += tagsString
		for _, tag := range val.([]string) {
			res += tagsIndentString + tag
		}
	}
	if val, ok := r.config["operationId"]; ok && val != "" {
		res += operationIdString + val.(string)
	}
	if val, ok := r.config["summary"]; ok && val != "" {
		res += summaryString + val.(string)
	}
	if r.parameters != nil {
		res += parametersString
		for _, parameter := range r.parameters {
			str, err := parameter.ToString()
			if err != nil {
				return "", err
			}
			res += strings.ReplaceAll(str, "\n", parametersIndentString)
		}
	}
	if r.responses != nil {
		res += responsesString
		for _, resp := range r.responses {
			str, err := resp.ToString()
			if err != nil {
				return "", err
			}
			res += strings.ReplaceAll(str, "\n", responsesIndentString)
		}
	}
	return res, nil
}

func NewRequestSwagger(params map[string]interface{}, parameters []parameters.SwaggParameter, resp []ResponseSwagger) RequestSwagger {
	return &requestSwagger{
		config:     params,
		parameters: parameters,
		responses:  resp,
	}
}
