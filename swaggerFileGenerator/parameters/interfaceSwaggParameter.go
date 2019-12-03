package parameters

import "errors"

const (
	inString              = "\n  in: "
	typeString            = "\n  type: "
	typeDeficeString      = "\n- type: "
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
)

var (
	errorEmptyName = errors.New("ERROR_NAME_IS_MANDATORY")
	errorEmptyIn   = errors.New("ERROR_IN_IS_MANDATORY")
)

type SwaggParameter interface {
	ToString(isDefinition bool) (string, error)
}
