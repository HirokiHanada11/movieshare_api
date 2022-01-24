package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rikuhatano09/movieshare_api/internal/infrastructure/authentication"
)

func Authenticate() gin.HandlerFunc {
	return func(context *gin.Context){
		authClient, err := authentication.NewAuthClient(context)
		if err != nil {
			// throws http error and aborts
			errorAndAbort(context, err)
			return
		}

		cookie, err := context.Cookie("session")
		if err != nil {
			// throws http error and aborts
			errorAndAbort(context, err)
			return
		}

		// verify and check if the session cookie is revoked
		_, err = authClient.Client.VerifySessionCookieAndCheckRevoked(context, cookie)
		if err != nil {
			// throws http error and aborts
			errorAndAbort(context, err)
			return
		}

		context.Next()
	}
}