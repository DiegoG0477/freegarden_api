package middlewares

import (
	"time"
	"github.com/golang-jwt/jwt/v4"
)


func GenerateJWT(clientID int64, email string) (string, error) {
	claims := CustomClaims{
		ClientID: clientID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)), // Expira en 72 horas
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // Fecha de emisión
			NotBefore: jwt.NewNumericDate(time.Now()),                     // No válido antes de
			Issuer:    "myapp",                                            // Emisor del token
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
