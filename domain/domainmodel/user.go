package domainmodel

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Email    bool   `json:"email"`
}

func (User) TableName() string {
	return "onlinetests.users"
}
