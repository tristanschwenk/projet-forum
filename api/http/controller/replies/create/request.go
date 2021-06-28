package create

import "goyave.dev/goyave/v3/validation"

var (
	CreateRequest = validation.RuleSet{
		"body":   {"required", "string", "min:1"},
		"postID": {"required", "integer"},
	}
)
