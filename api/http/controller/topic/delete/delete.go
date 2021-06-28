package delete

import (
	"fmt"
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

type DeleteRequestStruct struct {
	Name string
	Id   int
}

func Delete(response *goyave.Response, request *goyave.Request) {

	data := DeleteRequestStruct{}

	fmt.Println(data)

	if err := request.ToStruct(&data); err != nil {
		panic(err)
	}

	db := database.Conn()

	user := request.User.(*model.User)

	toFind := model.Topic{}

	if result := db.Where("name = ? AND userID = ?", data.Name, user.ID).Or("id = ? AND userID = ?", int(data.Id), user.ID).Delete(&toFind); result.Error != nil {
		panic(result.Error)
	}

	response.JSON(http.StatusOK, toFind)
}
