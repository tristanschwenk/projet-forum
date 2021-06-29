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
	ID int
}

type SubscribeResponseStruct struct {
	model.Post
	IsUserSubscribed bool `json:"isUserSubscribed"`
}

func Subscribe(response *goyave.Response, request *goyave.Request) {

	data := SubscribeRequestStruct{}

	if err := request.ToStruct(&data); err != nil {
		panic(err)
	}

	user := request.User.(*model.User)

	db := database.Conn()

	post := model.Post{}

	result := db.First(&post, data.ID)

	notFound := errors.Is(result.Error, gorm.ErrRecordNotFound)

	if result.Error != nil && !notFound {
		panic(result.Error)
	}

	response.JSON(http.StatusOK, SubscribeResponseStruct{
		Post:             post,
		IsUserSubscribed: services.ToggleSubscription(int(user.ID), int(post.ID)),
	})

	response.JSON(http.StatusOK, post)
}
