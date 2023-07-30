package models

import "github.com/golang-jwt/jwt"

type Claims struct {
	Identifier string `json:"identifier"`
	Type       string `json:"type"`
	jwt.StandardClaims
}
