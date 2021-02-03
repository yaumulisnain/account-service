package model

// import "github.com/dgrijalva/jwt-go"

// type ClaimsPayload struct {
// 	Key string `json:"key"`
// }

// type Claims struct {
// 	jwt.StandardClaims
// 	Data ClaimsPayload
// }

type TokenResponse struct {
	User         User   `json:"user"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
