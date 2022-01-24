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
	engine.Use(middleware.CorsConfig())

	engine.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Hello MovieShare",
		})
	})
	// Public Movie endpoints
	engine.GET("/movies/random", handler.GetMovieAtRandom)
	engine.GET("/movies", handler.GetMovieList)
	engine.GET("/movies/:id", handler.GetMovieByID)
	engine.PUT("/movies/:id", handler.PutMovie)
	
	// Private Movie endpoints
	authorized := engine.Group("/")
	authorized.Use(middleware.Authenticate())
	{
		authorized.POST("/movies", handler.CreateMovie)
	}

	engine.POST("/auth/login", handler.LoginHandler)
	engine.POST("/auth/logout", handler.LogoutHandler)
	engine.POST("/auth/verify", handler.VerificationHandler)

	engine.Run(":8000")
}
