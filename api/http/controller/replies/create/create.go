package create

import (
	"net/http"
	"time"

	"goyave.dev/goyave/v3"
	"goyave.dev/goyave/v3/database"

	"ezyo/forum/database/model"
	"ezyo/forum/services"
)

type CreateRequestStruct struct {
	Body   string
	PostID int
}

func Create(response *goyave.Response, request *goyave.Request) {
	data := CreateRequestStruct{}
	if err := request.ToStruct(&data); err != nil {
		panic(err)
	}

	db := database.Conn()

	user := request.User.(*model.User)

	reply := model.Replies{
		Body:      data.Body,
		UserID:    int(user.ID),
		PostID:    int(data.PostID),
		CreatedAt: int(time.Now().Unix()),
	}

	post := model.Post{}

	db.First(&post, int(data.PostID))

	subscribedUsers := services.FetchUserSubscribed(int(post.ID))

	if !services.ContainsUser(subscribedUsers, user) {
		services.CreateSubscription(int(user.ID), int(post.ID))
	}

	sendNotificationTo := services.OmitUser(subscribedUsers, user)

	services.CreateNotificationBulk("New reply to {NAME}", "@XXX added a reply to {NAME}", sendNotificationTo, "topics/1/2#bob")

	if result := db.Create(&reply); result.Error != nil {
		panic(result.Error)
	}

	response.JSON(http.StatusOK, reply)
}
