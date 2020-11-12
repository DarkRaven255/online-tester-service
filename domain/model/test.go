package model

import (
	"github.com/jinzhu/gorm"
)

type Test struct {
	gorm.Model
	UserID         uint   `json:"user_id"`
	NumOfQuestions int    `json:"num_of_questions"`
	TestCode       string `json:"test_code"`
}

func (Test) TableName() string {
	return "onlinetests.tests"
}
