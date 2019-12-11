package structures

type mainConfig struct {
	requiredMainConfig
	description string
	host        string
	basePath    string
}

func (c *mainConfig) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"version":     c.version,
		"title":       c.title,
		"description": c.description,
		"host":        c.host,
		"basePath":    c.basePath,
	}
}

func NewMainConfig(version, title, description, host, basePath string) Config {
	return &mainConfig{
		requiredMainConfig: requiredMainConfig{
			version: version,
			title:   title,
		},
		description: description,
		host:        host,
		basePath:    basePath,
	}
}
