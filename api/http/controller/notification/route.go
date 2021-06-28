package notification

import (
	"ezyo/forum/http/controller/notification/all"
	"ezyo/forum/http/controller/notification/one"

	"goyave.dev/goyave/v3"
)

func Register(rootRouter *goyave.Router) {
	router := rootRouter.Subrouter("/notification")

	router.Get("/?", all.All)
	router.Get("/{id:[0-9]+}", one.ById)
}
