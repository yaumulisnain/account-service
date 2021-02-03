package repository

import (
	h "account-service/src/helper"
	"account-service/src/v1/model"
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
		music []model.Music
		err   error
		limit int
		page  int
	)
	db := getDB()

	db, limit, page, err = h.GetPagination(db, paginationModel)

	if err != nil {
		return music, limit, page, err
	}

	if err := db.Find(&music).Order("position ASC", true).Error; err != nil {
		return music, limit, page, err
	}

	return music, limit, page, nil
}
