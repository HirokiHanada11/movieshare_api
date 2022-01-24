package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func errorAndAbort(context *gin.Context, err error) {
	context.JSON(http.StatusBadRequest, gin.H{
		"message": fmt.Sprintf("Bad request error: %s", err.Error()),
	})
	context.Abort()
}