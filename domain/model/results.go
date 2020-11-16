package model

import (
	"github.com/jinzhu/gorm"
)

type Results struct {
	gorm.Model
	TestID uint    `json:"test_id"`
	UserID uint    `json:"user_id"`
	Result float32 `json:"result"`
}

func (Results) TableName() string {
	return "onlinetests.results"
}
