package one

import (
	"errors"
	"net/http"

	"gorm.io/gorm"
	"goyave.dev/goyave/v3"
	"goyave.dev/goyave/v3/database"

	"ezyo/forum/database/model"
	"ezyo/forum/services"
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

type OneRequestStruct struct {
	Name string
	Id   int
}

type OneResponseStruct struct {
	model.Post
	IsUserSubscribed bool `json:"isUserSubscribed"`
}

func One(response *goyave.Response, request *goyave.Request) {

	db := database.Conn()

	user := request.User.(*model.User)
	post := model.Post{}

	result := db.First(&post, request.Params["id"])

	notFound := errors.Is(result.Error, gorm.ErrRecordNotFound)

	if result.Error != nil && !notFound {
		panic(result.Error)
	}

	services.FetchUserSubscribed(int(post.ID))

	response.JSON(http.StatusOK, OneResponseStruct{
		Post:             post,
		IsUserSubscribed: services.IsUserSubscribed(int(user.ID), int(post.ID)),
	})
}
