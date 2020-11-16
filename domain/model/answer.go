package model

import (
	"online-tests/delivery/model"

	"github.com/jinzhu/gorm"
)

type Answer struct {
	*gorm.Model
	Answer     string `json:"answer"`
	Correct    bool   `json:"correct"`
	QuestionID uint   `json:"question_id"`
}

func (Answer) TableName() string {
	return "onlinetests.answers"
}

func NewAnswer(a *model.Answer) *Answer {
	return &Answer{
		Answer:  a.Answer,
		Correct: a.Correct,
	}
}

func NewAnswerArray(aArr *[]model.Answer) *[]Answer {
	answers := []Answer{}
	for _, a := range *aArr {
		answers = append(answers, *NewAnswer(&a))
	}

	return &answers
}
