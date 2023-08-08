package models

import "github.com/golang-jwt/jwt"

type GHCredentials struct {
	Code string `json:"code"`
}

type GHCredsReq struct {
	Code         string `json:"code"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type GHToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

type GHUser struct {
	Username string `json:"login"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type GoogleUser struct {
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	FirstName     string `json:"given_name"`
	LastName      string `json:"family_name"`
	jwt.StandardClaims
}

type GoogleToken struct {
	GoogleJWT string `json:"token"`
}
