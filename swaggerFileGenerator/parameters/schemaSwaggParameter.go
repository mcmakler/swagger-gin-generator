package parameters

import (
	"errors"
	"strconv"
	"strings"
)

const (
	refString = "\n    $ref: '#/definitions/"
)

type schemaSwaggParameter struct {
	configs map[string]interface{}
	obj     SwaggParameter
}

var (
	errorEmptySchema = errors.New("EMPTY_SCHEMA_LINK")
)

func (a *schemaSwaggParameter) ToString(isDefinition bool) (string, error) {
	var res string
	if val, ok := a.configs["in"]; ok && val != "" {
		res += inDeficeString + val.(string)
	} else {
		return "", errorEmptyIn
	}
	if val, ok := a.configs["name"]; ok && val != "" {
		res += nameString + val.(string)
	} else {
		return "", errorEmptyName
	}
	if val, ok := a.configs["description"]; ok && val != "" {
		res += descriptionString + val.(string)
	}
	if val, ok := a.configs["required"]; ok && val != ""{
		res += requiredString + strconv.FormatBool(val.(bool))
	}
	if a.obj != nil {
		str, err := a.obj.ToString(false)
		if err != nil {
			return "", err
		}
		res += linkOnSchemaString + strings.Replace(str, "\n", "\n  ", -1)
		return res, nil
	}
	if val, ok := a.configs["schema"]; ok {
		res += linkOnSchemaString + refString + val.(string) + "'"
	} else {
		return "", errorEmptySchema
	}
	return res, nil
}

func (a *schemaSwaggParameter) IsObject() bool {
	if a.obj == nil {
		return false
	}
	return a.obj.IsObject()
}

func (a *schemaSwaggParameter) getConfigs() map[string]interface{} {
	return a.configs
}

func NewSchemaSwaggParameter(params SwaggParameter) SwaggParameter {
	configs := params.getConfigs()
	if val, ok := configs["nameOfVariable"]; ok {
		configs["schema"] = val
		return &schemaSwaggParameter{
			configs: configs,
			obj:     nil,
		}
	}
	return &schemaSwaggParameter{
		configs: configs,
		obj:     params,
	}
}
