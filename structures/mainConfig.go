package structures

type mainConfig struct {
	requiredMainConfig
	description string
	basePath    string
}

func (c *mainConfig) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"version":     c.version,
		"title":       c.title,
		"description": c.description,
		"basePath":    c.basePath,
	}
}

func NewMainConfig(version, title, description, basePath string) Config {
	return &mainConfig{
		requiredMainConfig: requiredMainConfig{
			version: version,
			title:   title,
		},
		description: description,
		basePath:    basePath,
	}
}
