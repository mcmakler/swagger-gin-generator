package structures

type requiredMainConfig struct {
	version string
	title   string
}

func (c *requiredMainConfig) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"version": c.version,
		"title":   c.title,
	}
}

func NewRequiredMainConfig(version, title string) Config {
	return &requiredMainConfig{
		version: version,
		title:   title,
	}
}
