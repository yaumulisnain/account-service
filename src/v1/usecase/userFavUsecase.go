package usecase

import (
	"account-service/src/v1/model"
	"account-service/src/v1/repository"
)

func ListUserFavorite(token string) ([]model.Music, error) {
	var (
		music  []model.Music
		userId int32
		err    error
	)

	// TODO: Replace User ID with token User ID
	userId = 1

	music, err = repository.GetUserFavMusic(userId)

	if err != nil {
		return music, err
	}

	return music, nil
}
