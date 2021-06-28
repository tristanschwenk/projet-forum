package mail

import (
	"reflect"

	"goyave.dev/goyave/v3/config"
)

func init() {
	config.Register("mail.host", config.Entry{
		Value:            nil,
		Type:             reflect.String,
		IsSlice:          false,
		AuthorizedValues: []interface{}{},
	})

	config.Register("mail.port", config.Entry{
		Value:            587,
		Type:             reflect.Int,
		IsSlice:          false,
		AuthorizedValues: []interface{}{},
	})

	config.Register("mail.username", config.Entry{
		Value:            nil,
		Type:             reflect.String,
		IsSlice:          false,
		AuthorizedValues: []interface{}{},
	})

	config.Register("mail.fromEmail", config.Entry{
		Value:            nil,
		Type:             reflect.String,
		IsSlice:          false,
		AuthorizedValues: []interface{}{},
	})

	config.Register("mail.password", config.Entry{
		Value:            nil,
		Type:             reflect.String,
		IsSlice:          false,
		AuthorizedValues: []interface{}{},
	})
}
