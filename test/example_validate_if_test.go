package main

import (
	"fmt"
	"testing"

	govalid "github.com/harrysan/govalid/validator"
)

type UserVIf struct {
	IsActive bool   `validate:"isTrue"`
	Reason   string `validate_if:"IsActive=true,required"`
}

func TestValidateIfFail(t *testing.T) {
	user := UserVIf{
		IsActive: true,
		Reason:   "",
	}

	err := govalid.ValidateStruct(user)
	if len(err) > 0 {
		fmt.Println("Validation failed:", err)
	} else {
		fmt.Println("Validation successful!")
	}
}

func TestValidateIfTrue(t *testing.T) {
	user := UserVIf{
		IsActive: true,
		Reason:   "Reason",
	}

	err := govalid.ValidateStruct(user)
	if len(err) > 0 {
		fmt.Println("Validation failed:", err)
	} else {
		fmt.Println("Validation successful!")
	}
}

func TestValidateIfMany(t *testing.T) {
	tests := []struct {
		user   UserVIf
		hasErr bool
	}{
		{UserVIf{IsActive: true, Reason: ""}, true},            // Error karena Reason kosong
		{UserVIf{IsActive: true, Reason: "Active"}, false},     // Valid
		{UserVIf{IsActive: false, Reason: ""}, false},          // Tidak ada validasi karena IsActive=false
		{UserVIf{IsActive: false, Reason: "Not Active"}, true}, // Tidak ada validasi karena IsActive=false
	}

	for _, tt := range tests {
		errs := govalid.ValidateStruct(tt.user)
		fmt.Println(errs)
		if (len(errs) > 0) != tt.hasErr {
			// 	fmt.Println(errs)
			fmt.Printf("Test failed for input %+v. Expected hasErr=%v, got %v.\n", tt.user, tt.hasErr, len(errs) > 0)
		}
		fmt.Println("============================")
	}
}
