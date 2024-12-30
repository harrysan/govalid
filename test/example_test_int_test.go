package main

import (
	"fmt"
	"testing"

	govalid "github.com/harrysan/govalid/validator"
	"github.com/stretchr/testify/assert"
)

type User struct {
	Name  string `validate:"required,min=3"`
	Age   int    `validate:"min=18,max=99"`
	Email string `validate:"required,email"`
}

func TestExampleInt(t *testing.T) {
	user := User{
		Name:  "John",
		Age:   17,
		Email: "passed_email@example.com",
	}

	errs := govalid.ValidateStruct(user)
	assert.NotNil(t, errs)

	if len(errs) > 0 {
		fmt.Println("Validation Errors : ")
		for _, err := range errs {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Validation Passed!")
	}
}
