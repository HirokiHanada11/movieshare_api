package authentication

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
)

// app is a singleton instance of firebase app
var app *firebase.App

// creates firebase instance if it doen't exist
func getFirebaseConnection() (*firebase.App) {
	if app == nil {
		app, err := firebase.NewApp(context.Background(), nil)
		if err != nil {
			log.Printf("error initializing app: %v\n", err)
			return app
		}
		return app
	}
	return app
}