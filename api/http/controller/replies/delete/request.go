package delete

import "goyave.dev/goyave/v3/validation"

var (
	DeleteRequest = validation.RuleSet{
		"id": {"integer"},
	}
)
