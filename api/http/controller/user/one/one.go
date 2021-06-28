package one

import (
	"errors"
	"net/http"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"goyave.dev/goyave/v3"
	"goyave.dev/goyave/v3/database"

	"ezyo/forum/database/model"
)

func ByUsername(response *goyave.Response, request *goyave.Request) {

	db := database.Conn()

	user := model.User{}

	result := db.
		Preload(clause.Associations).
		Preload("Subscription.Post").
		Where(&model.User{UserName: request.Params["username"]}).
		First(&user)

	notFound := errors.Is(result.Error, gorm.ErrRecordNotFound)

	if result.Error != nil && !notFound {
		panic(result.Error)
	}

	if !notFound {
		response.JSON(http.StatusOK, user)
		return
	}

	response.Status(http.StatusNotFound)
}
