package govalid

import (
	"fmt"
	"reflect"
)

func validateField(data interface{}, field reflect.StructField, rule string) error {
	value := reflect.ValueOf(data).FieldByName(field.Name)

	switch {
	case rule == "required":
		if validateRequired(value) {
			return fmt.Errorf("field is required")
		}
		// Add another rule (e.g., min, max)
	}

	return nil
}
