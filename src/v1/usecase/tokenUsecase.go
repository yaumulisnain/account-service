package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"account-service/src/core"
	"account-service/src/v1/model"
)

func jwtKey() []byte {
	return []byte(os.Getenv("JWT_SECRET_KEY"))
}

func CleanUpToken(token string) (string, error) {
	if hasBearer := strings.Index(token, "Bearer"); hasBearer < 0 || token == "" {
		return "", errors.New("need bearer authorization")
	}

	return strings.Split(token, "Bearer ")[1], nil
}

func createToken(user model.User, expired int) (string, error) {
	expirationTime := time.Now().Add(time.Duration(expired) * time.Hour)

	claims := &model.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    os.Getenv("JWT_OWNER"),
		},
		Data: user,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey())

	return tokenString, err
}

func CreateBundleToken(user model.User) (model.TokenResponse, error) {
	var (
		token, refreshToken string
		err                 error
		tokenBundle         model.TokenResponse
		redisValue          []byte
	)

	token, err = createToken(user, 1)

	if err != nil {
		return tokenBundle, err
	}

	refreshToken, err = createToken(user, 24*7)

	if err != nil {
		return tokenBundle, err
	}

	tokenBundle = model.TokenResponse{
		User:         user,
		Token:        token,
		RefreshToken: refreshToken,
	}

	redisValue, err = json.Marshal(tokenBundle)

	if err != nil {
		return tokenBundle, err
	}

	if err = core.RedisSet(user.UserName, string(redisValue)); err != nil {
		return tokenBundle, err
	}

	if err = core.RedisSetExpire(user.UserName, 3600*24*7); err != nil {
		return tokenBundle, err
	}

	return tokenBundle, nil
}

func GetTokenPayload(tokenStr string) (model.User, error) {
	var user model.User

	claims := &model.Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey(), nil
	})

	if err != nil {
		return user, err
	}

	claims, ok := token.Claims.(*model.Claims)

	if !ok {
		return user, errors.New("failed read claims")
	}

	user = claims.Data

	return user, nil
}

func VerifyToken(tokenStr string) (string, bool, error) {
	var (
		tokenResponse model.TokenResponse
		tokenCache    string
	)

	claims := &model.Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey(), nil
	})

	if err != nil {
		return "", false, err
	}

	claims, ok := token.Claims.(*model.Claims)

	if !ok {
		return "", false, errors.New("failed read claims")
	}

	tokenCache, err = core.RedisGet(claims.Data.UserName)

	if err != nil {
		return "", false, err
	}

	if err = json.Unmarshal([]byte(tokenCache), &tokenResponse); err != nil {
		return "", false, err
	}

	if tokenResponse.RefreshToken == tokenStr {
		return "refresh", true, nil
	} else if tokenResponse.Token == tokenStr {
		return "token", true, nil
	}

	return "", false, errors.New("token is not in whitelist, please re-login")
}
