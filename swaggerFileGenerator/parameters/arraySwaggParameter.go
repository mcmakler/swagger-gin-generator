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

type arraySwaggParameter struct {
	configs map[string]interface{}
	items   SwaggParameter
}

var (
	ErrorNillItemsParameter = errors.New("ERROR_EMPTY_ITEMS")
)

func (a *arraySwaggParameter) ToString(isDefinition bool) (string, error) {
	if a.items == nil {
		return "", ErrorNillItemsParameter
	}
	var res string
	res = typeString + arrayType
	if val, ok := a.configs["minItems"]; ok {
		res += minItemsString + strconv.FormatInt(int64(val.(int)), 10)
	}
	if val, ok := a.configs["maxItems"]; ok {
		res += maxItemsString + strconv.FormatInt(int64(val.(int)), 10)
	}
	if val, ok := a.configs["uniqueItems"]; ok {
		res += uniqueItemsString + strconv.FormatBool(val.(bool))
	}
	items, err := a.items.ToString(true)
	if err != nil {
		return "", err
	}
	res += itemsString + strings.Replace(items, "\n", "\n  ", -1)
	return res, nil
}

func (a *arraySwaggParameter) IsObject() bool {
	return false
}

func (a *arraySwaggParameter) getConfigs() map[string]interface{} {
	return a.configs
}

func NewArraySwaggParameter(params map[string]interface{}, items SwaggParameter) SwaggParameter {
	return &arraySwaggParameter{
		configs: params,
		items:   items,
	}
}
