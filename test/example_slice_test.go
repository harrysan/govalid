package main

import (
	"fmt"
	"testing"

	govalid "github.com/harrysan/govalid/validator"
)

type DataSlice struct {
	Names []string `validate:"slice,required,min=3"`
	Age   []int    `validate:"slice,min=18"`
	// Tags  map[string]string `validate:"map,keys=required,values=required"`
}

func TestValidationSlice(t *testing.T) {
	data := DataSlice{
		Names: []string{},
		Age:   []int{17, 25},
		// Tags:  map[string]string{"Key1": "Value1", "": "Value2"},
	}

	errors := govalid.ValidateStruct(data)
	if len(errors) > 0 {
		fmt.Println("Validation failed:", errors)
	} else {
		fmt.Println("Validation successful!")
	}
}
