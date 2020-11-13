package model

import "github.com/jinzhu/gorm"

type Answer struct {
	ModelEntity
	gorm.Model
	QuestionID uint   `json:"question_id"`
	Answer     string `json:"question"`
	Correct    bool   `json:"correct"`
	Required   bool   `json:"required"`
}

func (Answer) TableName() string {
	return "onlinetests.answers"
}
