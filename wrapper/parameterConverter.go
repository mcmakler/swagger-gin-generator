package wrapper

import (
	"github.com/mcmakler/swagger-gin-generator/structures"
	"github.com/mcmakler/swagger-gin-generator/swaggerFileGenerator/parameters"
	"reflect"
	"strings"
	"time"
)

type Parameter interface {
	getSwaggerParameter(withSchema bool) parameters.SwaggParameter
}

type parameter struct {
	listOfParameters map[string]interface{}
	object           interface{}
}

func NewParameter(params structures.Config, obj interface{}) Parameter {
	if params != nil {
		return &parameter{
			listOfParameters: params.ToMap(),
			object:           obj,
		}
	}
	return &parameter{
		listOfParameters: nil,
		object:           obj,
	}
}

func (p *parameter) getSwaggerParameter(withSchema bool) parameters.SwaggParameter {
	obj := setValueByType(p.listOfParameters, p.object, false, false)
	if withSchema {
		return parameters.NewSchemaSwaggParameter(obj)
	}
	return obj
}

//TODO: required
//TODO: watch the case, when object is a pointer
func ConvertObjectToSwaggerParameter(params map[string]interface{}, object interface{}, subObj bool) parameters.SwaggParameter {
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
	if params == nil {
		params = make(map[string]interface{})
	}
	params["required"] = []string{}

	if reflect.TypeOf(object).Kind() == reflect.Struct {
		for i := 0; i < typ.NumField(); i++ {
			if strings.Contains(typ.Field(i).Tag.Get("binding"), "required") {
				params["required"] = append(params["required"].([]string), typ.Field(i).Name)
			}
			properties[typ.Field(i).Name] = setValueByType(nil, val.Field(i).Interface(), true, false)
		}
	}
	params["nameOfVariable"] = typ.Name()
	//params["required"] = nil //TODO: add required params
	res := parameters.NewObjectSwaggerParameter(params, properties, subObj)

	return res
}

func setValueByType(params map[string]interface{}, object interface{}, subObj bool, isNotStruct bool) parameters.SwaggParameter {
	switch reflect.TypeOf(object).Kind() {
	case reflect.Bool:
		return parameters.NewBoolSwaggerParameter(params)
	case reflect.String:
		return parameters.NewStringSwaggerParameter(params)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return parameters.NewIntegerSwaggerParameter(params)
	case reflect.Float32, reflect.Float64:
		return parameters.NewNumberSwaggerParameter(params)
	case reflect.Array, reflect.Slice:
		return parameters.NewArraySwaggerParameter(params, setValueByType(params, reflect.Zero(reflect.TypeOf(object).Elem()).Interface(), false, isNotStruct))
	//TODO: map?
	default:
		//some unusuall cases
		if reflect.TypeOf(time.Time{}).String() == reflect.TypeOf(object).String() {
			return parameters.NewStringSwaggerParameter(params)
		}

		if isNotStruct {
			return nil
		}
		return ConvertObjectToSwaggerParameter(params, object, subObj)
	}
}

func ReturnNonStructureObject(params map[string]interface{}, object interface{}) parameters.SwaggParameter {
	return setValueByType(params, object, false, true)
}
