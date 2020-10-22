package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Healthcheck : checks whether the app is running
func Healthcheck(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "App up and running"})
}
