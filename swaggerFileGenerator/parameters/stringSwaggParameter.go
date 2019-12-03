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
	params map[string]interface{}
}

func (s *stringSwaggParameter) ToString(isDefinition bool) (string, error) {
	var res string
	if isDefinition {
		res = typeString + stringType
	} else {
		res = typeDeficeString + stringType
	}
	if s.params == nil {
		return res, nil
	}
	if val, ok := s.params["in"]; ok {
		res += inString + val.(string)
	} else if !isDefinition {
		return "", errorEmptyIn
	}
	if val, ok := s.params["name"]; ok {
		res += nameString + val.(string)
	} else if !isDefinition {
		return "", errorEmptyName
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
