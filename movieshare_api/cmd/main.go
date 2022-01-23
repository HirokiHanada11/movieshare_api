package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/rikuhatano09/movieshare_api/internal/interfaces/handler"
)

func main() {
	engine := gin.Default()

	// Set CORS config.
	engine.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge: 24 * time.Hour,
	}))

	engine.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Hello MovieShare",
		})
	})
	engine.GET("/movies/random", handler.GetMovieAtRandom)
	engine.GET("/movies", handler.GetMovieList)
	engine.GET("/movies/:id", handler.GetMovieByID)
	engine.POST("/movies", handler.CreateMovie)
	engine.PUT("/movies/:id", handler.PutMovie)
	engine.POST("/auth/login", handler.LoginHandler)
	engine.POST("/auth/logout", handler.LogoutHandler)
	engine.POST("/auth/verify", handler.VerificationHandler)

	engine.Run(":8000")
}
