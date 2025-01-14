package main

import (
	"fmt"
	"testing"

	govalid "github.com/harrysan/govalid/validator"
)

type UserCustomError struct {
	Age   int    `validate:"min=18" error_message:"Age must be at least 18 (from custom)"`
	Email string `validate:"regex=email" error_message:"Invalid email format (from custom)"`
}

func TestValidateCustomError(t *testing.T) {
	user := UserCustomError{
		Age:   17, // Invalid, should trigger error_message
		Email: "", // Invalid, should trigger error_message
	}

	err := govalid.ValidateStruct(user)
	if err != nil {
		fmt.Println("Validation Error:", err)
	} else {
		fmt.Println("Validation Passed")
	}
}
