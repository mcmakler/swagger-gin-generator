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
	config     map[string]interface{}
	properties map[string]SwaggParameter
	subObject  bool
}

var (
	errorNilObjectVariableName = errors.New("NIL_OBJECT_VARIABLE_NAME")
)

func (a *objectSwaggerParameter) ToString() (string, error) {
	if a.config == nil {
		return "", errorNilObjectVariableName
	}
	if _, ok := a.config["nameOfVariable"]; !ok {
		return "", errorNilObjectVariableName
	}
	res := ""
	if !a.subObject {
		res = "\n" + a.config["nameOfVariable"].(string) + ":"
	}
	res += typeString + objectType
	if val, ok := a.config["required"]; ok && val != nil {
		for _, val := range val.([]string) {
			res += requiredIndentStr + val
		}
	}
	if a.properties != nil {
		res += propertiesStr
		for index, val := range a.properties {
			res += propertyIndentStr + index + ":"
			str, err := val.ToString()
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
	return a.config
}

func NewObjectSwaggerParameter(config map[string]interface{}, properties map[string]SwaggParameter, subObject bool) SwaggParameter {
	return &objectSwaggerParameter{
		config:     config,
		properties: properties,
		subObject:  subObject,
	}
}
