package refresh

import "goyave.dev/goyave/v3/validation"

var (
	RefreshRequest = validation.RuleSet{
		"refreshToken": {"required", "string"},
	}
)
