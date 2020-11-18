package domainmodel

import (
	"online-tests/delivery/command/cmdmodel"

	"github.com/jinzhu/gorm"
)

type Answer struct {
	gorm.Model
	Answer     string `json:"answer"`
	Correct    bool   `json:"-"`
	QuestionID uint   `json:"question_id"`
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
