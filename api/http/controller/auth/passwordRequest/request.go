package passwordRequest

import "goyave.dev/goyave/v3/validation"

var (
	PasswordRequestRequest = validation.RuleSet{
		"email": {"required", "string"},
	}
)

var (
	PasswordResetRequest = validation.RuleSet{
		"id":       {"required", "string"},
		"password": {"required", "string", "between:8,255"},
	}
)
