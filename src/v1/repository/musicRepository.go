package repository

import (
	"account-service/src/core"
	h "account-service/src/helper"
	"account-service/src/v1/model"
	"encoding/json"
	"fmt"
)

func FetchChartCount() (int, error) {
	var (
		music []model.Music
		count int
	)

	db := getDB()

	if err := db.Find(&music).Count(&count).Error; err != nil {
		return count, err
	}

	return count, nil
}

func FetchChartWithPaginate(paginationModel map[string][]string) ([]model.Music, int, int, error) {
	var (
		music      []model.Music
		musicRedis string
		redisData  []byte
		err        error
		limit      int
		page       int
	)
	db := getDB()

	db, limit, page, err = h.GetPagination(db, paginationModel)

	if err != nil {
		return music, limit, page, err
	}

	redisKey := fmt.Sprintf("music_%d_%d", page, limit)

	musicRedis, err = core.RedisGet(redisKey)

	if err == nil {
		if err := json.Unmarshal([]byte(musicRedis), &music); err == nil {
			return music, limit, page, nil
		}
	}

	if err := db.Find(&music).Order("position ASC", true).Error; err != nil {
		return music, limit, page, err
	}

	redisData, err = json.Marshal(music)

	if err != nil {
		return music, limit, page, err
	}

	if err = core.RedisSet(redisKey, string(redisData)); err != nil {
		return music, limit, page, err
	}

	if err = core.RedisSetExpire(redisKey, 3600); err != nil {
		return music, limit, page, err
	}

	return music, limit, page, nil
}
