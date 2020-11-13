package model

import "github.com/jinzhu/gorm"

type User struct {
	ModelEntity
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Email    bool   `json:"email"`
}

func (User) TableName() string {
	return "onlinetests.users"
}
