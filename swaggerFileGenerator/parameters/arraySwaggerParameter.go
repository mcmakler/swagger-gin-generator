package parameters

import (
	"errors"
	"strconv"
	"strings"
)

const (
	arrayType = "array"

	minItemsString    = "\n  minItems: "
	maxItemsString    = "\n  maxItems: "
	uniqueItemsString = "\n  uniqueItems: "
	itemsString       = "\n  items: "
)

type arraySwaggerParameter struct {
	config map[string]interface{}
	items  SwaggParameter
}

var (
	ErrorNillItemsParameter = errors.New("ERROR_EMPTY_ITEMS")
)

func (a *arraySwaggerParameter) ToString() (string, error) {
	if a.items == nil {
		return "", ErrorNillItemsParameter
	}
	var res string
	res = typeString + arrayType
	if a.config != nil {
		if val, ok := a.config["minItems"]; ok {
			res += minItemsString + strconv.FormatInt(int64(val.(int)), 10)
		}
		if val, ok := a.config["maxItems"]; ok {
			res += maxItemsString + strconv.FormatInt(int64(val.(int)), 10)
		}
		if val, ok := a.config["uniqueItems"]; ok {
			res += uniqueItemsString + strconv.FormatBool(val.(bool))
		}
	}
	items, err := a.items.ToString()
	if err != nil {
		return "", err
	}
	res += itemsString + strings.ReplaceAll(items, "\n", "\n  ")
	return res, nil
}

func (a *arraySwaggerParameter) IsObject() bool {
	return false
}

func (a *arraySwaggerParameter) getConfigs() map[string]interface{} {
	return a.config
}

func NewArraySwaggerParameter(config map[string]interface{}, items SwaggParameter) SwaggParameter {
	return &arraySwaggerParameter{
		config: config,
		items:  items,
	}
}
