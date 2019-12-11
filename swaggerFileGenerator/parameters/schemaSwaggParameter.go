package parameters

import (
	"errors"
	"strconv"
	"strings"
)

const (
	inDeficeString     = "\n- in: "
	nameString         = "\n  name: "
	requiredString     = "\n  required: "
	descriptionString  = "\n  description: "
	linkOnSchemaString = "\n  schema:"
	refString          = "\n    $ref: '#/definitions/"
)

type schemaSwaggParameter struct {
	configs map[string]interface{}
	obj     SwaggParameter
}

var (
	errorEmptyIn     = errors.New("ERROR_IN_IS_MANDATORY")
	errorEmptyName   = errors.New("ERROR_NAME_IS_MANDATORY")
	errorEmptySchema = errors.New("EMPTY_SCHEMA_LINK")
)

func (a *schemaSwaggParameter) ToString() (string, error) {
	var res string
	if a.configs == nil {
		return "", errorEmptyIn
	}
	if val, ok := a.configs["in"]; ok && val.(string) != "" {
		res += inDeficeString + val.(string)
	} else {
		return "", errorEmptyIn
	}
	if val, ok := a.configs["name"]; ok && val.(string) != "" {
		res += nameString + val.(string)
	} else {
		return "", errorEmptyName
	}
	if val, ok := a.configs["description"]; ok && val.(string) != "" {
		res += descriptionString + val.(string)
	}
	if a.obj != nil && !a.IsObject() {
		if val, ok := a.configs["required"]; ok && val != "" {
			res += requiredString + strconv.FormatBool(val.(bool))
		}
	}
	if a.obj != nil {
		str, err := a.obj.ToString()
		if err != nil {
			return "", err
		}
		if a.configs["in"] == "body" {
			res += linkOnSchemaString + strings.ReplaceAll(str, "\n", "\n  ")
			return res, nil
		}
		res += str
		return res, nil
	}
	res += linkOnSchemaString + refString + a.configs["schema"].(string) + "'"
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

func NewSchemaSwaggParameter(parameter SwaggParameter) SwaggParameter {
	configs := parameter.getConfigs()
	if val, ok := configs["nameOfVariable"]; ok {
		configs["schema"] = val
		return &schemaSwaggParameter{
			configs: configs,
			obj:     nil,
		}
	}
	return &schemaSwaggParameter{
		configs: configs,
		obj:     parameter,
	}
}
