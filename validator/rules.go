package govalid

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
)

type TypeParam interface {
	int | int32 | int64 | float32 | float64
}

// check if nil / empty / 0
func validateRequired(value any) error {
	typ := reflect.TypeOf(value)
	val := reflect.ValueOf(value)

	if typ.Kind() == reflect.Slice {
		ret := make([]interface{}, val.Len())
		if len(ret) == 0 {
			return fmt.Errorf("field is required")
		}
	} else {
		if value == nil || value == "" || value == 0 {
			return fmt.Errorf("field is required")
		}
	}

	return nil
}

func validateBool(value any, rule string) error {
	typ := reflect.TypeOf(value)
	val := reflect.ValueOf(value)

	switch rule {
	case "isTrue":
		if typ.Kind() == reflect.Bool && !val.Bool() {
			return fmt.Errorf("value must be true")
		}
	case "isFalse":
		if typ.Kind() == reflect.Bool && val.Bool() {
			return fmt.Errorf("value must be false")
		}
	}

	return nil
}

// validate min value
func validateMin[T TypeParam](value any, min T) error {
	typ := reflect.TypeOf(value) // value.(int)

	if typ.Kind() == reflect.Int {
		v, _ := value.(int)

		if v < int(min) {
			return fmt.Errorf("value must be greater than or equal to %d", int(min))
		}
	} else if typ.Kind() == reflect.Int32 {
		v, _ := value.(int32)

		if v < int32(min) {
			return fmt.Errorf("value must be greater than or equal to %d", int32(min))
		}
	} else if typ.Kind() == reflect.Int64 {
		v, _ := value.(int64)

		if v < int64(min) {
			return fmt.Errorf("value must be greater than or equal to %d", int64(min))
		}
	} else if typ.Kind() == reflect.String {
		v, _ := value.(string)

		if len(v) < int(min) {
			return fmt.Errorf("value must be greater than or equal to %d", int(min))
		}
	} else if typ.Kind() == reflect.Float64 {
		v, _ := value.(float64)

		if v < float64(min) {
			return fmt.Errorf("value must be greater than or equal to %f", float64(min))
		}
	} else if typ.Kind() == reflect.Float32 {
		v, _ := value.(float32)

		if v < float32(min) {
			return fmt.Errorf("value must be greater than or equal to %f", float32(min))
		}
	}

	return nil
}

// validate max value
func validateMax[T TypeParam](value any, max T) error {
	typ := reflect.TypeOf(value) // value.(int)

	if typ.Kind() == reflect.Int {
		v, _ := value.(int)

		if v > int(max) {
			return fmt.Errorf("value must be less than or equal to %d", int(max))
		}
	} else if typ.Kind() == reflect.Int32 {
		v, _ := value.(int32)

		if v > int32(max) {
			return fmt.Errorf("value must be less than or equal to %d", int32(max))
		}
	} else if typ.Kind() == reflect.Int64 {
		v, _ := value.(int64)

		if v > int64(max) {
			return fmt.Errorf("value must be less than or equal to %d", int64(max))
		}
	} else if typ.Kind() == reflect.String {
		v, _ := value.(string)

		if len(v) > int(max) {
			return fmt.Errorf("value must be less than or equal to %d", int(max))
		}
	} else if typ.Kind() == reflect.Float32 {
		v, _ := value.(float32)

		if v > float32(max) {
			return fmt.Errorf("value must be less than or equal to %.1f", float32(max))
		}
	} else if typ.Kind() == reflect.Float64 {
		v, _ := value.(float64)

		if v > float64(max) {
			return fmt.Errorf("value must be less than or equal to %.1f", float64(max))
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

func validateSlice(value any) error {
	typ := reflect.TypeOf(value)

	if typ.Kind() != reflect.Slice {
		return fmt.Errorf("value must be slice")
	}

	return nil
}
