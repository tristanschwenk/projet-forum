package services

import (
	"ezyo/forum/database/model"
)

func OmitUser(users []model.User, user *model.User) []model.User {
	i := 0
	for _, x := range users {
		if x.ID != user.ID {
			users[i] = x
			i++
		}
	}
	users = users[:i]
	return users
}

func ContainsUser(users []model.User, user *model.User) bool {
	for _, x := range users {
		if x.ID == user.ID {
			return true
		}
	}
	return false
}
