package auth

import (
	"ezyo/forum/http/controller/auth/login"
	"ezyo/forum/http/controller/auth/register"

	"goyave.dev/goyave/v3"
)

func Register(rootRouter *goyave.Router) {
	rooter := rootRouter.Subrouter("/auth")

	rooter.Post("/register", register.Register).Validate(register.RegisterRequest)
	rooter.Post("/login", login.Login).Validate(login.LoginRequest)
}
