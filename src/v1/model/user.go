package model

import "time"

type User struct {
	ID        int32      `json:"id" gorm:"primary_key;AUTO_INCREMENT;unique_index"`
	UserName  string     `json:"user_name" validate:"min=3,max=50,alphanum,required" gorm:"not_null;unique_index"`
	FullName  string     `json:"full_name" validate:"min=3,required" gorm:"not_null"`
	Password  string     `json:"password" validate:"min=5,required" gorm:"not_null"`
	Gender    string     `json:"gender,omitempty" validate:"eq=M|eq=F|eq=" gorm:"default:null"`
	Hobby     string     `json:"hobby,omitempty" gorm:"default:null"`
	Address   string     `json:"address,omitempty" gorm:"default:null"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}

func (User) TableName() string {
	return "user"
}

type UserLogin struct {
	UserName string `json:"user_name" validate:"min=3,max=50,required"`
	Password string `json:"password" validate:"min=5,required"`
}
