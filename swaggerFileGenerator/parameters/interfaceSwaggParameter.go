package parameters

const (
	inString              = "\n  in: "
	typeString            = "\n- type: "
	nameString            = "\n  name: "
	requiredString        = "\n  required: "
	allowEmptyValueString = "\n  allowEmptyValue: "
	descriptionString     = "\n  description: "

	defaultValueString          = "\n  default: "
	minimumValueString          = "\n  minimum: "
	exclusiveMinimumValueString = "\n  exclusiveMinimum: "
	maximumValueString          = "\n  maximum: "
	exclusiveMaximumValueString = "\n  exclusiveMaximum: "
	multipleOfString            = "\n  multipleOf: "

	errorNilInParameter = "ERROR_EMPTY_IN"
)

type SwaggParameter interface {
	ToString() (string, error)
}
