package subscribe

import "goyave.dev/goyave/v3/validation"

var (
	ToggleRequest = validation.RuleSet{
		"id": {"required", "integer"},
	}
)
