package parameters

import (
	"strconv"
)

const (
	stringType = "string"

	formatString    = "\n  format: "
	minLengthString = "\n  minLength: "
	maxLengthString = "\n  maxLength: "
	patternString   = "\n  pattern: "
	enumString      = "\n  enum: "
	enumNewString   = "\n- "
)

type stringSwaggParameter struct {
	params          map[string]interface{}
}

func (s *stringSwaggParameter) ToString() (string, error) {
	res := typeString + stringType
	if s.params == nil {
		return res, nil
	}
	if val, ok := s.params["in"]; ok {
		res += inString + val.(string)
	}
	if val, ok := s.params["name"]; ok {
		res += nameString + val.(string)
	}
	if val, ok := s.params["required"]; ok {
		res += requiredString + strconv.FormatBool(val.(bool))
	}
	if val, ok := s.params["format"]; ok { //TODO: make checking of format
		res += formatString + val.(string)
	}
	if val, ok := s.params["minLength"]; ok { //TODO: make check?
		res += minLengthString + strconv.FormatInt(int64(val.(int)), 10)
	}
	if val, ok := s.params["maxLength"]; ok {
		res += maxLengthString + strconv.FormatInt(int64(val.(int)), 10)
	}
	if val, ok := s.params["pattern"]; ok {
		res += patternString + val.(string)
	}
	if val, ok := s.params["allowEmptyValue"]; ok {
		res += allowEmptyValueString + strconv.FormatBool(val.(bool))
	}
	if val, ok := s.params["description"]; ok {
		res += descriptionString + val.(string)
	}
	if val, ok := s.params["enum"]; ok {
		res += enumString
		for _, enum := range val.([]string) {
			res += enumNewString + enum
		}
	}
	return res, nil
}

func NewStringSwagParameter(params map[string]interface{}) SwaggParameter {
	return &stringSwaggParameter{
		params: params,
	}
}
