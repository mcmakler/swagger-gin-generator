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

func (a *boolSwaggParameter) ToString(isDefinition bool) (string, error) {
	var res string
	if isDefinition {
		res = typeString + boolType
	} else {
		res = typeDeficeString + boolType
	}
	if a.configs == nil {
		return res, nil
	}
	if val, ok := a.configs["in"]; ok {
		res += inString + val.(string)
	} else if !isDefinition {
		return "", errorEmptyIn
	}
	if val, ok := a.configs["name"]; ok {
		res += nameString + val.(string)
	} else if !isDefinition {
		return "", errorEmptyName
	}
	if val, ok := a.configs["required"]; ok {
		res += requiredString + strconv.FormatBool(val.(bool))
	}
	if val, ok := a.configs["allowEmptyValue"]; ok {
		res += allowEmptyValueString + strconv.FormatBool(val.(bool))
	}
	if val, ok := a.configs["description"]; ok {
		res += descriptionString + val.(string)
	}
	return res, nil
}

func (a *boolSwaggParameter) IsObject() bool {
	return false
}

func (a *boolSwaggParameter) getConfigs() map[string]interface{} {
	return a.configs
}

func NewBoolSwagParameter(params map[string]interface{}) SwaggParameter {
	return &boolSwaggParameter{
		configs: params,
	}
}
