package rules

import (
	"reflect"
)

func ValidateField(data interface{}, field reflect.StructField, rule string) error {
	value := reflect.ValueOf(data).FieldByName(field.Name)

	switch {
	case rule == "required":
		return ValidateRuleRequired(value)
		// Add another rule (e.g., min, max)
	}

	return nil
}
