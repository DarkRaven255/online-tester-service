package domainmodel

import (
	"online-tests/delivery/command/cmdmodel"
	"time"
)

type Question struct {
	ID        uint64     `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-" sql:"index"`
	Question  string     `json:"question"`
	Required  bool       `json:"required"`
	TestID    uint64     `json:"-"`
	Answers   []Answer   `json:"answers" gorm:"foreignKey:QuestionID,constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
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
