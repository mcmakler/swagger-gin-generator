package wrapper

import (
	"github.com/mcmakler/swagger-gin-generator/wrapper/swaggerFileGenerator/parameters"
	"reflect"
	"time"
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
	return setValueByType(p.listOfparameters, p.object, false)
}

//TODO: required
//TODO: watch the case, when object is a pointer
func ConvertObjectToSwaggParameter(params map[string]interface{}, object interface{}, subObj bool) parameters.SwaggParameter {
	var typ reflect.Type
	var val reflect.Value
	if reflect.ValueOf(object).Kind() == reflect.Ptr {
		typ = reflect.ValueOf(object).Elem().Type()
		val = reflect.ValueOf(object).Elem()
		object = reflect.ValueOf(object).Elem()
	} else {
		typ = reflect.TypeOf(object)
		val = reflect.ValueOf(object)
	}

	properties := make(map[string]parameters.SwaggParameter)

	if reflect.TypeOf(object).Kind() == reflect.Struct {
		for i := 0; i < typ.NumField(); i++ {
			properties[typ.Field(i).Name] = setValueByType(nil, val.Field(i).Interface(), true)
		}
	}

	if params == nil {
		params = make(map[string]interface{})
	}
	params["name"] = typ.Name()
	res := parameters.NewObjectSwaggerParameter(params, properties, subObj)

	return res
}

func setValueByType(params map[string]interface{}, object interface{}, subObj bool) parameters.SwaggParameter {
	switch reflect.TypeOf(object).Kind() {
	case reflect.Bool:
		return parameters.NewBoolSwagParameter(params)
	case reflect.String:
		return parameters.NewStringSwagParameter(params)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return parameters.NewIntegerSwagParameter(params)
	case reflect.Float32, reflect.Float64:
		return parameters.NewNumberSwagParameter(params)
	case reflect.Array, reflect.Slice:
		return parameters.NewArraySwaggParameter(params, setValueByType(params, reflect.Zero(reflect.TypeOf(object).Elem()).Interface(), false))
	//TODO: map?
	default:
		//some unusuall cases
		if reflect.TypeOf(time.Time{}).String() == reflect.TypeOf(object).String() {
			return parameters.NewStringSwagParameter(params)
		}

		return ConvertObjectToSwaggParameter(params, object, subObj)
	}
}
