package services

import (
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// ValidateJwtToken verify token authenticity and signature. Returns true/false & error
func ValidateJwtToken(tokenString string) (bool, error) {
	tokenString = strings.Trim(tokenString, " ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["HS256"])
		}
		return []byte(os.Getenv("JWT_TOKEN_SECRET")), nil
	})
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, nil
	}

	return false, err
}
