package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Secret key used for signing and verifying tokens
var secretKey = []byte("No_Ra1nB0wT4bl3_c4n_H3lp!")

func main() {
	// Hard-coded Refresh Token and Expired Access Token
	refreshToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDY2MzEzNDIsImlzcyI6IkFQSSBTZXJ2aWNlcyBDZW50ZXIiLCJzdWIiOiJ1c2VyMTIzIn0.RBUCLBWDRY6M0U7fxfYL5SOd8Ec-ZHeYib_c5lu20Vw"
	expiredAccessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im1vQGdtYWlsLmNvbSIsImV4cCI6MTcwNjYyNDM4MywicGFzc3dvcmQiOiJwYXNzMTIzIiwidXNlcm5hbWUiOiJtbzEyMyJ9.D52udm4cgNSf5S10wlBFd_uiTmkSUi_CXbCpVQSfyKc"

	// Check if Refresh Token is valid
	validRefreshToken, err := isValidRefreshToken(refreshToken)
	if err != nil {
		fmt.Println("Error checking Refresh Token:", err)
		return
	}

	if validRefreshToken {
		// Decode the expired Access Token
		claims, err := decodeAccessToken(expiredAccessToken)
		if err != nil {
			fmt.Println("Error decoding Access Token:", err)
			return
		}

		// Remove 'exp' field from claims
		delete(claims, "exp")

		// Encode new Access Token
		newAccessToken, err := encodeAccessToken(claims)
		if err != nil {
			fmt.Println("Error encoding new Access Token:", err)
			return
		}

		// Print results
		fmt.Println("Refresh Token is valid.")
		fmt.Println("New Access Token:", newAccessToken)
		fmt.Println("Old Refresh Token:", refreshToken)
	} else {
		fmt.Println("Invalid Refresh Token.")
	}
}

// isValidRefreshToken checks if the Refresh Token is not expired and has a valid signature.
func isValidRefreshToken(refreshToken string) (bool, error) {
	// Parse the Refresh Token without validation
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return false, err
	}

	// Check if the token is not expired and has a valid signature
	return token.Valid, nil
}

// decodeAccessToken decodes the Access Token and returns the claims.
func decodeAccessToken(accessToken string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			// Check if the error is due to token expiration
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				fmt.Println("Access Token is expired. Generating a new one...")

				// Create a new map to store the claims
				newClaims := make(jwt.MapClaims)

				// Extract claims from the expired token
				if claims, ok := token.Claims.(jwt.MapClaims); ok {
					for key, value := range claims {
						// Exclude the 'exp' field
						if key != "exp" {
							newClaims[key] = value
						}
					}
				}

				// Return the new claims for a new token
				return newClaims, nil
			}
		}
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid Access Token")
	}

	return claims, nil
}

// encodeAccessToken encodes the claims into a new Access Token.
func encodeAccessToken(claims jwt.MapClaims) (string, error) {
	// Add an expiration time to the new token (e.g., 1 hour from now)
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
