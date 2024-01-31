package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// DecodeToken decodes a JWT and returns the claims in a dictionary form along with an expiration flag.
func DecodeToken(tokenString string, secret string) (map[string]interface{}, bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify the signing method and return the secret key
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, false, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, false, fmt.Errorf("invalid token")
	}

	// Check if the token is expired
	expired := claims.VerifyExpiresAt(time.Now().Unix(), false)

	// Remove the "exp" field from the claims
	delete(claims, "exp")

	return claims, expired, nil
}

func main() {
	// Example usage: decode a token
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im1vQGdtYWlsLmNvbSIsImV4cCI6MTcwNjYyNDM4MywicGFzc3dvcmQiOiJwYXNzMTIzIiwidXNlcm5hbWUiOiJtbzEyMyJ9.D52udm4cgNSf5S10wlBFd_uiTmkSUi_CXbCpVQSfyKc" // Replace with the actual encoded token
	secret := "No_Ra1nB0wT4bl3_c4n_H3lp!"                                                                                                                                                                          // Replace with the secret key used for signing the token

	decodedClaims, expired, err := DecodeToken(tokenString, secret)
	if err != nil {
		fmt.Println("Error decoding token:", err)
		return
	}

	fmt.Println("Decoded Claims:", decodedClaims)
	fmt.Println("Valid:", expired)
}
