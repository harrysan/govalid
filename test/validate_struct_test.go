package main

import (
	"fmt"
	"testing"

	govalid "github.com/harrysan/govalid/validator"
)

type Address struct {
	City    string `validate:"required"`
	ZipCode string `validate:"required,min=5"`
}

type UserStruct struct {
	Name    string  `validate:"required"`
	Age     int     `validate:"min=18"`
	Address Address `validate:"struct"`
}

func TestValidateStruct(t *testing.T) {
	user := UserStruct{
		Name: "",
		Age:  17,
		Address: Address{
			City:    "",
			ZipCode: "", // Tidak valid (harus numeric)
		},
	}

	errors := govalid.ValidateStruct(user)
	if len(errors) > 0 {
		fmt.Println("Validation failed:", errors)
	} else {
		fmt.Println("Validation successful!")
	}
}
