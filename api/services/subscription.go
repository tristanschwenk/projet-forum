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

func RemoveSubscription(userID int, postID int) {
	db := database.Conn()

	if result := db.Where("userId = ? AND postId = ?", userID, postID).Delete(model.Subscription{}); result.Error != nil {
		panic(result.Error)
	}
}

func IsUserSubscribed(userID int, postID int) bool {
	db := database.Conn()

	var count int64

	db.Model(&model.Subscription{}).Where("userId = ? AND postId = ?", userID, postID).Count(&count)
	return count >= 1
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

func ToggleSubscription(userID int, postID int) bool {
	if IsUserSubscribed(userID, postID) {
		RemoveSubscription(userID, postID)
		return false
	} else {
		CreateSubscription(userID, postID)
		return true
	}
}
