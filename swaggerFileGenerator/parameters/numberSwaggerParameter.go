package parameters

import (
	"strconv"
)

const (
	numberType = "number"

	digitsAfterDot = -1
)

type numberSwaggerParameter struct {
	config map[string]interface{}
}

func (a *numberSwaggerParameter) ToString() (string, error) {
	var res string
	res = typeString + numberType
	if a.config == nil {
		return res, nil
	}
	if val, ok := a.config["defaultValue"]; ok {
		res += defaultValueString + strconv.FormatFloat(val.(float64), 'f', digitsAfterDot, 64)
	}
	if val, ok := a.config["minimumValue"]; ok {
		res += minimumValueString + strconv.FormatFloat(val.(float64), 'f', digitsAfterDot, 64)
	}
	if val, ok := a.config["exclusiveMinimumValue"]; ok {
		res += exclusiveMinimumValueString + strconv.FormatBool(val.(bool))
	}
	if val, ok := a.config["maximumValue"]; ok {
		res += maximumValueString + strconv.FormatFloat(val.(float64), 'f', digitsAfterDot, 64)
	}
	if val, ok := a.config["exclusiveMaximumValue"]; ok {
		res += exclusiveMaximumValueString + strconv.FormatBool(val.(bool))
	}
	if val, ok := a.config["multipleOf"]; ok {
		res += multipleOfString + strconv.FormatFloat(val.(float64), 'f', digitsAfterDot, 64)
	}
	return res, nil
}

func (a *numberSwaggerParameter) IsObject() bool {
	return false
}

func (a *numberSwaggerParameter) getConfigs() map[string]interface{} {
	return a.config
}

func NewNumberSwaggerParameter(config map[string]interface{}) SwaggParameter {
	return &numberSwaggerParameter{
		config: config,
	}
}
