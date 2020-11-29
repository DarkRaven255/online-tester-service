package domainmodel

import (
	"online-tests/delivery/models/testmodel"
	"time"

	"gorm.io/gorm"
)

type Question struct {
	ID        uint64         `json:"id" gorm:"primary_key"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" sql:"index"`
	Question  string         `json:"question"`
	Required  bool           `json:"required"`
	TestID    uint64         `json:"-"`
	Answers   []Answer       `json:"answers" gorm:"foreignKey:QuestionID"`
}

func (Question) TableName() string {
	return "onlinetests.questions"
}

func newQuestion(q *testmodel.Question, tID uint64) *Question {
	return &Question{
		ID:       q.ID,
		Question: q.Question,
		Required: q.Required,
		Answers:  *newAnswerArray(&q.Answers, q.ID),
		TestID:   tID,
	}
}

func newQuestionsArray(qArr *[]testmodel.Question, tID uint64) *[]Question {
	questions := []Question{}
	for _, q := range *qArr {
		questions = append(questions, *newQuestion(&q, tID))
	}

	return &questions
}
