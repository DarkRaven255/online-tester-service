package domainmodel

import (
	"online-tester-service/delivery/models/testmodel"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Question struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `sql:"index"`
	Question  string
	Required  bool
	TestID    uuid.UUID
	Answers   []Answer `gorm:"foreignKey:QuestionID"`
}

func (Question) TableName() string {
	return "online_tester_service.questions"
}

func (question *Question) BeforeCreate(tx *gorm.DB) (err error) {
	nullUUID := uuid.UUID{}
	if question.ID == nullUUID {
		question.ID = uuid.New()
	}
	return
}

func newQuestion(q *testmodel.Question, tID uuid.UUID) *Question {
	return &Question{
		ID:       q.ID,
		Question: q.Question,
		Required: q.Required,
		Answers:  *newAnswerArray(&q.Answers, q.ID),
		TestID:   tID,
	}
}

func newQuestionsArray(qArr *[]testmodel.Question, tID uuid.UUID) *[]Question {
	questions := []Question{}
	for _, q := range *qArr {
		questions = append(questions, *newQuestion(&q, tID))
	}

	return &questions
}
