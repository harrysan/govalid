package main

import (
	"fmt"
	"testing"

	govalid "govalid/validator"
)

type User struct {
	Name  string `validate:"required,min=3"`
	Age   int    `validate:"min=18,max=99"`
	Email string `validate:"required,email"`
}

func TestExample(t *testing.T) {
	user := User{
		Name:  "Jo",
		Age:   17,
		Email: "invalid_email",
	}

	errs := govalid.ValidateStruct(user)
	if len(errs) > 0 {
		fmt.Println("Validation Errors : ")
		for _, err := range errs {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Validation Passed!")
	}
}
