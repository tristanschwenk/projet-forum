package create

import "goyave.dev/goyave/v3/validation"

var (
	CreateRequest = validation.RuleSet{
		"name":    {"required", "string", "between:1,255"},
		"body":    {"required", "string", "min:1"},
		"topicId": {"required", "integer"},
	}
)
