package parameters

import (
	"errors"
	"strings"
)

const (
	objectType = "object"

	requiredIndentStr = "\n    - "
	propertiesStr     = "\n  properties:"
	propertyIndentStr = "\n    "
)

type objectSwaggerParameter struct {
	configs    map[string]interface{}
	properties map[string]SwaggParameter
	subObject  bool
}

var (
	errorNilObjectName = errors.New("NIL_OBJECT_NAME")
)

func (a *objectSwaggerParameter) ToString(isDefinition bool) (string, error) {
	//TODO: other parameters (description, required ...)
	if a.configs == nil {
		return "", errorNilObjectName
	}
	if _, ok := a.configs["nameOfVariable"]; !ok {
		return "", errorNilObjectName
	}
	res := ""
	if !a.subObject {
		res = "\n" + a.configs["nameOfVariable"].(string) + ":"
	}
	res += typeString + objectType
	if val, ok := a.configs["required"]; ok && val != nil {
		for _, val := range val.([]string) {
			res += requiredIndentStr + val
		}
	}
	if a.properties != nil {
		res += propertiesStr
		for index, val := range a.properties {
			res += propertyIndentStr + index + ":"
			str, err := val.ToString(isDefinition)
			if err != nil {
				return "", err
			}
			res += strings.Replace(str, "\n", propertyIndentStr, -1)
		}
	}
	return res, nil
}

func (a *objectSwaggerParameter) IsObject() bool {
	return true
}

func (a *objectSwaggerParameter) getConfigs() map[string]interface{} {
	return a.configs
}

func NewObjectSwaggerParameter(params map[string]interface{}, prop map[string]SwaggParameter, subObj bool) SwaggParameter {
	return &objectSwaggerParameter{
		configs:    params,
		properties: prop,
		subObject:  subObj,
	}
}
