package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func generateToken(claims map[string]interface{}) (string, error) {
	// Append the expiration claim (4 minutes)
	claims["exp"] = time.Now().Add(time.Minute * 4).Unix()

	// Create a new token object, specifying signing method and the claims you want to include
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(claims))

	// Sign the token with a secret key
	secret := "No_Ra1nB0wT4bl3_c4n_H3lp!"
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func main() {
	// Get parameters from the user
	fmt.Print("Enter username: ")
	var username string
	fmt.Scanln(&username)

	fmt.Print("Email: ")
	var email string
	fmt.Scanln(&email)

	fmt.Print("Password: ")
	var password string
	fmt.Scanln(&password)

	// Example usage with user-provided claims
	userClaims := map[string]interface{}{
		"username": username,
		"email":    email,
		"password": password,
		// You don't need to provide "exp" here; it will be added in generateToken function
	}

	token, err := generateToken(userClaims)
	if err != nil {
		fmt.Println("Error generating token:", err)
		return
	}

	fmt.Println("Generated Token:", token)
}
