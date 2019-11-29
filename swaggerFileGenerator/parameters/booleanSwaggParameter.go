package parameters

import (
	"strconv"
)

const (
	boolType = "boolean"
)

type boolSwaggParameter struct {
	params map[string]interface{}
}

func (s *boolSwaggParameter) ToString() (string, error) {
	res := typeString + boolType
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
	if val, ok := s.params["allowEmptyValue"]; ok {
		res += allowEmptyValueString + strconv.FormatBool(val.(bool))
	}
	if val, ok := s.params["description"]; ok {
		res += descriptionString + val.(string)
	}
	return res, nil
}

func NewBoolSwagParameter(params map[string]interface{}) SwaggParameter {
	return &boolSwaggParameter{
		params: params,
	}
}
