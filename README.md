# GoValid

GoValid is a lightweight validation library for Golang designed to simplify data structure validation using tags in structs. It supports built-in rules like `required`, `min`, `max`, and `email`, and allows users to define custom validation rules.

---

## 🎯 Features

- ✅ Tag-based validation for structs.
- ✅ Built-in rules:
  - `required`: Ensures the field is not empty.
  - `min`: Minimum value for integers.
  - `max`: Maximum value for integers.
  - `email`: Validates email format.
- ✅ Support for custom rules.
- ✅ Easy-to-follow documentation and examples.

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
	"github.com/harrysan/govalid"
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

	errs := govalid.ValidateStruct(user)
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

### **2. Conditional Validation**

##### Conditional Validation: `validate_if`

Use `validate_if` to apply rules conditionally based on another field's value.

#### Example:

```go
type User struct {
	IsActive bool   `validate:"isTrue"`
	Reason   string `validate_if:"IsActive=true,required"`
}
```

---

### **3. Slice and Map Validation**

Use `validate` to apply rules on slice and maps.

#### Example:

```go
type DataSlice struct {
	Names []string `validate:"slice,required,min=3"`
	Age   []int    `validate:"slice,max=30"`
	Email []string `validate:"slice,email"`
}
```

```go
type DataMap struct {
	Names []string          `validate:"slice,required,min=3"`
	Tags  map[string]string `validate:"map,keys=required;min=3,values=required;min=5"`
}
```

---

### **4. Struct Validation**

##### Use `validate` to apply rules on struct.

#### Example:

```go
type Address struct {
	City    string `validate:"required"`
	ZipCode string `validate:"required,min=5"`
}

type UserStruct struct {
	Name    string  `validate:"required"`
	Age     int     `validate:"min=18"`
	Address Address `validate:"struct"`
}
```

---

### **5. Add Custom Validation Rules**

You can register custom validation rules using the `RegisterCustomRule` function.

```go
package main

import (
	"fmt"
	"github.com/harrysan/govalid"
)

func main() {
	// Register a custom rule
	err := govalid.RegisterCustomRule("isEven", func(field string, value interface{}) error {
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
	errs := govalid.ValidateStruct(data)

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
| `bool`     | The field must be true.                                     | `validate:"isTrue"`   |
| `bool`     | The field must be false.                                    | `validate:"isFalse"`  |
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
.
└── govalid/
    ├── go.mod
    ├── validator/
    │   ├── validator.go   # Core validation logic
    │   ├── rules.go       # Rules for validation
    │   ├── rules_if.go    # Rules for validation_if
    │   └── custom.go      # Custom rule support
    ├── rules/
    │   └── regex_rules.go  # Regex rules
    ├── test/
    │   ├── validate_string_test.go     # Example usage
    │   ├── validate_custom_test.go     # Example usage
    │   ├── validate_struct_test.go     # Example usage
    │   └── etc..
    └── README.md          # Documentation
```

---

## 🛠️ Roadmap

Planned features for future updates:

- Support for nested struct validation.
- Customizable error messages (including multi-language support).
- Validation for additional data types (e.g., float, time).

---

## 🤝 Contributing

Contributions are welcome! Feel free to open a pull request or report issues in the [issues](https://github.com/username/validation-lib/issues) section.

---

## 📄 License

This library is licensed under the [MIT License]().

---

Happy coding ! 😊
