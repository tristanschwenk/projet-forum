package login

import (
	"errors"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"goyave.dev/goyave/v3"
	"goyave.dev/goyave/v3/database"
	"goyave.dev/goyave/v3/lang"

	"ezyo/forum/database/model"
	"ezyo/forum/http/controller/auth/helpers"
)

// Controllers are files containing a collection of Handlers related to a specific feature.
// Each feature should have its own package.
//
// Controller handlers contain the business logic of your application.
// They should be concise and focused on what matters for this particular feature in your application.
// Learn more about controllers here: https://goyave.dev/guide/basics/controllers.html

// ----------------------------------------------------------------------

// SayHi is a controller handler writing "Hi!" as a response.
//
// The Response object is used to write your response.
// https://goyave.dev/guide/basics/responses.html
//
// The Request object contains all the information about the incoming request, including it's parsed body,
// query params and route parameters.
// https://goyave.dev/guide/basics/requests.html

type LoginRequestStruct struct {
	Email    string
	Password string
}

type LoginResponseStruct struct {
	User         model.User `json:"user"`
	AccessToken  string     `json:"accessToken"`
	RefreshToken string     `json:"refreshToken"`
}

func Login(response *goyave.Response, request *goyave.Request) {

	// Retreive body as a struct
	data := LoginRequestStruct{}

	if err := request.ToStruct(&data); err != nil {
		panic(err)
	}

	user := model.User{}

	db := database.Conn()

	// Search the user by its email
	result := db.Where("email = ?", data.Email).First(&user)

	notFound := errors.Is(result.Error, gorm.ErrRecordNotFound)

	if result.Error != nil && !notFound {
		panic(result.Error)
	}

	if !notFound && bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)) == nil {

		tokens, err := helpers.GenerateTokens(user)

		if err != nil {
			goyave.Logger.Fatal(err)
		}

		response.JSON(http.StatusOK, LoginResponseStruct{
			User:         user,
			AccessToken:  tokens.Access,
			RefreshToken: tokens.Refresh,
		})

		// Update the user in db
		user.LastLoginAt = int(time.Now().Unix())
		user.RefreshToken = tokens.Refresh
		db.Save(&user)

		return
	}

	response.JSON(http.StatusUnauthorized, map[string]string{"validationError": lang.Get(request.Lang, "auth.invalid-credentials")})
}
