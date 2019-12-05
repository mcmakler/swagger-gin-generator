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

type stringSwaggerParameter struct {
	config map[string]interface{}
}

func (a *stringSwaggerParameter) ToString() (string, error) {
	var res string
	res = typeString + stringType
	if a.config == nil {
		return res, nil
	}
	if val, ok := a.config["format"]; ok { //TODO: make checking of format
		res += formatString + val.(string)
	}
	if val, ok := a.config["minLength"]; ok {
		res += minLengthString + strconv.FormatInt(int64(val.(int)), 10)
	}
	if val, ok := a.config["maxLength"]; ok {
		res += maxLengthString + strconv.FormatInt(int64(val.(int)), 10)
	}
	if val, ok := a.config["pattern"]; ok {
		res += patternString + val.(string)
	}
	if val, ok := a.config["enum"]; ok && val.([]string) != nil {
		res += enumString
		for _, enum := range val.([]string) {
			res += enumNewString + enum
		}
	}
	return res, nil
}

func (a *stringSwaggerParameter) IsObject() bool {
	return false
}

func (a *stringSwaggerParameter) getConfigs() map[string]interface{} {
	return a.config
}

func NewStringSwaggerParameter(params map[string]interface{}) SwaggParameter {
	return &stringSwaggerParameter{
		config: params,
	}
}
