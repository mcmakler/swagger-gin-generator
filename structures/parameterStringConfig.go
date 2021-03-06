package structures

type stringParameterConfig struct {
	basicParameterConfig
	Format    string
	MinLength int
	MaxLength int
	Pattern   string
	Enum      []string
}

func (c *stringParameterConfig) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"in":          c.In,
		"name":        c.Name,
		"required":    c.Required,
		"description": c.Description,
		"format":      c.Format,
		"minLength":   c.MinLength,
		"maxLength":   c.MaxLength,
		"pattern":     c.Pattern,
		"enum":        c.Enum,
	}
}

func NewStringParameterConfig(in, name, description string, required bool, format string, minLength, maxLength int, pattern string, enum []string) Config {
	return &stringParameterConfig{
		basicParameterConfig: basicParameterConfig{
			requiredParameterConfig: requiredParameterConfig{
				In:   in,
				Name: name,
			},
			Required:    required,
			Description: description,
		},
		Format:    format,
		MinLength: minLength,
		MaxLength: maxLength,
		Pattern:   pattern,
		Enum:      enum,
	}
}
