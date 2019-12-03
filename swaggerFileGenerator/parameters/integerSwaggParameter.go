package parameters

import (
	"strconv"
)

const (
	integerType = "integer"
)

type integerSwaggParameter struct {
	params map[string]interface{}
}

func (i *integerSwaggParameter) ToString(isDefinition bool) (string, error) {
	var res string
	if isDefinition {
		res = typeString + integerType
	} else {
		res = typeDeficeString + integerType
	}
	if i.params == nil {
		return res, nil
	}
	if val, ok := i.params["in"]; ok {
		res += inString + val.(string)
	} else if !isDefinition {
		return "", errorEmptyIn
	}
	if val, ok := i.params["name"]; ok {
		res += nameString + val.(string)
	} else if !isDefinition {
		return "", errorEmptyName
	}
	if val, ok := i.params["required"]; ok {
		res += requiredString + strconv.FormatBool(val.(bool))
	}
	if val, ok := i.params["defaultValue"]; ok {
		res += defaultValueString + strconv.FormatInt(int64(val.(int)), 10)
	}
	if val, ok := i.params["minimumValue"]; ok {
		res += minimumValueString + strconv.FormatInt(int64(val.(int)), 10)
	}
	if val, ok := i.params["exclusiveMinimumValue"]; ok {
		res += exclusiveMinimumValueString + strconv.FormatBool(val.(bool))
	}
	if val, ok := i.params["maximumValue"]; ok {
		res += maximumValueString + strconv.FormatInt(int64(val.(int)), 10)
	}
	if val, ok := i.params["exclusiveMaximumValue"]; ok {
		res += exclusiveMaximumValueString + strconv.FormatBool(val.(bool))
	}
	if val, ok := i.params["multipleOf"]; ok {
		res += multipleOfString + strconv.FormatInt(int64(val.(int)), 10)
	}
	if val, ok := i.params["allowEmptyValue"]; ok {
		res += allowEmptyValueString + strconv.FormatBool(val.(bool))
	}
	if val, ok := i.params["description"]; ok {
		res += descriptionString + val.(string)
	}

	return res, nil
}

func NewIntegerSwagParameter(params map[string]interface{}) SwaggParameter {
	return &integerSwaggParameter{
		params: params,
	}
}
