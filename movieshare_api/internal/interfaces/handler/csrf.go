package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
	"github.com/rikuhatano09/movieshare_api/pkg/domain/contract"
)

func CSRF(context *gin.Context) {
	// generate csrf token to be used for subsequent POST requests
	// should be called at the very first load of the website
	context.JSON(http.StatusOK, contract.CsrfResponse{
		CsrfToken: csrf.Token(context.Request),
	})
}