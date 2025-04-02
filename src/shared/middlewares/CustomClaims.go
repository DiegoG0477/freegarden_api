package middlewares

import "github.com/golang-jwt/jwt/v4"

type CustomClaims struct {
	ClientID int64  `json:"client_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}