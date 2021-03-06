package subscribe

import (
	"errors"
	"net/http"

	"gorm.io/gorm"
	"goyave.dev/goyave/v3"
	"goyave.dev/goyave/v3/database"

	"ezyo/forum/database/model"
	"ezyo/forum/services"
)

type SubscribeRequestStruct struct {
	Id int
}

type SubscribeResponseStruct struct {
	model.Post
	IsUserSubscribed bool `json:"isUserSubscribed"`
}

func Toggle(response *goyave.Response, request *goyave.Request) {

	data := SubscribeRequestStruct{}

	if err := request.ToStruct(&data); err != nil {
		panic(err)
	}

	user := request.User.(*model.User)

	db := database.Conn()

	post := model.Post{}

	result := db.First(&post, int(data.Id))

	notFound := errors.Is(result.Error, gorm.ErrRecordNotFound)

	if result.Error != nil && !notFound {
		panic(result.Error)
	}

	response.JSON(http.StatusOK, SubscribeResponseStruct{
		Post:             post,
		IsUserSubscribed: services.ToggleSubscription(int(user.ID), int(post.ID)),
	})
}
