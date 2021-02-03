package model

import "time"

type User struct {
	ID        int32      `json:"id" gorm:"primary_key;AUTO_INCREMENT;unique_index"`
	UserName  string     `json:"user_name" validate:"min=3,max=50,required" gorm:"not_null;unique_index"`
	FullName  string     `json:"full_name" validate:"required" gorm:"not_null"`
	Password  string     `json:"password" validate:"min=5,required" gorm:"not_null"`
	Gender    string     `json:"gender" gorm:"default:null"`
	Hobby     string     `json:"hobby" gorm:"default:null"`
	Address   string     `json:"address" gorm:"default:null"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}

func (User) TableName() string {
	return "user"
}
