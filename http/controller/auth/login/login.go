package login

import (
	"net/http"
	"time"

	"goyave.dev/goyave/v3"
	"goyave.dev/goyave/v3/database"

	"ezyo/forum/database/model"
	_ "ezyo/forum/database/model"
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

func Login(response *goyave.Response, request *goyave.Request) {

	data := LoginRequestStruct{}

	if err := request.ToStruct(&data); err != nil {
		panic(err)
	}

	user := model.User{}

	db := database.Conn()

	if result := db.Where("email = ? AND password = ?", data.Email, data.Password).Find(&user); result.Error != nil {
		panic(result.Error)
	}

	user.LastLoginAt = int(time.Now().Unix())
	db.Save(&user)

	response.JSON(http.StatusOK, user)
}
