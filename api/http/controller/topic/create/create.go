package create

import (
	"errors"

	"net/http"
	"time"

	"goyave.dev/goyave/v3"
	"goyave.dev/goyave/v3/database"
	"goyave.dev/goyave/v3/lang"

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

type CreateRequestStruct struct {
	Name string
	Body string
}

func Create(response *goyave.Response, request *goyave.Request) {

	data := CreateRequestStruct{}

	if err := request.ToStruct(&data); err != nil {
		panic(err)
	}

	existingTopic := model.Topic{}
	db := database.Conn()
	result := db.Where("Name = ?", data.Name).First(&existingTopic)
	exist := errors.Is(result.Error, nil)

	if exist {
		var fieldName string
		if data.Name == existingTopic.Name {
			fieldName = "name"
		}
		response.JSON(http.StatusUnauthorized, map[string]string{"validationError": lang.Get(request.Lang, "validation.rules.unique", ":field", fieldName)})
		return
	}

	user := request.User.(*model.User)

	topic := model.Topic{
		Name:      data.Name,
		Body:      data.Body,
		UserID:    int(user.ID),
		CreatedAt: int(time.Now().Unix()),
	}

	if result := db.Create(&topic); result.Error != nil {
		panic(result.Error)
	}

	response.JSON(http.StatusOK, topic)
}
