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
	params     map[string]interface{}
	properties map[string]SwaggParameter
}

var (
	errorNilObjectName = errors.New("NIL_OBJECT_NAME")
)

func (o *objectSwaggerParameter) ToString(isDefinition bool) (string, error) {
	//TODO: other parameters (description, required ...)
	if o.params == nil {
		return "", errorNilObjectName
	}
	if _, ok := o.params["name"]; !ok {
		return "", errorNilObjectName
	}
	res := "\n" + o.params["name"].(string) + ":"
	if isDefinition {
		res += typeString + objectType
	} else {
		res += typeDeficeString + objectType
	}
	if val, ok := o.params["required"]; ok {
		for _, val := range val.([]string) {
			res += requiredIndentStr + val
		}
	}
	if o.properties != nil {
		res += propertiesStr
		for index, val := range o.properties {
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

func NewObjectSwaggerParameter(params map[string]interface{}, prop map[string]SwaggParameter) SwaggParameter {
	return &objectSwaggerParameter{
		params:     params,
		properties: prop,
	}
}
