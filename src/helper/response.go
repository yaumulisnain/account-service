package helper

import (
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
)

type ResponseMeta struct {
	TotalPage   int
	TotalRecord int
	Limit       int
	PageNumber  int
}

type ErrorResp struct {
	Error   string `json:"error"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func GetPagination(db *gorm.DB, paginationModel map[string][]string) (*gorm.DB, int, int, error) {

	var limit, page = 10, 1
	var errLimit, errPage error

	for i, v := range paginationModel {
		switch i {
		case "order":
			if len(v) == 1 {
				v = append([]string{"created_at"}, v...)
			}
			db = db.Order(strings.Join(v, " "))
		case "limit":
			limit, errLimit = strconv.Atoi(v[0])

			if val, ok := paginationModel["page"]; ok {
				page, errPage = strconv.Atoi(val[0])
				if errLimit != nil && errPage != nil {
					return db, limit, page, errPage
				}
				if limit == 0 {
					limit = 10
				}
			}
		default:
		}
	}
	db = db.Limit(limit)
	db = db.Offset((page - 1) * limit)

	return db, limit, page, nil
}
