package me

import (
	"net/http"

	"gorm.io/gorm/clause"
	"goyave.dev/goyave/v3"
	"goyave.dev/goyave/v3/database"

	"ezyo/forum/database/model"
)

type MeResponceStruct struct {
	User *model.User `json:"user"`
}

func Me(response *goyave.Response, request *goyave.Request) {

	db := database.GetConnection()

	user := model.User{}

	db.
		Preload(clause.Associations).
		Preload("Subscription.Post").First(&user, request.User.(*model.User).ID)

	res := MeResponceStruct{
		User: &user,
	}

	response.JSON(http.StatusOK, res)
}

type PatchRequestStruct struct {
	DisplayName string
	Email       string
}

func Patch(response *goyave.Response, request *goyave.Request) {
	data := PatchRequestStruct{}

	if err := request.ToStruct(&data); err != nil {
		panic(err)
	}

	user := request.User.(*model.User)
	user.DisplayName = data.DisplayName
	user.Email = data.Email

	db := database.Conn()
	db.Save(&user)

	response.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})

}
