package login

import (
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"goyave.dev/goyave/v3"
	goyaveAuth "goyave.dev/goyave/v3/auth"
	"goyave.dev/goyave/v3/config"
	"goyave.dev/goyave/v3/database"
	"goyave.dev/goyave/v3/lang"

	"ezyo/forum/database/model"
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

type LoginRequestStruct struct {
	Email    string
	Password string
}

type LoginResponseStruct struct {
	model.User
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func generateAccessToken(user model.User) (string, error) {
	return goyaveAuth.GenerateTokenWithClaims(jwt.MapClaims{
		"email": user.Email,
		"id":    user.ID,
	}, jwt.SigningMethodHS256)
}

func generateRefreshToken(user model.User) (string, error) {
	refreshExpiry := time.Duration(config.GetInt("auth.jwt.refreshExpiry")) * time.Second
	refreshSecret := []byte(config.GetString("auth.jwt.refreshSecret"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"nbf": time.Now().Unix(),
		"exp": time.Now().Add(refreshExpiry).Unix(),
	})

	return token.SignedString(refreshSecret)
}

func Login(response *goyave.Response, request *goyave.Request) {

	// Retreive body as a struct
	data := LoginRequestStruct{}

	if err := request.ToStruct(&data); err != nil {
		panic(err)
	}

	user := model.User{}

	db := database.Conn()

	// Search the user by its email
	result := db.Where("email = ?", data.Email).First(&user)

	notFound := errors.Is(result.Error, gorm.ErrRecordNotFound)

	if result.Error != nil && !notFound {
		panic(result.Error)
	}

	if !notFound && bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)) == nil {

		accessToken, err := generateAccessToken(user)

		if err != nil {
			goyave.Logger.Fatal(err)
		}

		refreshToken, err := generateRefreshToken(user)

		if err != nil {
			goyave.Logger.Fatal(err)
		}

		response.JSON(http.StatusOK, LoginResponseStruct{
			User:         user,
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		})

		// Update the user in db
		user.LastLoginAt = int(time.Now().Unix())
		user.RefreshToken = refreshToken
		db.Save(&user)

		return
	}

	response.JSON(http.StatusUnauthorized, map[string]string{"validationError": lang.Get(request.Lang, "auth.invalid-credentials")})
}
