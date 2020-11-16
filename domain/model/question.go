package model

import (
	"online-tests/delivery/model"

	"github.com/jinzhu/gorm"
)

type Question struct {
	*gorm.Model
	Question string   `json:"question"`
	Required bool     `json:"required"`
	TestID   uint     `json:"test_id"`
	Answers  []Answer `json:"answers" gorm:"foreignKey:QuestionID"`
}

func (Question) TableName() string {
	return "onlinetests.questions"
}

func NewQuestion(q *model.Question) *Question {
	return &Question{
		Question: q.Question,
		Required: q.Required,
		Answers:  *NewAnswerArray(&q.Answer),
	}
}

func NewQuestionsArray(qArr *[]model.Question) *[]Question {
	questions := []Question{}
	for _, q := range *qArr {
		questions = append(questions, *NewQuestion(&q))
	}

	return &questions
}
