package parameters

import (
	"strconv"
)

const (
	stringType = "string"

	formatString    = "\n  format: "
	minLengthString = "\n  minLength: "
	maxLengthString = "\n  maxLength: "
	patternString   = "\n  pattern: "
	enumString      = "\n  enum: "
	enumNewString   = "\n- "
)

type stringSwaggParameter struct {
	configs map[string]interface{}
}

func (a *stringSwaggParameter) ToString(isDefinition bool) (string, error) {
	var res string
	res = typeString + stringType
	if a.configs == nil {
		return res, nil
	}
	if val, ok := a.configs["format"]; ok { //TODO: make checking of format
		res += formatString + val.(string)
	}
	if val, ok := a.configs["minLength"]; ok { //TODO: make check?
		res += minLengthString + strconv.FormatInt(int64(val.(int)), 10)
	}
	if val, ok := a.configs["maxLength"]; ok {
		res += maxLengthString + strconv.FormatInt(int64(val.(int)), 10)
	}
	if val, ok := a.configs["pattern"]; ok {
		res += patternString + val.(string)
	}
	if val, ok := a.configs["enum"]; ok && val.([]string) != nil {
		res += enumString
		for _, enum := range val.([]string) {
			res += enumNewString + enum
		}
	}
	return res, nil
}

func (a *stringSwaggParameter) IsObject() bool {
	return false
}

func (a *stringSwaggParameter) getConfigs() map[string]interface{} {
	return a.configs
}

func NewStringSwagParameter(params map[string]interface{}) SwaggParameter {
	return &stringSwaggParameter{
		configs: params,
	}
}
