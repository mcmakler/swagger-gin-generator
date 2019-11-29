package utils

import (
	"SwaggerGin/swaggerFileGenerator/parameters"
	"reflect"
)

type Parameter interface {
	GetSwagParameter() parameters.SwaggParameter
}

type parameter struct {
	listOfparameters map[string]interface{}
	object           interface{}
}

func NewParameter(params map[string]interface{}, obj interface{}) Parameter {
	return &parameter{
		listOfparameters: params,
		object:           obj,
	}
}

func (p *parameter) GetSwagParameter() parameters.SwaggParameter {
	return setValueByType(p.listOfparameters, p.object)
}

//TODO: required
func ConvertObjectToSwaggParameter(params map[string]interface{}, object interface{}) parameters.SwaggParameter {
	e := reflect.ValueOf(&object).Elem()

	properties := make(map[string]parameters.SwaggParameter)

	for i := 0; i < e.NumField(); i++ {
		properties[e.Type().Field(i).Name] = setValueByType(nil, e.Field(i).Interface())
	}

	if params == nil {
		params = make(map[string]interface{})
	}
	params["name"] = reflect.ValueOf(&object).Type().Name()
	res := parameters.NewObjectSwaggerParameter(params, properties)

	return res
}

func setValueByType(params map[string]interface{}, object interface{}) parameters.SwaggParameter {
	switch reflect.TypeOf(&object).Kind() {
	case reflect.String:
		return parameters.NewStringSwagParameter(params)
	case reflect.Bool:
		return parameters.NewBoolSwagParameter(params)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return parameters.NewIntegerSwagParameter(params)
	case reflect.Float32, reflect.Float64:
		return parameters.NewNumberSwagParameter(params)
	case reflect.Array, reflect.Slice:
		//TODO: check is it work
		return parameters.NewArraySwaggParameter(params, setValueByType(params, reflect.Zero(reflect.TypeOf(object))))
	//TODO: map?
	default:
		return ConvertObjectToSwaggParameter(params, &object)
	}
}
