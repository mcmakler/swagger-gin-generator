package parameters

const (
	typeString = "\n  type: "

	//For Integer and Number
	defaultValueString          = "\n  default: "
	minimumValueString          = "\n  minimum: "
	exclusiveMinimumValueString = "\n  exclusiveMinimum: "
	maximumValueString          = "\n  maximum: "
	exclusiveMaximumValueString = "\n  exclusiveMaximum: "
	multipleOfString            = "\n  multipleOf: "
)

type SwaggParameter interface {
	ToString() (string, error)
	IsObject() bool
	getConfigs() map[string]interface{}
}
