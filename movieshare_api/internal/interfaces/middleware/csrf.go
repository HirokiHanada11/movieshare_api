package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
	adapter "github.com/gwatts/gin-adapter"
)

func CSRF() gin.HandlerFunc {
	return adapter.Wrap(
		csrf.Protect(
			[]byte(os.Getenv("CSRF_TOKEN_BASEKEY")),
			// csrf.Secure(os.Getenv("GO_ENV") == "prod"),
			// csrf.TrustedOrigins([]string{"http://localhost:3000"}),
			// csrf.Path("/auth"),
			),
		)
}