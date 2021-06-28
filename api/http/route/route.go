package route

import (
	"ezyo/forum/database/model"
	"ezyo/forum/http/controller/auth"
	"ezyo/forum/http/controller/notification"
	"ezyo/forum/http/controller/post"
	"ezyo/forum/http/controller/replies"
	"ezyo/forum/http/controller/topic"
	"ezyo/forum/http/controller/user"

	"goyave.dev/goyave/v3"
	goyaveAuth "goyave.dev/goyave/v3/auth"
	"goyave.dev/goyave/v3/cors"
	"goyave.dev/goyave/v3/middleware"
)

func Register(router *goyave.Router) {

	router.CORS(cors.Default())
	router.Middleware(middleware.Gzip(), middleware.Trim)

	// Auth Routes
	auth.Register(router)

	// Authentificated Routes
	authenticator := goyaveAuth.Middleware(&model.User{}, &goyaveAuth.JWTAuthenticator{
		ClaimName: "id",
	})

	authentificatedRouter := router.Group()
	authentificatedRouter.Middleware(authenticator)

	user.Register(authentificatedRouter)
	topic.Register(authentificatedRouter)
	post.Register(authentificatedRouter)
	replies.Register(authentificatedRouter)
	notification.Register(authentificatedRouter)
}
