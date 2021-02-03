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

func FetchUserByUserName(username string) (model.User, error) {
	var user model.User

	db := getDB()

	if err := db.Where("user_name = ?", username).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
