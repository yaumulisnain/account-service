package model

import "time"

type UserFav struct {
	ID        int32      `json:"id" gorm:"primary_key;AUTO_INCREMENT;unique_index"`
	UserID    int32      `json:"phone_number" validate:"required" gorm:"not_null"`
	MusicID   int32      `json:"name" validate:"required" gorm:"not_null"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}

func (UserFav) TableName() string {
	return "user_fav"
}
