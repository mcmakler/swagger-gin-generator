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

func (a *integerSwaggParameter) ToString() (string, error) {
	var res string
	res = typeString + integerType
	if a.configs == nil {
		return res, nil
	}
	if val, ok := a.configs["defaultValue"]; ok {
		res += defaultValueString + strconv.FormatInt(int64(val.(int)), 10)
	}
	if val, ok := a.configs["minimumValue"]; ok {
		res += minimumValueString + strconv.FormatInt(int64(val.(int)), 10)
	}
	if val, ok := a.configs["exclusiveMinimumValue"]; ok {
		res += exclusiveMinimumValueString + strconv.FormatBool(val.(bool))
	}
	if val, ok := a.configs["maximumValue"]; ok {
		res += maximumValueString + strconv.FormatInt(int64(val.(int)), 10)
	}
	if val, ok := a.configs["exclusiveMaximumValue"]; ok {
		res += exclusiveMaximumValueString + strconv.FormatBool(val.(bool))
	}
	if val, ok := a.configs["multipleOf"]; ok {
		res += multipleOfString + strconv.FormatInt(int64(val.(int)), 10)
	}
	return res, nil
}

func (a *integerSwaggParameter) IsObject() bool {
	return false
}

func (a *integerSwaggParameter) getConfigs() map[string]interface{} {
	return a.configs
}

func NewIntegerSwagParameter(params map[string]interface{}) SwaggParameter {
	return &integerSwaggParameter{
		configs: params,
	}
}
