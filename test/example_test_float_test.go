package main

import (
	"fmt"
	"testing"

	govalid "github.com/harrysan/govalid/validator"
	"github.com/stretchr/testify/assert"
)

type UserFloat struct {
	Name  string  `validate:"required,min=3"`
	Grade float32 `validate:"min=3.1,max=5.0"`
	Email string  `validate:"required,email"`
}

func TestExampleFloat(t *testing.T) {
	user := UserFloat{
		Name:  "John",
		Grade: 5.1,
		Email: "passed_email@example.com",
	}

	errs := govalid.ValidateStruct(user)
	assert.NotNil(t, errs)
	// assert.Nil(t, errs)

	if len(errs) > 0 {
		fmt.Println("Validation Errors : ")
		for _, err := range errs {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Validation Passed!")
	}
}
