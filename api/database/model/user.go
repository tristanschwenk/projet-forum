package model

import (
	"goyave.dev/goyave/v3/database"
)

// A model is a structure reflecting a database table structure. An instance of a model
// is a single database record. Each model is defined in its own file inside the database/models directory.
// Models are usually just normal Golang structs, basic Go types, or pointers of them.
// "sql.Scanner" and "driver.Valuer" interfaces are also supported.

// Learn more here: https://goyave.dev/guide/basics/database.html#models

func init() {
	// All models should be registered in an "init()" function inside their model file.
	database.RegisterModel(&User{})
}

type User struct {
	ID           int32          `gorm:"primary_key;AUTO_INCREMENT;column:id;type:integer;" json:"id" auth:"username"`
	Email        string         `gorm:"column:email;type:varchar;" json:"email"`
	DisplayName  string         `gorm:"column:displayName;type:varchar;" json:"displayName"`
	UserName     string         `gorm:"column:userName;type:varchar;size:30;" json:"userName"`
	Password     string         `gorm:"column:password;type:varchar;size:255;" json:"-"`
	LastLoginAt  int            `gorm:"column:lastLoginAt;type:integer;" json:"lastLoginAt"`
	RegisteredAt int            `gorm:"column:registeredAt;type:integer;" json:"registeredAt"`
	RefreshToken string         `gorm:"column:refreshToken;type:varchar;size:1000;" json:"-"`
	Topics       []Topic        `json:"topics"`
	Notification []Notification `json:"notification"`
	Post         []Post         `json:"post"`
	Replies      []Replies      `json:"replies"`
	Subscription []Subscription `json:"subscription"`
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
func UserGenerator() interface{} {

	user := &User{}
	/* user.Name = faker.Name()

	faker.SetGenerateUniqueValues(true)
	user.Email = faker.Email()
	faker.SetGenerateUniqueValues(false) */
	return user
}
