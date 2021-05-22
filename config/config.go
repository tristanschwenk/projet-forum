package config

import (
	"reflect"

	"goyave.dev/goyave/v3/config"
)

func init() {
	config.Register("auth.jwt.refreshExpiry", config.Entry{
		Value:            86400,
		Type:             reflect.Int,
		IsSlice:          false,
		AuthorizedValues: []interface{}{},
	})

	config.Register("auth.jwt.refreshSecret", config.Entry{
		Value:            nil,
		Type:             reflect.String,
		IsSlice:          false,
		AuthorizedValues: []interface{}{},
	})
}
