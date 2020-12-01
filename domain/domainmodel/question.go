package domainmodel

import (
	"online-tests/delivery/models/testmodel"
	"time"

	"gorm.io/gorm"
)

type Question struct {
	ID        uint64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `sql:"index"`
	Question  string
	Required  bool
	TestID    uint64
	Answers   []Answer `gorm:"foreignKey:QuestionID"`
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
