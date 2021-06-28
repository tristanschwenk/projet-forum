package all

import "goyave.dev/goyave/v3/validation"

var (
	AllRequest = validation.RuleSet{
		"topicId": {"required", "integer"},
	}
)
