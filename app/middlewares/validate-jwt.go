package middlewares

import (
	"net/http"

	"github.com/Tsuryu/arreglapp-commons/app/jwt"
	"github.com/gin-gonic/gin"
)

// ValidateJwt : validates jwt
func ValidateJwt(context *gin.Context) {
	auth := context.GetHeader("Authorization")

	claims, err := jwt.ValidateJWT(auth)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		context.Abort()
		return
	}

	if context.Keys == nil {
		context.Keys = make(map[string]interface{})
	}

	context.Keys["claims"] = claims
}
