package all

import (
	"net/http"

	"goyave.dev/goyave/v3"
	"goyave.dev/goyave/v3/database"

	"ezyo/forum/database/model"
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

type AllRequestStruct struct {
	PostID int
}

func AllGet(response *goyave.Response, request *goyave.Request) {

	db := database.Conn()

	toFind := []model.Replies{}

	if result := db.Find(&toFind); result.Error != nil {
		panic(result.Error)
	}
	response.JSON(http.StatusOK, toFind)
}

func AllbyPostId(response *goyave.Response, request *goyave.Request) {

	data := AllRequestStruct{}

	if err := request.ToStruct(&data); err != nil {
		panic(err)
	}

	db := database.Conn()

	toFind := []model.Replies{}

	if result := db.Where("PostID = ?", request.Params["id"]).Find(&toFind); result.Error != nil {
		panic(result.Error)
	}
	response.JSON(http.StatusOK, toFind)
}

func AllPost(response *goyave.Response, request *goyave.Request) {

	data := AllRequestStruct{}

	if err := request.ToStruct(&data); err != nil {
		panic(err)
	}

	db := database.Conn()

	toFind := []model.Replies{}

	if result := db.Where("postID = ?", data.PostID).Find(&toFind); result.Error != nil {
		panic(result.Error)
	}
	response.JSON(http.StatusOK, toFind)
}
