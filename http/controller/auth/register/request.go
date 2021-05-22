package register

import "goyave.dev/goyave/v3/validation"

var (
	RegisterRequest = validation.RuleSet{
		"displayName": {"required", "string", "between:1,255"},
		"userName":    {"required", "string", "between:1,30"},
		"email":       {"required", "string", "email", "between:1,255"},
		"password":    {"required", "string", "between:8,255"},
	}
)
