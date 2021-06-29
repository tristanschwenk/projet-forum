package model

import (
	"goyave.dev/goyave/v3/database"
)

func init() {
	database.RegisterModel(&Like{})
}

type Like struct {
	ID        int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:integer;" json:"id"`
	UserID    int   `gorm:"column:userId;type:integer;" json:"userId"`
	User      User  `json:"user"`
	PostID    int   `gorm:"column:postId;type:integer;" json:"postId"`
	Post      Post  `json:"post"`
	ReplyID   int   `gorm:"column:replyId;type:integer;" json:"replyId"`
	Reply     Post  `json:"reply"`
	CreatedAt int   `gorm:"column:createdAt;type:integer;" json:"createdAt"`
}

// You may need to test features interacting with your database.
// Goyave provides a handy way to generate and save records in your database: factories.
// Factories need a generator function. These functions generate a single random record.
//
// "database.Generator" is an alias for "func() interface{}"
//
// Learn more here: https://goyave.dev/guide/advanced/testing.html#database-testing

// UserGenerator generator function for the User model.
// Generate users using the following:
//  database.NewFactory(model.UserGenerator).Generate(5)
func LikeGenerator() interface{} {

	like := &Like{}
	/* user.Name = faker.Name()

	faker.SetGenerateUniqueValues(true)
	user.Email = faker.Email()
	faker.SetGenerateUniqueValues(false) */
	return like
}
