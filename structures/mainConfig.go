package structures

type mainConfig struct {
	requiredMainConfig
	description string
}

func (c *mainConfig) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"version": c.version,
		"title":   c.title,
		"description": c.description,
	}
}

func NewMainConfig(version, title, description string) Config {
	return &mainConfig{
		requiredMainConfig: requiredMainConfig{
			version: version,
			title:   title,
		},
		description: description,
	}
}
