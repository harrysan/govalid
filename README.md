# GoValid

GoValid is a lightweight validation library for Golang designed to simplify data structure validation using tags in structs. It supports built-in rules like `required`, `min`, `max`, and `email`, and allows users to define custom validation rules.

---

## 🎯 Features

* ✅ Tag-based validation for structs.
* ✅ Built-in rules:
  * `required`: Ensures the field is not empty.
  * `min`: Minimum value for integers.
  * `max`: Maximum value for integers.
  * `email`: Validates email format.
* ✅ Support for custom rules.
* ✅ Easy-to-follow documentation and examples.

---

## 🚀 Installation

Install the library using `go get`:

`go get github.com/harrysan/govalid`

---

## 🔧 Usage

### **1. Validate Structs with Built-In Rules**

Use the `validate` tag to define rules for your struct fields.

```go
package main

import (
	"fmt"
	"github.com/harrysan/govalid/validator"
)

type User struct {
	Name  string validate:"required,min=3"
	Age   int    validate:"min=18,max=99"
	Email string validate:"required,email"
}

func main() {
	user := User{
		Name:  "Jo",
		Age:   17,
		Email: "invalid-email",
	}  

	errs := validator.ValidateStruct(user)
	if len(errs) > 0 {
		fmt.Println("Validation Errors:")
		for _, err := range errs {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Validation Passed!")
	}
}
```

**Output:**

```go
Validation Errors:
Field 'Name' failed validation 'min=3': value must be greater than or equal to 3
Field 'Age' failed validation 'min=18': value must be greater than or equal to 18
Field 'Email' failed validation 'email': invalid email format
```

---

### **2. Add Custom Validation Rules**

You can register custom validation rules using the `RegisterCustomRule` function.

```go
package main

import (
	"fmt"
	"github.com/username/validation-lib/validator"
)

func main() {
	// Register a custom rule
	err := validator.RegisterCustomRule("isEven", func(field string, value interface{}) error {
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("%s must be an integer", field)
		}
		if v%2 != 0 {
			return fmt.Errorf("%s must be an even number", field)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Failed to register custom rule:", err)
		return
	}

	type Data struct {
		Number int validate:"isEven"
	}

	data := Data{Number: 3}
	errs := validator.ValidateStruct(data)
  
	if len(errs) > 0 {
		fmt.Println("Validation Errors:")
		for _, err := range errs {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Validation Passed!")
	}
}
```

**Output:**

```go
Validation Errors:
Field 'Number' failed validation 'isEven': Number must be an even number
```

---

## 📜 Built-In Rules

| Rule         | Description                                                 | Example Tag             |
| ------------ | ----------------------------------------------------------- | ----------------------- |
| `required` | Ensures the field is not empty.                             | `validate:"required"` |
| `min`      | The field must be greater than or equal to a minimum value. | `validate:"min=3"`    |
| `max`      | The field must be less than or equal to a maximum value.    | `validate:"max=10"`   |
| `email`    | The field must be in a valid email format.                  | `validate:"email"`    |

---

## ⚙️ API Reference

### **1. ValidateStruct**

```go
func ValidateStruct(s interface{}) []ValidationError
```

### **2. RegisterCustomRule**

```go
func RegisterCustomRule(name string, rule CustomRule) error
```

Registers a custom validation rule with a unique name and a function that implements the rule.

### **3. ValidationError**

Struct representing a validation error:

```go
type ValidationError struct {
	Field string      // Name of the field that failed validation
	Tag   string      // The validation rule that failed
	Value interface{} // The value of the field that failed validation
	Err   error       // Details about the error
}
```

---

## 📂 Project Structure

```
govalid/
├── go.mod                  		# Go module file
├── validator/
│   ├── validator.go        		# Core validation logic
│   ├── rules.go            		# Built-in validation rules
│   ├── custom.go           		# Custom validation rule support
├── test/
│   ├── example_test.go     		# Unit test for basic validation
│   ├── example_custom_test.go 		# Unit test for custom validation rules
├── README.md               		# Project documentation
```

---

## 🛠️ Roadmap

Planned features for future updates:

* Support for nested struct validation.
* Customizable error messages (including multi-language support).
* Validation for additional data types (e.g., float, time).

---

## 🤝 Contributing

Contributions are welcome! Feel free to open a pull request or report issues in the [issues](https://github.com/username/validation-lib/issues) section.

---

## 📄 License

This library is licensed under the [MIT License]().

---

Happy coding ! 😊
