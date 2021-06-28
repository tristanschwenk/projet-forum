package delete

import "goyave.dev/goyave/v3/validation"

var (
	DeleteRequest = validation.RuleSet{
		"name": {"string", "between:1,255"},
		"id":   {"integer"},
	}
)
