package me

import "goyave.dev/goyave/v3/validation"

var (
	MeRequest   = validation.RuleSet{}
	PathRequest = validation.RuleSet{
		"displayName": {"required", "string", "between:1,255"},
		"email":       {"required", "string", "email", "between:1,255"},
	}
)
