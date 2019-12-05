package parameters

import (
	"strconv"
)

const (
	integerType = "integer"
)

type integerSwaggerParameter struct {
	config map[string]interface{}
}

func (a *integerSwaggerParameter) ToString() (string, error) {
	var res string
	res = typeString + integerType
	if a.config == nil {
		return res, nil
	}
	if val, ok := a.config["defaultValue"]; ok {
		res += defaultValueString + strconv.FormatInt(int64(val.(int)), 10)
	}
	if val, ok := a.config["minimumValue"]; ok {
		res += minimumValueString + strconv.FormatInt(int64(val.(int)), 10)
	}
	if val, ok := a.config["exclusiveMinimumValue"]; ok {
		res += exclusiveMinimumValueString + strconv.FormatBool(val.(bool))
	}
	if val, ok := a.config["maximumValue"]; ok {
		res += maximumValueString + strconv.FormatInt(int64(val.(int)), 10)
	}
	if val, ok := a.config["exclusiveMaximumValue"]; ok {
		res += exclusiveMaximumValueString + strconv.FormatBool(val.(bool))
	}
	if val, ok := a.config["multipleOf"]; ok {
		res += multipleOfString + strconv.FormatInt(int64(val.(int)), 10)
	}
	return res, nil
}

func (a *integerSwaggerParameter) IsObject() bool {
	return false
}

func (a *integerSwaggerParameter) getConfigs() map[string]interface{} {
	return a.config
}

func NewIntegerSwaggerParameter(config map[string]interface{}) SwaggParameter {
	return &integerSwaggerParameter{
		config: config,
	}
}
