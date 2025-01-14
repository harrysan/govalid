package main

import (
	"fmt"
	"testing"

	govalid "github.com/harrysan/govalid/validator"
	"github.com/stretchr/testify/assert"
)

type UserRegex struct {
	Username string `validate:"required,regex=username"`
	Email    string `validate:"required,regex=email"`
	Phone    string `validate:"required,regex=phone_number"`
}

func TestValidateRegexFail(t *testing.T) {
	userRegexFail := UserRegex{
		Username: "John",
		Email:    "invalid_email.com",
		Phone:    "1234",
	}

	errs := govalid.ValidateStruct(userRegexFail)
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

func TestValidateRegexPass(t *testing.T) {
	userRegexPass := UserRegex{
		Username: "Johns",
		Email:    "test@example.com",
		Phone:    "+12345678901",
	}

	errs := govalid.ValidateStruct(userRegexPass)
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
