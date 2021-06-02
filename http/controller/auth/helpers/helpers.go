package helpers

import (
	"ezyo/forum/database/model"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	goyaveAuth "goyave.dev/goyave/v3/auth"
	"goyave.dev/goyave/v3/config"
)

type TokensStruct struct {
	Refresh string
	Access  string
}

func GenerateTokens(user model.User) (TokensStruct, error) {
	accessToken, err := GenerateAccessToken(user)
	if err != nil {
		return TokensStruct{}, err
	}

	refreshToken, err := GenerateRefreshToken(user)
	if err != nil {
		return TokensStruct{}, err
	}

	return TokensStruct{
		Access:  accessToken,
		Refresh: refreshToken,
	}, nil
}

func GenerateAccessToken(user model.User) (string, error) {
	return goyaveAuth.GenerateTokenWithClaims(jwt.MapClaims{
		"email": user.Email,
		"id":    user.ID,
	}, jwt.SigningMethodHS256)
}

func GenerateRefreshToken(user model.User) (string, error) {
	refreshExpiry := time.Duration(config.GetInt("auth.jwt.refreshExpiry")) * time.Second
	refreshSecret := []byte(config.GetString("auth.jwt.refreshSecret"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"nbf": time.Now().Unix(),
		"exp": time.Now().Add(refreshExpiry).Unix(),
	})

	return token.SignedString(refreshSecret)
}

func KeyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(config.GetString("auth.jwt.secret")), nil
}

func RefreshKeyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(config.GetString("auth.jwt.refreshSecret")), nil
}
