package usecase

import (
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"github.com/rikuhatano09/movieshare_api/internal/infrastructure/authentication"
)

func CreateSessionCookie(context *gin.Context) error {
	authClient, err := authentication.NewAuthClient(context)
	if err != nil {
		return err
	}

	err = authClient.CreateSessionCookie(context)
	if err != nil {
		return err
	}
	return nil
}

func DestroySessionCookie(context *gin.Context) error {
	authClient, err := authentication.NewAuthClient(context)
	if err != nil {
		return err
	}

	err = authClient.DestroySessionCookie(context)
	if err != nil {
		return err
	}
	return nil
}

func VerifySessionCookie(context *gin.Context) (*auth.UserInfo, error) {
	authClient, err := authentication.NewAuthClient(context)
	if err != nil {
		return nil, err
	}

	userInfo, err := authClient.VerifySessionCookie(context)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}