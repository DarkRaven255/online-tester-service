package domainmodel

import (
	"online-tests/delivery/command/cmdmodel"
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

func NewQuestion(q *cmdmodel.Question, tID uint64) *Question {
	return &Question{
		ID:       q.ID,
		Question: q.Question,
		Required: q.Required,
		Answers:  *NewAnswerArray(&q.Answer, q.ID),
		TestID:   tID,
	}
}

func NewQuestionsArray(qArr *[]cmdmodel.Question, tID uint64) *[]Question {
	questions := []Question{}
	for _, q := range *qArr {
		questions = append(questions, *NewQuestion(&q, tID))
	}

	return &questions
}
