package authentication

import (
	"errors"
	"os"
	"time"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"github.com/rikuhatano09/movieshare_api/pkg/domain/contract"
)

type (
	// NewAuthClient is a client for firebase sdk
	AuthClient struct {
		Client *auth.Client
	}

	// AuthRepository defines session management methods available
	AuthActions interface {
		CreateSessionCookie(*gin.Context)(error)
		DestroySessionCookie(*gin.Context)(error)
		VerifySessionCookie(*gin.Context)(*auth.UserInfo, error)
	}
)

// NewAuthClient creates a new auth client
func NewAuthClient(context *gin.Context) (AuthActions, error) {
	app := getFirebaseConnection()
	client, err := app.Auth(context)
	if err != nil {
		return AuthClient{}, errors.New("failed to create initiate client")
	}
	return AuthClient{
		Client: client,
	}, nil
}

// CreateSessionCookie creates and set session cookie from idToken
func (authClient AuthClient) CreateSessionCookie(context *gin.Context) error {
	requestBody := contract.LoginPostRequestBody{}

	err := context.ShouldBindJSON(&requestBody)
	if err != nil {
		return err
	}
	
	// cookie expires in 30 minutes
	expiresIn := time.Minute * 30 

	// generate session cookie from IdToken 
	cookie, err := authClient.Client.SessionCookie(context, requestBody.IdToken, expiresIn)
	if err != nil {
		return err
	}

	// get the domain of the client
	stage := os.Getenv("GO_ENV")
	var domain string
	switch stage {
	case "local":
		domain = "localhost"
	case "prod":
		domain = ""
	}

	context.SetCookie(
		"session",
		cookie,
		int(expiresIn.Seconds()),
		"/",
		domain,
		true,
		true,
	)

	return nil
}

// DestroySessionCookie destroys session cookie
func (authClient AuthClient) DestroySessionCookie(context *gin.Context) error {
	cookie, err := context.Cookie("session")
	if err != nil {
		return err 
		// errors.New("not logged in, session cookie is unavailable")
	}

	// decode the cookie
	decoded, err := authClient.Client.VerifySessionCookie(context, cookie)
	if err != nil {
		return errors.New("not logged in, session cookie invalid")
	}

	// revoke refresh token
	if err:= authClient.Client.RevokeRefreshTokens(context, decoded.UID); err != nil {
		return errors.New("failed to revoke refresh token")
	}
	
	// clear session cookie
	context.SetCookie(
		"session",
		"",
		0,
		"/",
		"",
		true,
		true,
	)
	return nil
}

// VerifySessionCookie verifies session cookie attached to the request
func (authClient AuthClient) VerifySessionCookie(context *gin.Context) (*auth.UserInfo, error) {

	cookie, err := context.Cookie("session")
	if err != nil {
		return nil, errors.New("not logged in, session cookie is unavailable")
	}

	// verify and check if the session cookie is revoked
	decoded, err := authClient.Client.VerifySessionCookieAndCheckRevoked(context, cookie)
	if err != nil {
		return nil, errors.New("verification did not pass")
	}

	// retrieve user info using uid in the token
	user, err := authClient.Client.GetUser(context, decoded.UID)
	if err != nil {
		return nil, errors.New("failed to retrieve user information")
	}
	return user.UserInfo, nil 
}