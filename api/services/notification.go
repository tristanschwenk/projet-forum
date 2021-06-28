package services

import (
	"time"

	"goyave.dev/goyave/v3/database"

	"ezyo/forum/database/model"
)

func CreateNotification(name string, body string, userID int, link string) {
	db := database.Conn()
	notification := model.Notification{
		Name:      name,
		Body:      body,
		UserID:    userID,
		Link:      link,
		CreatedAt: int(time.Now().Unix()),
	}

	if result := db.Create(&notification); result.Error != nil {
		panic(result.Error)
	}
}

func CreateNotificationBulk(name string, body string, users []model.User, link string) {
	db := database.Conn()

	notifications := []model.Notification{}

	if len(users) > 0 {
		for _, user := range users {
			notifications = append(notifications, model.Notification{
				Name:      name,
				Body:      body,
				UserID:    int(user.ID),
				Link:      link,
				CreatedAt: int(time.Now().Unix()),
			})
		}

		if result := db.Create(&notifications); result.Error != nil {
			panic(result.Error)
		}
	}

}
