package structures

type basicParameterConfig struct {
	requiredParameterConfig
	Required        bool
	AllowEmptyValue bool
	Description     string
}

func (c *basicParameterConfig) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"in":              c.In,
		"name":            c.Name,
		"required":        c.Required,
		"allowEmptyValue": c.AllowEmptyValue,
		"description":     c.Description,
	}
}

func NewBasicParameterConfig(in, name, description string, required, allowEmptyValue bool) Config {
	return &basicParameterConfig{
		requiredParameterConfig: requiredParameterConfig{
			In:   in,
			Name: name,
		},
		Required:        required,
		AllowEmptyValue: allowEmptyValue,
		Description:     description,
	}
}
