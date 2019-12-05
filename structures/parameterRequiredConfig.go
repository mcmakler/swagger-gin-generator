package structures

type requiredParameterConfig struct {
	In   string `binding:"required"` //TODO: enum
	Name string `binding:"required"`
}

func (c *requiredParameterConfig) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"in":   c.In,
		"name": c.Name,
	}
}

func NewRequiredParameterConfig(in, name string) Config {
	return &requiredParameterConfig{
		In:   in,
		Name: name,
	}
}
