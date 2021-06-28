package passwordRequest

import (
	"errors"
	"ezyo/forum/database/model"
	"ezyo/forum/mail"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"goyave.dev/goyave/v3"
	"goyave.dev/goyave/v3/database"
)

type PasswordRequestStruct struct {
	Email string
}

type PasswordRequestMail struct {
	Email string
	Code  string
}

type PasswordResetStruct struct {
	Id       string
	Password string
}

func PasswordRequest(response *goyave.Response, request *goyave.Request) {

	//Fetch request
	data := PasswordRequestStruct{}

	if err := request.ToStruct(&data); err != nil {
		panic(err)
	}

	fmt.Println(data)

	//Check if email requested exist
	existingUser := model.User{}
	db := database.Conn()
	result := db.Where("email = ?", data.Email).First(&existingUser)
	fmt.Println(result)

	if result.Error != nil {
		response.JSON(http.StatusOK, "")
		return
	}

	//Generate code
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
		"abcdefghijklmnopqrstuvwxyzåäö" +
		"0123456789")
	length := 60
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()

	fmt.Println(str)

	code := model.PasswordCode{
		Body:      str,
		Email:     data.Email,
		CreatedAt: int(time.Now().Unix()),
		ExpireAt:  int(time.Now().Unix() + 600),
	}

	if result := db.Create(&code); result.Error != nil {
		panic(result.Error)
	}

	message := &mail.Message{
		SendTo: []string{data.Email},
		Title:  "Here is your new password !",
	}

	mailContent := PasswordRequestMail{Email: data.Email, Code: str}

	err2 := message.FromTemplate("passwordreset", mailContent)
	fmt.Println(err2)

	message.SendAsync()

	response.JSON(http.StatusOK, "")

}

func PasswordReset(response *goyave.Response, request *goyave.Request) {

	data := PasswordResetStruct{}

	if err := request.ToStruct(&data); err != nil {
		panic(err)
	}

	fmt.Println(data)

	db := database.Conn()

	existingCode := model.PasswordCode{}
	result := db.Where("body = ?", data.Id).First(&existingCode)
	exist := errors.Is(result.Error, nil)

	if exist {
		fmt.Println(existingCode.ExpireAt)

		//Check still valid
		if int(time.Now().Unix()) > existingCode.ExpireAt {
			response.JSON(http.StatusRequestTimeout, map[string]string{"validationError": "Link has expire"})
			return
		}

		user := model.User{}
		db.Where("email = ?", existingCode.Email).First(&user)
		fmt.Println(user)
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}

		user.Password = string(hashedPassword)

		db.Save(&user)

		response.JSON(http.StatusOK, map[string]interface{}{
			"user": user,
		})
		return
	}

	response.JSON(http.StatusBadRequest, map[string]string{"validationError": "Cannot find link"})

}
