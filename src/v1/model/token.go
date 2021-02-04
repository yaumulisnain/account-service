package model

import jwt "github.com/dgrijalva/jwt-go"

type Claims struct {
	jwt.StandardClaims
	Data User
}

type TokenResponse struct {
	User         User   `json:"user"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
