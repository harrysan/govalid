package main

import (
	"fmt"
	"testing"

	govalid "github.com/harrysan/govalid/validator"
	"github.com/stretchr/testify/assert"
)

func TestCustomValidation(t *testing.T) {
	err := govalid.RegisterCustomRule("isEven", func(field string, value interface{}) error {
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("%s must be an integer", field)
		}
		if v%2 != 0 {
			return fmt.Errorf("%s must be an even number", field)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Failed to register custom rule:", err)
		return
	}

	type Data struct {
		Number int `validate:"custom=isEven"`
	}

	data := Data{Number: 3} // Expect error
	errs := govalid.ValidateStruct(data)

	assert.Nil(t, errs)

	if len(errs) > 0 {
		fmt.Println("Validation Errors:")
		for _, err := range errs {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Validation Passed!")
	}
}
