package parameters

import (
	"strconv"
)

const (
	boolType = "boolean"
)

type boolSwaggParameter struct {
	configs map[string]interface{}
}

func (s *boolSwaggParameter) ToString(isDefinition bool) (string, error) {
	var res string
	if isDefinition {
		res = typeString + boolType
	} else {
		res = typeDeficeString + boolType
	}
	if s.configs == nil {
		return res, nil
	}
	if val, ok := s.configs["in"]; ok {
		res += inString + val.(string)
	} else if !isDefinition {
		return "", errorEmptyIn
	}
	if val, ok := s.configs["name"]; ok {
		res += nameString + val.(string)
	} else if !isDefinition {
		return "", errorEmptyName
	}
	if val, ok := s.configs["required"]; ok {
		res += requiredString + strconv.FormatBool(val.(bool))
	}
	if val, ok := s.configs["allowEmptyValue"]; ok {
		res += allowEmptyValueString + strconv.FormatBool(val.(bool))
	}
	if val, ok := s.configs["description"]; ok {
		res += descriptionString + val.(string)
	}
	return res, nil
}

func NewBoolSwagParameter(params map[string]interface{}) SwaggParameter {
	return &boolSwaggParameter{
		configs: params,
	}
}
