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

	keyRules := []string{}
	valueRules := []string{}

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

				// for Map
				if fieldType.Type.Kind() == reflect.Map {
					if strings.HasPrefix(rule, "keys=") {
						keyRules = strings.Split(strings.TrimPrefix(rule, "keys="), ";")
					} else if strings.HasPrefix(rule, "values=") {
						valueRules = strings.Split(strings.TrimPrefix(rule, "values="), ";")

					}
				}

				if err != nil {
					errs = append(errs, ValidationError{
						Field: fieldType.Name,
						Tag:   rule,
						Value: field.Interface(),
						Err:   err,
					})
				}
			}

			if fieldType.Type.Kind() == reflect.Map {
				for _, key := range field.MapKeys() {
					mapValue := field.MapIndex(key).Interface()

					for _, rule := range keyRules {
						err := applyRule(fieldType.Name, key.Interface(), rule)

						if err != nil {
							errs = append(errs, ValidationError{
								Field: fieldType.Name,
								Tag:   rule,
								Value: key.Interface(), //field.Interface(),
								Err:   fmt.Errorf(key.String() + err.Error()),
							})
						}
					}

					for _, rule := range valueRules {
						err := applyRule(fieldType.Name, mapValue, rule)

						if err != nil {
							errs = append(errs, ValidationError{
								Field: fieldType.Name,
								Tag:   rule,
								Value: mapValue, //field.Interface(),
								Err:   fmt.Errorf(field.MapIndex(key).String() + err.Error()),
							})
						}
					}
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
		return validateRuleRequired(value)
	case strings.HasPrefix(rule, "min="):
		min, _ := strconv.ParseFloat(strings.TrimPrefix(rule, "min="), 64)
		return validateRuleMin(value, min)
	case strings.HasPrefix(rule, "max="):
		max, _ := strconv.ParseFloat(strings.TrimPrefix(rule, "max="), 64)
		return validateRuleMax(value, max)
	case rule == "email":
		return validateRuleEmail(value)
	case rule == "isTrue" || rule == "isFalse":
		return validateRuleBool(value, rule)
	case rule == "slice":
		return validateRuleSlice(value)
	case rule == "maps":
		return validateRuleMap(value)
	case rule == "custom":
		return applyCustomRule(rule, fieldName, value)
	}

	return nil
}
