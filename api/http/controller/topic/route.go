package topic

import (
	"ezyo/forum/http/controller/topic/all"
	"ezyo/forum/http/controller/topic/create"
	"ezyo/forum/http/controller/topic/delete"
	"ezyo/forum/http/controller/topic/one"

	"goyave.dev/goyave/v3"
)

func Register(rootRouter *goyave.Router) {
	router := rootRouter.Subrouter("/topic")

	router.Get("/?", all.All)
	router.Get("/{id:[0-9]+}", one.ById)
	router.Get("/{name}", one.ByName)

	router.Post("/create", create.Create).Validate(create.CreateRequest)

	router.Delete("/delete", delete.Delete).Validate(delete.DeleteRequest)
}
