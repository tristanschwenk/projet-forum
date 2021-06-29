package all

import (
	"ezyo/forum/database/model"
	"net/http"

	"goyave.dev/goyave/v3"
	"goyave.dev/goyave/v3/database"
)

func All(response *goyave.Response, request *goyave.Request) {

	db := database.Conn()

	notifications := []model.Notification{}

	if result := db.Find(&notifications, "userId = ?", request.User.(*model.User).ID); result.Error != nil {
		panic(result.Error)
	}
	response.JSON(http.StatusOK, notifications)
}
