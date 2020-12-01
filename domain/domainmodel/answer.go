package domainmodel

import (
	"online-tests/delivery/models/testmodel"
	"time"

	"gorm.io/gorm"
)

type Answer struct {
	ID         uint64 `gorm:"primary_key"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `sql:"index"`
	Answer     string
	Correct    bool
	QuestionID uint64
}

func (Answer) TableName() string {
	return "onlinetests.answers"
}

func newAnswer(a *testmodel.Answer, qID uint64) *Answer {
	return &Answer{
		ID:         a.ID,
		Answer:     a.Answer,
		Correct:    a.Correct,
		QuestionID: qID,
	}
}

func newAnswerArray(aArr *[]testmodel.Answer, qID uint64) *[]Answer {
	answers := []Answer{}
	for _, a := range *aArr {
		answers = append(answers, *newAnswer(&a, qID))
	}

	return &answers
}
