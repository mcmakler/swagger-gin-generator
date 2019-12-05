package structures

type arrayParameterConfig struct {
	basicParameterConfig
	MinItems    int
	MaxItems    int
	UniqueItems bool
}

func (c *arrayParameterConfig) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"in":              c.In,
		"name":            c.Name,
		"required":        c.Required,
		"allowEmptyValue": c.AllowEmptyValue,
		"description":     c.Description,
		"minItems":        c.MinItems,
		"maxItems":        c.MaxItems,
		"uniqueItems":     c.UniqueItems,
	}
}

func NewArrayParameterConfig(in, name, description string, required, allowEmptyValue bool, minItems, maxItems int, uniqueItems bool) Config {
	return &arrayParameterConfig{
		basicParameterConfig: basicParameterConfig{
			requiredParameterConfig: requiredParameterConfig{
				In:   in,
				Name: name,
			},
			Required:        required,
			AllowEmptyValue: allowEmptyValue,
			Description:     description,
		},
		MinItems:    minItems,
		MaxItems:    maxItems,
		UniqueItems: uniqueItems,
	}
}
