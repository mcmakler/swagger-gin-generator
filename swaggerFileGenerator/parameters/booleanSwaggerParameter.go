package parameters

const (
	boolType = "boolean"
)

type boolSwaggerParameter struct {
	config map[string]interface{}
}

func (a *boolSwaggerParameter) ToString() (string, error) {
	var res string
	res = typeString + boolType
	if a.config == nil {
		return res, nil
	}
	return res, nil
}

func (a *boolSwaggerParameter) IsObject() bool {
	return false
}

func (a *boolSwaggerParameter) getConfigs() map[string]interface{} {
	return a.config
}

func NewBoolSwaggerParameter(config map[string]interface{}) SwaggParameter {
	return &boolSwaggerParameter{
		config: config,
	}
}
