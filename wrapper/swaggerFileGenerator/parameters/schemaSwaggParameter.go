package parameters

import "errors"

const (
	linkOnSchemaString = "\n  schema:\n    $ref: '#/definitions/"
)

type schemaSwaggParameter struct {
	configs map[string]interface{}
}

var (
	errorEmptySchema = errors.New("EMPTY_SCHEMA_LINK")
)

func (a *schemaSwaggParameter) ToString(isDefinition bool) (string, error) {
	var res string
	if val, ok := a.configs["in"]; ok {
		res += inString + val.(string)
	} else {
		return "", errorEmptyIn
	}
	if val, ok := a.configs["name"]; ok {
		res += nameString + val.(string)
	} else {
		return "", errorEmptyName
	}
	if val, ok := a.configs["description"]; ok {
		res += descriptionString + val.(string)
	}
	if val, ok := a.configs["schema"]; ok {
		res += nameString + val.(string)
	} else {
		return "", errorEmptySchema
	}
	return res, nil
}

func (a *schemaSwaggParameter) IsObject() bool {
	return false
}

func (a *schemaSwaggParameter) getConfigs() map[string]interface{} {
	return a.configs
}

func NewSchemaSwaggParameter(params SwaggParameter) SwaggParameter {
	configs := params.getConfigs()
	configs["schema"] = configs["nameOfVariable"]
	return &schemaSwaggParameter{
		configs: configs,
	}
}
