package domainmodel

import (
	"online-tests/delivery/command/cmdmodel"

	"github.com/jinzhu/gorm"
)

type Question struct {
	gorm.Model
	Question string   `json:"question"`
	Required bool     `json:"-"`
	TestID   uint     `json:"-"`
	Answers  []Answer `json:"answers" gorm:"foreignKey:QuestionID"`
}

func (Question) TableName() string {
	return "onlinetests.questions"
}

func NewQuestion(q *cmdmodel.Question) *Question {
	return &Question{
		Question: q.Question,
		Required: q.Required,
		Answers:  *NewAnswerArray(&q.Answer),
	}
}

func NewQuestionsArray(qArr *[]cmdmodel.Question) *[]Question {
	questions := []Question{}
	for _, q := range *qArr {
		questions = append(questions, *NewQuestion(&q))
	}

	return &questions
}
