package structures

type requestConfig struct {
	description string
	security    []string
	consumes    []string
	produces    []string
	tags        []string
	operationId string
	summary     string
}

func (c *requestConfig) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"description": c.description,
		"security":    c.security,
		"consumes":    c.consumes,
		"produces":    c.produces,
		"tags":        c.tags,
		"operationId": c.operationId,
		"summary":     c.summary,
	}
}

func NewRequestConfig(description, operationId, summary string, security, consumes, produces, tags []string) Config {
	return &requestConfig{
		description: description,
		security:    security,
		consumes:    consumes,
		produces:    produces,
		tags:        tags,
		operationId: operationId,
		summary:     summary,
	}
}
