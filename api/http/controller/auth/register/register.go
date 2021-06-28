package register

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"net/http"
	"time"

	"goyave.dev/goyave/v3"
	"goyave.dev/goyave/v3/database"
	"goyave.dev/goyave/v3/lang"

	"ezyo/forum/database/model"
	"ezyo/forum/mail"
)

// Controllers are files containing a collection of Handlers related to a specific feature.
// Each feature should have its own package.
//
// Controller handlers contain the business logic of your application.
// They should be concise and focused on what matters for this particular feature in your application.
// Learn more about controllers here: https://goyave.dev/guide/basics/controllers.html

// ----------------------------------------------------------------------

// SayHi is a controller handler writing "Hi!" as a response.
//
// The Response object is used to write your response.
// https://goyave.dev/guide/basics/responses.html
//
// The Request object contains all the information about the incoming request, including it's parsed body,
// query params and route parameters.
// https://goyave.dev/guide/basics/requests.html

type RegisterRequestStruct struct {
	DisplayName string
	UserName    string
	Email       string
	Password    string
}

func Register(response *goyave.Response, request *goyave.Request) {

	data := RegisterRequestStruct{}

	if err := request.ToStruct(&data); err != nil {
		panic(err)
	}

	existingUser := model.User{}
	db := database.Conn()
	result := db.Where("email = ? OR userName = ?", data.Email, data.UserName).First(&existingUser)
	exist := errors.Is(result.Error, nil)

	if exist {
		var fieldName string
		if data.Email == existingUser.Email {
			fieldName = "email"
		} else if data.UserName == existingUser.UserName {
			fieldName = "username"
		}
		response.JSON(http.StatusUnauthorized, map[string]string{"validationError": lang.Get(request.Lang, "validation.rules.unique", ":field", fieldName)})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	user := model.User{
		DisplayName:  data.DisplayName,
		UserName:     data.UserName,
		Email:        data.Email,
		Password:     string(hashedPassword),
		RegisteredAt: int(time.Now().Unix()),
	}

	if result := db.Create(&user); result.Error != nil {
		panic(result.Error)
	}

	message := &mail.Message{
		SendTo: []string{user.Email},
		Title:  "Welcome to EzyoForum",
	}

	message.FromTemplate("welcome", user)

	message.SendAsync()

	response.JSON(http.StatusOK, user)
}
