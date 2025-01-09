package govalid

import (
	"reflect"
)

func validateField(data interface{}, field reflect.StructField, rule string) error {
	value := reflect.ValueOf(data).FieldByName(field.Name)

	switch {
	case rule == "required":
		return validateRuleRequired(value)
		// Add another rule (e.g., min, max)
	}

	return nil
}
