package login

import "goyave.dev/goyave/v3/validation"

var (
	LoginRequest = validation.RuleSet{
		"email":    {"required", "string", "email", "between:1,255"},
		"password": {"required", "string", "between:8,255"},
	}
)
