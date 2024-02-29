package util

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
)

// CutstomValidator :
type CustomValidator struct {
	Validator *validator.Validate
}

// Validate : Validate Data
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

// GenerateJWT generates a JWT token for the given user
func GenerateJWT(user repository.UserData) (string, error) {
	// Define the expiration time for the token
	expirationTime := time.Now().Add(24 * time.Hour) // Token expires in 24 hours

	// Define the claims
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		IssuedAt:  time.Now().Unix(),
		Subject:   fmt.Sprintf("%d", user.Id),
	}

	// Create the token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secret key
	secretKey := []byte("secret")
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateUserByToken(tokenString string) int {
	tokenString = extractTokenFromHeader(tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil || !token.Valid {
		return 0
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0
	}

	userID, ok := claims["sub"].(string)
	if !ok {
		return 0
	}

	id, err := strconv.Atoi(userID)
	if err != nil {
		return 0
	}

	return id
}

func extractTokenFromHeader(token string) string {
	authHeaderParts := strings.Split(token, " ")
	if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
		return ""
	}

	return authHeaderParts[1]
}
