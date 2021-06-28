package model

import (
	"goyave.dev/goyave/v3/database"
)

func init() {
	database.RegisterModel(&Notification{})
}

type Notification struct {
	ID        int32  `gorm:"primary_key;AUTO_INCREMENT;column:id;type:integer;" json:"id"`
	Name      string `gorm:"column:name;type:varchar;" json:"name"`
	Body      string `gorm:"column:body;type:varchar;" json:"body"`
	Link      string `gorm:"column:link;type:varchar;" json:"link"`
	UserID    int    `gorm:"column:userId;type:integer;" json:"userId"`
	User      User   `json:"user"`
	CreatedAt int    `gorm:"column:createdAt;type:integer;" json:"createdAt"`
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
func NotificationGenerator() interface{} {

	notification := &Notification{}
	/* user.Name = faker.Name()

	faker.SetGenerateUniqueValues(true)
	user.Email = faker.Email()
	faker.SetGenerateUniqueValues(false) */
	return notification
}
