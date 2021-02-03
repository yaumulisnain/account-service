package usecase

import (
	"math"

	h "account-service/src/helper"
	"account-service/src/v1/model"
	"account-service/src/v1/repository"
)

func GetMusicChart(paginationModel map[string][]string) ([]model.Music, h.ResponseMeta, error) {
	var (
		chartCount int
		music      []model.Music
		err        error
		Response   h.ResponseMeta
		limit      int
		page       int
	)

	chartCount, err = repository.FetchChartCount()

	if err != nil {
		return music, Response, err
	}

	music, limit, page, err = repository.FetchChartWithPaginate(paginationModel)

	if err != nil {
		return music, Response, err
	}

	Response = h.ResponseMeta{
		TotalPage:   int(math.Ceil(float64(chartCount) / float64(limit))),
		TotalRecord: chartCount,
		Limit:       limit,
		PageNumber:  page,
	}
	return music, Response, nil
}
