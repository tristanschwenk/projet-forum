package create

import (
	"net/http"
	"time"

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

type CreateRequestStruct struct {
	Name    string
	Body    string
	TopicId int
}

func Create(response *goyave.Response, request *goyave.Request) {

	data := CreateRequestStruct{}

	if err := request.ToStruct(&data); err != nil {
		panic(err)
	}

	db := database.Conn()

	user := request.User.(*model.User)

	post := model.Post{
		Name:      data.Name,
		Body:      data.Body,
		UserID:    int(user.ID),
		TopicID:   int(data.TopicId),
		CreatedAt: int(time.Now().Unix()),
	}

	if result := db.Create(&post); result.Error != nil {
		panic(result.Error)
	}

	services.CreateSubscription(int(user.ID), int(post.ID))

	response.JSON(http.StatusOK, post)
}
