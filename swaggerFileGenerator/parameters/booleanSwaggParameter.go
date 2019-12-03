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

func (s *boolSwaggParameter) ToString(isDefinition bool) (string, error) {
	var res string
	if isDefinition {
		res = typeString + boolType
	} else {
		res = typeDeficeString + boolType
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
