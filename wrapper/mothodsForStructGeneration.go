package wrapper

import "github.com/mcmakler/swagger-gin-generator/structures"

func NewRequiredParameterConfig(in, name string) structures.ParameterConfig {
	return structures.NewRequiredParameterConfig(in, name)
}

func NewBasicParameterConfig(in, name, description string, required, allowEmptyValue bool) structures.ParameterConfig {
	return structures.NewBasicParameterConfig(in, name, description, required, allowEmptyValue)
}

func NewArrayParameterConfig(in, name, description string, required, allowEmptyValue bool, minItems, maxItems int, uniqueItems bool) structures.ParameterConfig {
	return structures.NewArrayParameterConfig(in, name, description, required, allowEmptyValue, minItems, maxItems, uniqueItems)
}

func NewIntegerParameterConfig(in, name, description string, required, allowEmptyValue bool, defaultValue, min, max, multipleOf int, exclusiveMin, exclusiveMax bool) structures.ParameterConfig {
	return structures.NewIntegerParameterConfig(in, name, description, required, allowEmptyValue, defaultValue, min, max, multipleOf, exclusiveMin, exclusiveMax)
}

func NewNumberParameterConfig(in, name, description string, required, allowEmptyValue bool, defaultValue, min, max, multipleOf float64, exclusiveMin, exclusiveMax bool) structures.ParameterConfig {
	return structures.NewNumberParameterConfig(in, name, description, required, allowEmptyValue, defaultValue, min, max, multipleOf, exclusiveMin, exclusiveMax)
}

func NewStringParameterConfig(in, name, description string, required, allowEmptyValue bool, format string, minLength, maxLength int, pattern string, enum []string) structures.ParameterConfig {
	return structures.NewStringParameterConfig(in, name, description, required, allowEmptyValue, format, minLength, maxLength, pattern, enum)
}
