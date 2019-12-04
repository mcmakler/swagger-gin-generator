package parameters

import (
	"errors"
	"strconv"
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

func (o *objectSwaggerParameter) ToString(isDefinition bool) (string, error) {
	//TODO: other parameters (description, required ...)
	if o.configs == nil {
		return "", errorNilObjectName
	}
	if _, ok := o.configs["name"]; !ok {
		return "", errorNilObjectName
	}
	res := ""
	if !o.subObject {
		res = "\n" + o.configs["name"].(string) + ":"
	}
	if isDefinition {
		res += typeString + objectType
	} else {
		res += typeDeficeString + objectType
	}
	if val, ok := o.configs["required"]; ok {
		res += requiredString + strconv.FormatBool(val.(bool))
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

func NewObjectSwaggerParameter(params map[string]interface{}, prop map[string]SwaggParameter, subObj bool) SwaggParameter {
	return &objectSwaggerParameter{
		configs:    params,
		properties: prop,
		subObject:  subObj,
	}
}
