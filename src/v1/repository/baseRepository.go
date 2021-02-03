package repository

import (
	"github.com/jinzhu/gorm"

	"account-service/src/core"
)

func getDB() *gorm.DB {
	return core.GetDB()
}
