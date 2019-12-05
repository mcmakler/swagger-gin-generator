package structures

type integerParameterConfig struct {
	basicParameterConfig
	ExclusiveMinimum bool
	ExclusiveMaximum bool
	Default          int
	Minimum          int
	Maximum          int
	MultipleOf       int
}

func (c *integerParameterConfig) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"in":               c.In,
		"name":             c.Name,
		"required":         c.Required,
		"allowEmptyValue":  c.AllowEmptyValue,
		"description":      c.Description,
		"exclusiveMinimum": c.ExclusiveMinimum,
		"exclusiveMaximum": c.ExclusiveMaximum,
		"minimum":          c.Minimum,
		"maximum":          c.Maximum,
		"default":          c.Default,
		"multipleOf":       c.MultipleOf,
	}
}

func NewIntegerParameterConfig(in, name, description string, required, allowEmptyValue bool, defaultValue, min, max, multipleOf int, exclusiveMin, exclusiveMax bool) Config {
	return &integerParameterConfig{
		basicParameterConfig: basicParameterConfig{
			requiredParameterConfig: requiredParameterConfig{
				In:   in,
				Name: name,
			},
			Required:        required,
			AllowEmptyValue: allowEmptyValue,
			Description:     description,
		},
		Default:          defaultValue,
		Minimum:          min,
		Maximum:          max,
		MultipleOf:       multipleOf,
		ExclusiveMaximum: exclusiveMax,
		ExclusiveMinimum: exclusiveMin,
	}
}
