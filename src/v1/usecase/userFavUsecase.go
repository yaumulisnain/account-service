package usecase

import (
	"account-service/src/v1/model"
	"account-service/src/v1/repository"
)

func ListUserFavorite(token string) ([]model.Music, error) {
	var (
		music  []model.Music
		user   model.User
		userId int32
		err    error
	)

	// TODO: Replace User ID with token User ID
	user, err = GetTokenPayload(token)

	if err != nil {
		return music, err
	}

	userId = user.ID

	music, err = repository.GetUserFavMusic(userId)

	if err != nil {
		return music, err
	}

	return music, nil
}
