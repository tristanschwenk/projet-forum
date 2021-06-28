package post

import (
	"ezyo/forum/http/controller/post/all"
	"ezyo/forum/http/controller/post/create"
	"ezyo/forum/http/controller/post/delete"
	"ezyo/forum/http/controller/post/one"

	"goyave.dev/goyave/v3"
)

func Register(rootRouter *goyave.Router) {
	router := rootRouter.Subrouter("/post")

	router.Post("/create", create.Create).Validate(create.CreateRequest)

	router.Get("/{id:[0-9]+}", one.One)

	router.Get("/all", all.All)
	router.Get("/all/{id:[0-9]+}", all.AllbyTopics)
	router.Get("/all/byUser/{id:[0-9]+}", all.AllbyUser)

	router.Delete("/delete", delete.Delete).Validate(delete.DeleteRequest)
}
