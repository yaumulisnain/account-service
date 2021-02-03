package repository

import (
	"account-service/src/v1/model"
)

func CreateUser(user *model.User) error {
	db := getDB()

	if err := db.Create(user).Error; err != nil {
		return err
	}

	return nil
}
