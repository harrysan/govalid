package govalid

import (
	"errors"
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
		fieldType := typ.Field(i)

		// Tag "validate"
		tag := fieldType.Tag.Get("validate")
		if tag == "" {
			continue
		}

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

	return errs
}

// applyRule => validate a field
func applyRule(fieldName string, value any, rule string) error {
	switch {
	case rule == "required":
		if validateRequired(value) {
			return errors.New("field is required")
		}
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
	default:
		return applyCustomRule(rule, fieldName, value)
	}

	return nil
}
