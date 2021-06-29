package github

import (
	"errors"
	"ezyo/forum/database/model"
	"ezyo/forum/http/controller/auth/helpers"
	"ezyo/forum/mail"
	"net/http"
	"time"

	"goyave.dev/goyave/v3"
	"goyave.dev/goyave/v3/database"
)

type GithubRequestStruct struct {
	Token      string
	Login      string
	Avatar_url string
	Name       string
	Email      string
}

type GithubResponseStruct struct {
	User         model.User `json:"user"`
	AccessToken  string     `json:"accessToken"`
	RefreshToken string     `json:"refreshToken"`
}

func Register(response *goyave.Response, request *goyave.Request) {

	data := GithubRequestStruct{}

	if err := request.ToStruct(&data); err != nil {
		panic(err)
	}

	goyave.Logger.Println(data)
	existingUser := model.User{}
	db := database.Conn()
	result := db.Where("email = ?", data.Email).First(&existingUser)
	exist := errors.Is(result.Error, nil)
	var user model.User
	if exist {
		user = existingUser
	} else {
		user = model.User{
			DisplayName:  data.Name,
			UserName:     data.Login,
			Email:        data.Email,
			Avatar:       data.Avatar_url,
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
	}

	tokens, err := helpers.GenerateTokens(user)

	if err != nil {
		goyave.Logger.Fatal(err)
	}

	user.Avatar = data.Avatar_url

	response.JSON(http.StatusOK, GithubResponseStruct{
		User:         user,
		AccessToken:  tokens.Access,
		RefreshToken: tokens.Refresh,
	})

	// Update the user in db
	user.LastLoginAt = int(time.Now().Unix())
	user.RefreshToken = tokens.Refresh
	db.Save(&user)
}
