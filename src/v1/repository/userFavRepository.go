package repository

import "account-service/src/v1/model"

func GetUserFavMusic(userId int32) ([]model.Music, error) {
	var music []model.Music

	db := getDB()

	// select music.* from user_fav
	// join music on music.id = user_fav.music_id
	// where user_fav.user_id = 1
	// order by music."position" asc;

	if err := db.Table("user_fav").Select(`music.*`).Joins(`Join music on music.id = user_fav.music_id`).Where(`user_fav.user_id = ?`, userId).Order(`music."position" ASC`, true).Find(&music).Error; err != nil {
		return music, err
	}

	return music, nil
}
