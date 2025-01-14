package rules

import (
	"errors"
	"sync"
)

// RegexRules is a map containing reusable regex validation rules
var regexRules = struct {
	sync.RWMutex
	m map[string]string
}{
	m: map[string]string{
		"email":        `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`,
		"phone_number": `^\+?[0-9]{10,15}$`,
		"username":     `^[a-zA-Z0-9_]{3,16}$`,
		"zipcode":      `^[0-9]{5}(?:-[0-9]{4})?$`,
		"url":          `/https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#()?&//=]*)/`,
		"ipv4":         `/^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$/`,
		"ipv6":         `/(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))/`,
		"slug":         `/^[a-z0-9]+(?:-[a-z0-9]+)*$/`,
	},
}

// GetRegexRule retrieves a regex rule by name
func GetRegexRule(name string) (string, error) {
	regexRules.RLock()
	defer regexRules.RUnlock()

	rule, exists := regexRules.m[name]
	if !exists {
		return "", errors.New("regex rule not found")
	}
	return rule, nil
}

// AddOrUpdateRegexRule adds or updates a regex rule
func AddOrUpdateRegexRule(name, pattern string) {
	regexRules.Lock()
	defer regexRules.Unlock()

	regexRules.m[name] = pattern
}

// DeleteRegexRule removes a regex rule by name
func DeleteRegexRule(name string) error {
	regexRules.Lock()
	defer regexRules.Unlock()

	if _, exists := regexRules.m[name]; !exists {
		return errors.New("regex rule not found")
	}
	delete(regexRules.m, name)
	return nil
}
