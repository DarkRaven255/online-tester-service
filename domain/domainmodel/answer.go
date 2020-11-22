package domainmodel

import (
	"online-tests/delivery/command/cmdmodel"
	"time"
)

type Answer struct {
	ID         uint64     `json:"id" gorm:"primary_key"`
	CreatedAt  time.Time  `json:"-"`
	UpdatedAt  time.Time  `json:"updatedAt"`
	DeletedAt  *time.Time `json:"-" sql:"index"`
	Answer     string     `json:"answer"`
	Correct    bool       `json:"correct"`
	QuestionID uint64     `json:"-"`
}

func (Answer) TableName() string {
	return "onlinetests.answers"
}

func NewAnswer(a *cmdmodel.Answer) *Answer {
	return &Answer{
		Answer:  a.Answer,
		Correct: a.Correct,
	}
}

func NewAnswerArray(aArr *[]cmdmodel.Answer) *[]Answer {
	answers := []Answer{}
	for _, a := range *aArr {
		answers = append(answers, *NewAnswer(&a))
	}

	return &answers
}
