package domainmodel

import (
	"online-tests/delivery/models/testmodel"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Answer struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `sql:"index"`
	Answer     string
	Correct    bool
	QuestionID uuid.UUID
}

func (Answer) TableName() string {
	return "onlinetests.answers"
}

func (answer *Answer) BeforeCreate(tx *gorm.DB) (err error) {
	nullUUID := uuid.UUID{}
	if answer.ID == nullUUID {
		answer.ID = uuid.New()
	}
	return
}

func newAnswer(a *testmodel.Answer, qID uuid.UUID) *Answer {
	return &Answer{
		ID:         a.ID,
		Answer:     a.Answer,
		Correct:    a.Correct,
		QuestionID: qID,
	}
}

func newAnswerArray(aArr *[]testmodel.Answer, qID uuid.UUID) *[]Answer {
	answers := []Answer{}
	for _, a := range *aArr {
		answers = append(answers, *newAnswer(&a, qID))
	}

	return &answers
}
