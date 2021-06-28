package user

import (
	"ezyo/forum/http/controller/user/me"
	"ezyo/forum/http/controller/user/one"

	"goyave.dev/goyave/v3"
)

func Register(rootRouter *goyave.Router) {
	router := rootRouter.Subrouter("/user")

	router.Get("/me", me.Me).Validate(me.MeRequest)
	router.Patch("/me", me.Patch).Validate(me.PathRequest)
	router.Get("/{username}", one.ByUsername)
}
