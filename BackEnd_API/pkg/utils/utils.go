package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var secretKey = []byte("No_Ra1nB0wT4bl3_c4n_H3lp!")

func JWTGenerateToken(claims map[string]interface{}) (string, error) {
	// Append the expiration claim (4 minutes)
	claims["exp"] = time.Now().Add(time.Minute * 4).Unix()

	// Create a new token object, specifying signing method and the claims you want to include
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(claims))

	// Sign the token with a secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func JWTDecodeToken(tokenString string, secret string) (map[string]interface{}, bool, error) {
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

	refreshTokenString, err := refreshToken.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return refreshTokenString, nil
}
