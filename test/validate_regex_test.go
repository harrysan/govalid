package main

import (
	"fmt"
	"testing"

	"github.com/harrysan/govalid/rules"
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

func TestValidateCustomRegex(t *testing.T) {
	// Override phone number regex to allow only Indonesian numbers
	rules.AddOrUpdateRegexRule("phone_number", `^\+62[0-9]{9,13}$`)

	user := UserRegex{
		Username: "Pepsi",
		Email:    "pepsi@man.com",
		Phone:    "+628123456789", // Valid for Indonesian number
	}

	err := govalid.ValidateStruct(user)
	fmt.Println("Validation 1")
	if err != nil {
		fmt.Println("Validation Error:", err)
	} else {
		fmt.Println("Validation Passed")
	}

	// Test invalid phone number
	user.Phone = "+1234567890" // Not valid for Indonesian number
	err = govalid.ValidateStruct(user)
	fmt.Println("Validation 2")
	if err != nil {
		fmt.Println("Validation Error:", err)
	} else {
		fmt.Println("Validation Passed")
	}
}
