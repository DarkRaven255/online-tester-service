package model

import (
	"github.com/jinzhu/gorm"
)

type Question struct {
	gorm.Model
	TestID   uint   `json:"test_id"`
	Question string `json:"question"`
}

func (Question) TableName() string {
	return "onlinetests.questions"
}
