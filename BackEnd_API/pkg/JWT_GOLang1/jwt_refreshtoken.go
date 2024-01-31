package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenerateRefreshToken generates a refresh token with the user's information.
func GenerateRefreshToken(userID string) (string, error) {
	expirationTime := time.Now().Add(999 * 24 * time.Hour)

	// Create a new refresh token object with user-specific claims
	refreshTokenClaims := jwt.MapClaims{
		"sub": userID,                // Subject (user ID)
		"exp": expirationTime.Unix(), // Expiration time
		"iss": "API Services Center", // Issuer
	}

	// Create a new refresh token, specifying the signing method and claims
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	// Sign the refresh token with a secret key
	refreshTokenSecret := "No_Ra1nB0wT4bl3_c4n_H3lp!"
	refreshTokenString, err := refreshToken.SignedString([]byte(refreshTokenSecret))
	if err != nil {
		return "", err
	}

	return refreshTokenString, nil
}

func main() {
	// Example usage: generate a refresh token for a user with ID "user123"
	userID := "user123"
	refreshToken, err := GenerateRefreshToken(userID)
	if err != nil {
		fmt.Println("Error generating refresh token:", err)
		return
	}

	fmt.Println("Generated Refresh Token:", refreshToken)
}
