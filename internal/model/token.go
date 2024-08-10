package model

import "github.com/dgrijalva/jwt-go"

// AdminClaims represents the claims of the admin JWT token.
type AdminClaims struct {
	UserID   uint
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}
