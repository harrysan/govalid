package main

import (
	"fmt"
	"testing"

	govalid "github.com/harrysan/govalid/validator"
)

type DataMap struct {
	Names []string          `validate:"slice,required,min=3"`
	Tags  map[string]string `validate:"map,keys=required;min=3,values=required;min=5"`
}

func TestValidationMap(t *testing.T) {
	data := DataMap{
		Names: []string{"John", "Does", "Game"},
		Tags:  map[string]string{"Key1": "Vl", "": "Value2"},
	}

	errors := govalid.ValidateStruct(data)
	if len(errors) > 0 {
		fmt.Println("Validation failed:", errors)
	} else {
		fmt.Println("Validation successful!")
	}
}
