package parameters

import (
	"strconv"
)

const (
	numberType = "number"

	digitsAfterDot = -1
)

type numberSwaggParameter struct {
	params map[string]interface{}
}

func (i *numberSwaggParameter) ToString(isDefinition bool) (string, error) {
	var res string
	if isDefinition {
		res = typeString + numberType
	} else {
		res = typeDeficeString + numberType
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
		res += defaultValueString + strconv.FormatFloat(val.(float64), 'f', digitsAfterDot, 64)
	}
	if val, ok := i.params["minimumValue"]; ok {
		res += minimumValueString + strconv.FormatFloat(val.(float64), 'f', digitsAfterDot, 64)
	}
	if val, ok := i.params["exclusiveMinimumValue"]; ok {
		res += exclusiveMinimumValueString + strconv.FormatBool(val.(bool))
	}
	if val, ok := i.params["maximumValue"]; ok {
		res += maximumValueString + strconv.FormatFloat(val.(float64), 'f', digitsAfterDot, 64)
	}
	if val, ok := i.params["exclusiveMaximumValue"]; ok {
		res += exclusiveMaximumValueString + strconv.FormatBool(val.(bool))
	}
	if val, ok := i.params["multipleOf"]; ok {
		res += multipleOfString + strconv.FormatFloat(val.(float64), 'f', digitsAfterDot, 64)
	}
	if val, ok := i.params["allowEmptyValue"]; ok {
		res += allowEmptyValueString + strconv.FormatBool(val.(bool))
	}
	if val, ok := i.params["description"]; ok {
		res += descriptionString + val.(string)
	}

	return res, nil
}

func NewNumberSwagParameter(params map[string]interface{}) SwaggParameter {
	return &numberSwaggParameter{
		params: params,
	}
}
