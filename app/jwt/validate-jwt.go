package jwt

import (
	"errors"
	"os"
	"strings"

	"github.com/Tsuryu/arreglapp-commons/app/models"
	"github.com/dgrijalva/jwt-go"
)

// ValidateJWT : validates a JWT to allow the user to keep working with the app
func ValidateJWT(token string) (*models.Claim, error) {
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	myKey := []byte(jwtSecretKey)
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return &models.Claim{}, errors.New("Invalid token format")
	}

	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})

	if !tkn.Valid {
		return &models.Claim{}, errors.New("Invalid token")
	}

	return tkn.Claims.(*models.Claim), err
}
