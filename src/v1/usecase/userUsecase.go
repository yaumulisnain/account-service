package usecase

import (
	"errors"

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
		token model.TokenResponse
		user  model.User
		err   error
	)

	user, err = repository.FetchUserByUserName(userLogin.UserName)

	if err != nil {
		return token, err
	}

	if !h.ComparePasswords(user.Password, []byte(userLogin.Password)) {
		return token, errors.New("invalid password")
	}

	// TODO: Create JWT Token

	return token, nil
}
