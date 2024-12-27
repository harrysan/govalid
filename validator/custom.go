package govalid

import "errors"

type CustomRule func(field string, value any) error

var customRules = map[string]CustomRule{}

func RegisterCustomRule(name string, rule CustomRule) error {
	if _, exists := customRules[name]; exists {
		return errors.New("rule already exists: " + name)
	}
	customRules[name] = rule
	return nil
}

func applyCustomRule(ruleName string, fieldName string, value any) error {
	rule, exists := customRules[ruleName]
	if !exists {
		return errors.New("custom rule not found: " + ruleName)
	}
	return rule(fieldName, value)
}
