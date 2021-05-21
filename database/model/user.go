package model

import (
	"database/sql"
	"time"

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
	//[ 0] id                                             integer              null: false  primary: true   isArray: false  auto: true   col: integer         len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:integer;"`
	//[ 1] email                                          varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
	Email string `gorm:"column:email;type:varchar;"`
	//[ 2] displayName                                    varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: -1      default: []
	DisplayName string `gorm:"column:displayName;type:varchar;"`
	//[ 3] userName                                       varchar(30)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	UserName string `gorm:"column:userName;type:varchar;size:30;"`
	//[ 4] password                                       varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Password string `gorm:"column:password;type:varchar;size:255;"`
	//[ 5] lastLoginAt                                    timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	LastLoginAt time.Time `gorm:"column:lastLoginAt;type:timestamp;"`
	//[ 6] registredAt                                    timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	RegistredAt time.Time `gorm:"column:registredAt;type:timestamp;"`
	//[ 7] refreshToken                                   varchar(1000)        null: true   primary: false  isArray: false  auto: false  col: varchar         len: 1000    default: []
	RefreshToken sql.NullString `gorm:"column:refreshToken;type:varchar;size:1000;"`
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
