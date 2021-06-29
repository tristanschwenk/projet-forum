package auth

import (
	"ezyo/forum/http/controller/auth/github"
	"ezyo/forum/http/controller/auth/login"
	"ezyo/forum/http/controller/auth/passwordRequest"
	"ezyo/forum/http/controller/auth/refresh"
	"ezyo/forum/http/controller/auth/register"

	"goyave.dev/goyave/v3"
)

func Register(rootRouter *goyave.Router) {
	router := rootRouter.Subrouter("/auth")

	router.Post("/register", register.Register).Validate(register.RegisterRequest)
	router.Post("/login", login.Login).Validate(login.LoginRequest)
	router.Post("/refresh", refresh.Refresh).Validate(refresh.RefreshRequest)
	router.Post("/passwordrequest", passwordRequest.PasswordRequest).Validate(passwordRequest.PasswordRequestRequest)
	router.Post("/passwordreset", passwordRequest.PasswordReset).Validate((passwordRequest.PasswordResetRequest))
	router.Post("/github", github.Register)

}
