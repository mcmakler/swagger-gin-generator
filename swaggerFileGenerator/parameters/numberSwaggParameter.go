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
	res = typeString + numberType
	if a.configs == nil {
		return res, nil
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
