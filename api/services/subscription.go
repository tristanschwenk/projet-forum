package services

import (
	"time"

	"goyave.dev/goyave/v3/database"

	"ezyo/forum/database/model"
)

func CreateSubscription(userID int, postID int) {
	db := database.Conn()
	subscription := model.Subscription{
		UserID:    userID,
		PostID:    postID,
		CreatedAt: int(time.Now().Unix()),
	}

	if result := db.Create(&subscription); result.Error != nil {
		panic(result.Error)
	}
}

func IsUserSubscribed(userID int, postID int) bool {
	db := database.Conn()
	subscription := model.Subscription{
		UserID: userID,
		PostID: postID,
	}

	var count int64

	db.Find(&subscription).Count(&count)

	if count == 1 {
		return true
	}

	return false
}

func FetchUserSubscribed(postID int) []model.User {
	db := database.Conn()
	subscriptions := []model.Subscription{}
	users := []model.User{}
	db.Joins("User").Find(&subscriptions, "postId = ?", postID)

	for _, subscription := range subscriptions {
		users = append(users, subscription.User)
	}

	return users
}
