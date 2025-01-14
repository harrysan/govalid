package main

import (
	"fmt"
	"testing"

	govalid "github.com/harrysan/govalid/validator"
	"github.com/stretchr/testify/assert"
)

type UserBoolean struct {
	Name    string `validate:"required,min=3"`
	IsHappy bool   `validate:"isFalse"`
}

func TestExampleBoolean(t *testing.T) {
	user := UserBoolean{
		Name:    "John",
		IsHappy: true,
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
