package validator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
)

// check if nil / empty / 0
func validateRequired(value any) bool {
	return value == nil || value == "" || value == 0
}

// validate min value
func validateMin(value any, min int) error {
	typ := reflect.TypeOf(value) // value.(int)

	if typ.Kind() == reflect.Int {
		v, _ := value.(int)
		if v < min {
			return fmt.Errorf("value must be greater than or equal to %d", min)
		}
	} else if typ.Kind() == reflect.String {
		v, _ := value.(string)
		if len(v) < min {
			return fmt.Errorf("value must be greater than or equal to %d", min)
		}
	}

	return nil
}

// validate max value
func validateMax(value any, max int) error {
	typ := reflect.TypeOf(value) // value.(int)

	if typ.Kind() == reflect.Int {
		v, _ := value.(int)
		if v > max {
			return fmt.Errorf("value must be less than or equal to %d", max)
		}
	} else if typ.Kind() == reflect.String {
		v, _ := value.(string)
		if len(v) > max {
			return fmt.Errorf("value must be less than or equal to %d", max)
		}
	}

	return nil
}

// validate format email
func validateEmail(value any) error {
	v, ok := value.(string)
	if !ok {
		return errors.New("email validation only supports strings")
	}
	re := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	if !re.MatchString(v) {
		return errors.New("invalid email format")
	}

	return nil
}
