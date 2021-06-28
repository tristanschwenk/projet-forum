package one

import "goyave.dev/goyave/v3/validation"

var (
	OneRequest = validation.RuleSet{
		"id": {"integer"},
	}
)
