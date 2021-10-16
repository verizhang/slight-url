package config

import "github.com/dgrijalva/jwt-go"

type CustomClaims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}
