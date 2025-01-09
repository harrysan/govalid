package main

import (
	"fmt"
	"testing"

	govalid "github.com/harrysan/govalid/validator"
)

type DataSlice struct {
	Names []string `validate:"slice,required,min=3"`
	Age   []int    `validate:"slice,max=30"`
	Email []string `validate:"slice,email"`
	// Tags  map[string]string `validate:"map,keys=required,values=required"`
}

func TestValidationSlice(t *testing.T) {
	data := DataSlice{
		Names: []string{"John", "Do", "JD"},
		Age:   []int{17, 31},
		Email: []string{"john@doe.com", "invalid_email"},
		// Tags:  map[string]string{"Key1": "Value1", "": "Value2"},
	}

	errors := govalid.ValidateStruct(data)
	if len(errors) > 0 {
		fmt.Println("Validation failed:", errors)
	} else {
		fmt.Println("Validation successful!")
	}
}
