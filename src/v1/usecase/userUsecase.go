package usecase

import (
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
