package parameters

const (
	boolType = "boolean"
)

type boolSwaggParameter struct {
	configs map[string]interface{}
}

func (a *boolSwaggParameter) ToString() (string, error) {
	var res string
	res = typeString + boolType
	if a.configs == nil {
		return res, nil
	}
	if val, ok := a.configs["description"]; ok {
		res += descriptionString + val.(string)
	}
	return res, nil
}

func (a *boolSwaggParameter) IsObject() bool {
	return false
}

func (a *boolSwaggParameter) getConfigs() map[string]interface{} {
	return a.configs
}

func NewBoolSwagParameter(params map[string]interface{}) SwaggParameter {
	return &boolSwaggParameter{
		configs: params,
	}
}
