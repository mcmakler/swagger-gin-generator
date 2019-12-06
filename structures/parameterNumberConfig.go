package structures

type numberParameterConfig struct {
	basicParameterConfig
	ExclusiveMinimum bool
	ExclusiveMaximum bool
	Default          float64
	Minimum          float64
	Maximum          float64
	MultipleOf       float64
}

func (c *numberParameterConfig) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"in":               c.In,
		"name":             c.Name,
		"required":         c.Required,
		"description":      c.Description,
		"exclusiveMinimum": c.ExclusiveMinimum,
		"exclusiveMaximum": c.ExclusiveMaximum,
		"minimum":          c.Minimum,
		"maximum":          c.Maximum,
		"default":          c.Default,
		"multipleOf":       c.MultipleOf,
	}
}

func NewNumberParameterConfig(in, name, description string, required bool, defaultValue, min, max, multipleOf float64, exclusiveMin, exclusiveMax bool) Config {
	return &numberParameterConfig{
		basicParameterConfig: basicParameterConfig{
			requiredParameterConfig: requiredParameterConfig{
				In:   in,
				Name: name,
			},
			Required:    required,
			Description: description,
		},
		Default:          defaultValue,
		Minimum:          min,
		Maximum:          max,
		MultipleOf:       multipleOf,
		ExclusiveMaximum: exclusiveMax,
		ExclusiveMinimum: exclusiveMin,
	}
}
