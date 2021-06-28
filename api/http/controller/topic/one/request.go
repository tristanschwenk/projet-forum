package one

import "goyave.dev/goyave/v3/validation"

var (
	OneRequest = validation.RuleSet{
		"name": {"string", "between:1,255"},
		"id":   {"integer"},
	}
)
