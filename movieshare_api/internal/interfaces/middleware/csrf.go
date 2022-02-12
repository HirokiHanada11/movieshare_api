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
			csrf.SameSite(csrf.SameSiteNoneMode),
			csrf.TrustedOrigins([]string{"http://localhost:3000", "https://mshare-web-app.vercel.app"}),
			// csrf.Path("/auth"),
			),
		)
}