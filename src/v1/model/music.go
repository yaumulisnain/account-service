package model

import "time"

type Music struct {
	ID        int32      `json:"id" gorm:"primary_key;AUTO_INCREMENT;unique_index"`
	Title     string     `json:"title" validate:"required" gorm:"not_null"`
	Artist    string     `json:"artist" validate:"required" gorm:"not_null"`
	Position  int        `json:"position" validate:"required" gorm:"not_null"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}

func (Music) TableName() string {
	return "music"
}
