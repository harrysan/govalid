package rules

// RegexRules is a map containing reusable regex validation rules
var RegexRules = map[string]string{
	"email":        `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`,
	"phone_number": `^\+?[0-9]{10,15}$`,
	"username":     `^[a-zA-Z0-9_]{3,16}$`,
	"zipcode":      `^[0-9]{5}(?:-[0-9]{4})?$`,
}
