package parameters

import (
	"strconv"
)

const (
	numberType = "number"

	digitsAfterDot = -1
)

type numberSwaggParameter struct {
	configs map[string]interface{}
}

func (a *numberSwaggParameter) ToString(isDefinition bool) (string, error) {
	var res string
	if isDefinition {
		res = typeString + numberType
	} else {
		res = typeDeficeString + numberType
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
	if val, ok := a.configs["defaultValue"]; ok {
		res += defaultValueString + strconv.FormatFloat(val.(float64), 'f', digitsAfterDot, 64)
	}
	if val, ok := a.configs["minimumValue"]; ok {
		res += minimumValueString + strconv.FormatFloat(val.(float64), 'f', digitsAfterDot, 64)
	}
	if val, ok := a.configs["exclusiveMinimumValue"]; ok {
		res += exclusiveMinimumValueString + strconv.FormatBool(val.(bool))
	}
	if val, ok := a.configs["maximumValue"]; ok {
		res += maximumValueString + strconv.FormatFloat(val.(float64), 'f', digitsAfterDot, 64)
	}
	if val, ok := a.configs["exclusiveMaximumValue"]; ok {
		res += exclusiveMaximumValueString + strconv.FormatBool(val.(bool))
	}
	if val, ok := a.configs["multipleOf"]; ok {
		res += multipleOfString + strconv.FormatFloat(val.(float64), 'f', digitsAfterDot, 64)
	}
	if val, ok := a.configs["allowEmptyValue"]; ok {
		res += allowEmptyValueString + strconv.FormatBool(val.(bool))
	}
	if val, ok := a.configs["description"]; ok {
		res += descriptionString + val.(string)
	}

	return res, nil
}

func (a *numberSwaggParameter) IsObject() bool {
	return false
}

func (a *numberSwaggParameter) getConfigs() map[string]interface{} {
	return a.configs
}

func NewNumberSwagParameter(params map[string]interface{}) SwaggParameter {
	return &numberSwaggParameter{
		configs: params,
	}
}
