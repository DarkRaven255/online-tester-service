package model

import (
	"github.com/jinzhu/gorm"
)

type Question struct {
	ModelEntity
	gorm.Model
	TestID   uint   `json:"test_id"`
	Question string `json:"question"`
}

func (Question) TableName() string {
	return "onlinetests.questions"
}
