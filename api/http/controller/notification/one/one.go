package one

import (
	"errors"
	"net/http"

	"gorm.io/gorm"
	"goyave.dev/goyave/v3"
	"goyave.dev/goyave/v3/database"

	"ezyo/forum/database/model"
)

func ById(response *goyave.Response, request *goyave.Request) {

	db := database.Conn()

	topic := model.Notification{}

	result := db.Where("id = ?", request.Params["id"]).First(&topic)

	notFound := errors.Is(result.Error, gorm.ErrRecordNotFound)

	if result.Error != nil && !notFound {
		panic(result.Error)
	}

	if !notFound {
		response.JSON(http.StatusOK, topic)
		return
	}

	response.Status(http.StatusNotFound)
}
