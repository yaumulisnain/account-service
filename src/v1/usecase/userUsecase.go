package usecase

import (
	"encoding/json"
	"errors"

	"github.com/gomodule/redigo/redis"

	"account-service/src/core"
	h "account-service/src/helper"
	"account-service/src/v1/model"
	"account-service/src/v1/repository"
)

func Registration(user *model.User) error {
	user.Password = h.HashPassword([]byte(user.Password))

	if err := repository.CreateUser(user); err != nil {
		return err
	}

	return nil
}

func Login(userLogin model.UserLogin) (model.TokenResponse, error) {
	var (
		token      model.TokenResponse
		user       model.User
		tokenCache string
		err        error
	)

	user, err = repository.FetchUserByUserName(userLogin.UserName)

	if err != nil {
		return token, err
	}

	if !h.ComparePasswords(user.Password, []byte(userLogin.Password)) {
		return token, errors.New("invalid password")
	}

	tokenCache, err = core.RedisGet(user.UserName)

	if err == nil {
		if tokenCache != "" {
			return token, errors.New("user has been loged in, you must logout first to re-login")
		}
	} else {
		if err != redis.ErrNil {
			return token, err
		}
	}

	token, err = CreateBundleToken(user)

	if err != nil {
		return token, err
	}

	return token, nil
}

func RefreshToken(token string) (model.TokenResponse, error) {
	var (
		user          model.User
		tokenResponse model.TokenResponse
		tokenCache    string
		err           error
	)

	user, err = GetTokenPayload(token)

	if err != nil {
		return tokenResponse, err
	}

	tokenCache, err = core.RedisGet(user.UserName)

	if err != nil {
		return tokenResponse, err
	}

	if err = json.Unmarshal([]byte(tokenCache), &tokenResponse); err != nil {
		return tokenResponse, err
	}

	if tokenResponse.RefreshToken != token {
		return tokenResponse, errors.New("refresh token is not valid")
	}

	tokenResponse, err = CreateBundleToken(user)

	if err != nil {
		return tokenResponse, err
	}

	return tokenResponse, nil
}

func Logout(token string) error {
	var (
		user model.User
		err  error
	)

	user, err = GetTokenPayload(token)

	if err != nil {
		return err
	}

	if err = core.RedisDelete(user.UserName); err != nil {
		return err
	}

	return nil
}
