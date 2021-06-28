package replies

import (
	"ezyo/forum/http/controller/replies/all"
	"ezyo/forum/http/controller/replies/create"
	"ezyo/forum/http/controller/replies/delete"
	"ezyo/forum/http/controller/replies/one"

	"goyave.dev/goyave/v3"
)

func Register(rootRouter *goyave.Router) {
	router := rootRouter.Subrouter("/replies")

	router.Post("/create", create.Create).Validate(create.CreateRequest)

	router.Post("/one", one.One).Validate(one.OneRequest)

	router.Post("/all", all.AllPost).Validate(all.AllRequest)
	router.Get("/all", all.AllGet)
	router.Get("/{id:[0-9]+}", all.AllbyPostId)

	router.Delete("/delete", delete.Delete)
}
