package route

import (
	"ezyo/forum/database/model"
	"ezyo/forum/http/controller/auth"

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
		Optional:  true,
		ClaimName: "id",
	})

	authentificatedRouter := router.Group()
	authentificatedRouter.Middleware(authenticator)

	// E.g: topic.Register(authentificatedRouter)
}
