package model

import (
	"goyave.dev/goyave/v3/database"
)

func init() {
	database.RegisterModel(&PasswordCode{})
}

type PasswordCode struct {
	ID        int32  `gorm:"primary_key;AUTO_INCREMENT;column:id;type:integer;" json:"id"`
	Body      string `gorm:"column:body;type:varchar;" json:"body"`
	Email     string `gorm:"column:email;type:varchar;" json:"email"`
	CreatedAt int    `gorm:"column:createdAt;type:integer;" json:"createdAt"`
	ExpireAt  int    `gorm:"column:expireAt;type:integer;" json:"expireAt"`
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
