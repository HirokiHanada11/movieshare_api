package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rikuhatano09/movieshare_api/internal/interfaces/handler"
	"github.com/rikuhatano09/movieshare_api/internal/interfaces/middleware"
)

func main() {
	engine := gin.Default()

	// Set CORS config.
	engine.Use(middleware.CORS())

	engine.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Hello MovieShare",
		})
	})
	// Public Movie endpoints
	engine.GET("/movies/random", handler.GetMovieAtRandom)
	engine.GET("/movies", handler.GetMovieList)
	engine.GET("/movies/:id", handler.GetMovieByID)
	
	csrfProtected := engine.Group("/")
	csrfProtected.Use(middleware.CSRF())
	{
		csrfProtected.GET("/auth", handler.CSRF)

		csrfProtected.POST("/auth/login", handler.LoginHandler)
		csrfProtected.POST("/auth/logout", handler.LogoutHandler)
		csrfProtected.POST("/auth/verify", handler.VerificationHandler)

		csrfProtected.PUT("/movies/:id", handler.PutMovie)

		// Private Movie endpoints
		authorized := csrfProtected.Group("/")
		authorized.Use(middleware.Authenticate())
		{
			authorized.POST("/movies", handler.CreateMovie)
		}
	}

	engine.Run(":8000")
}
