package parameters

import (
	"strconv"
)

const (
	integerType = "integer"
)

type integerSwaggParameter struct {
	configs map[string]interface{}
}

func (i *integerSwaggParameter) ToString(isDefinition bool) (string, error) {
	var res string
	if isDefinition {
		res = typeString + integerType
	} else {
		res = typeDeficeString + integerType
	}
	if i.configs == nil {
		return res, nil
	}
	if val, ok := i.configs["in"]; ok {
		res += inString + val.(string)
	} else if !isDefinition {
		return "", errorEmptyIn
	}
	if val, ok := i.configs["name"]; ok {
		res += nameString + val.(string)
	} else if !isDefinition {
		return "", errorEmptyName
	}
	if val, ok := i.configs["required"]; ok {
		res += requiredString + strconv.FormatBool(val.(bool))
	}
	if val, ok := i.configs["defaultValue"]; ok {
		res += defaultValueString + strconv.FormatInt(int64(val.(int)), 10)
	}
	if val, ok := i.configs["minimumValue"]; ok {
		res += minimumValueString + strconv.FormatInt(int64(val.(int)), 10)
	}
	if val, ok := i.configs["exclusiveMinimumValue"]; ok {
		res += exclusiveMinimumValueString + strconv.FormatBool(val.(bool))
	}
	if val, ok := i.configs["maximumValue"]; ok {
		res += maximumValueString + strconv.FormatInt(int64(val.(int)), 10)
	}
	if val, ok := i.configs["exclusiveMaximumValue"]; ok {
		res += exclusiveMaximumValueString + strconv.FormatBool(val.(bool))
	}
	if val, ok := i.configs["multipleOf"]; ok {
		res += multipleOfString + strconv.FormatInt(int64(val.(int)), 10)
	}
	if val, ok := i.configs["allowEmptyValue"]; ok {
		res += allowEmptyValueString + strconv.FormatBool(val.(bool))
	}
	if val, ok := i.configs["description"]; ok {
		res += descriptionString + val.(string)
	}

	return res, nil
}

func NewIntegerSwagParameter(params map[string]interface{}) SwaggParameter {
	return &integerSwaggParameter{
		configs: params,
	}
}
