package govalid

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type ValidationError struct {
	Field string
	Tag   string
	Value interface{}
	Err   error
}

func (ve ValidationError) Error() string {
	return fmt.Sprintf("Field '%s' failed validation '%s' : %v", ve.Field, ve.Tag, ve.Err)
}

// ValidateStruct validate struct based on tag
func ValidateStruct(s any) []ValidationError {
	var errs []ValidationError

	val := reflect.ValueOf(s)
	typ := reflect.TypeOf(s)

	if val.Kind() != reflect.Struct {
		panic("Validate: input must be a struct")
	}

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}

	// Iterate field
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		field2 := val.Type().Field(i)
		fieldType := typ.Field(i)

		// Tag "validate"
		tag := fieldType.Tag.Get("validate")
		if tag != "" {
			// Split tag
			rules := strings.Split(tag, ",")
			for _, rule := range rules {
				err := applyRule(fieldType.Name, field.Interface(), rule)
				if err != nil {
					errs = append(errs, ValidationError{
						Field: fieldType.Name,
						Tag:   rule,
						Value: field.Interface(),
						Err:   err,
					})
				}
			}
		}

		// Tag "validate If"
		tagVIf := fieldType.Tag.Get("validate_if")
		if tagVIf != "" {
			parts := strings.SplitN(strings.TrimPrefix(tagVIf, "validate_if:"), ",", 2)
			if len(parts) != 2 {
				panic("Invalid validate_if format. Expected 'Field=Value,rule'")
			}
			condition := parts[0]
			additionalRule := parts[1]

			// Parse condition (e.g., IsActive=true)
			condParts := strings.SplitN(condition, "=", 2)
			if len(condParts) != 2 {
				panic("Invalid condition format in validate_if. Expected 'Field=Value'")
			}
			condField, condValue := condParts[0], condParts[1]

			// Check condition field
			condFieldValue := reflect.ValueOf(s).FieldByName(condField)
			if !condFieldValue.IsValid() {
				panic("Condition field '" + condField + "' not found")
			}

			// Convert condition value to string for comparison
			condFieldValueStr := fmt.Sprintf("%v", condFieldValue.Interface())
			if condFieldValueStr == condValue {
				// Apply additional validation rule if condition is met
				validateField(s, field2, additionalRule)
			}
			// Check 2nd condition (e.g., required)
			if additionalRule != "" {
				err := applyRule(fieldType.Name, field.Interface(), additionalRule)
				if err != nil {
					errs = append(errs, ValidationError{
						Field: fieldType.Name,
						Tag:   additionalRule,
						Value: field.Interface(),
						Err:   err,
					})
				}
			}
		}
	}

	return errs
}

// applyRule => validate a field
func applyRule(fieldName string, value any, rule string) error {
	switch {
	case rule == "required":
		return validateRequired(value)
	case strings.HasPrefix(rule, "min="):
		min, _ := strconv.ParseFloat(strings.TrimPrefix(rule, "min="), 64)
		return validateMin(value, min)
	case strings.HasPrefix(rule, "max="):
		max, _ := strconv.ParseFloat(strings.TrimPrefix(rule, "max="), 64)
		return validateMax(value, max)
	case rule == "email":
		return validateEmail(value)
	case rule == "isTrue" || rule == "isFalse":
		return validateBool(value, rule)
	case rule == "slice":
		return validateSlice(value)
	default:
		return applyCustomRule(rule, fieldName, value)
	}
}

// TODO : validate slice nested = example slice test . go
