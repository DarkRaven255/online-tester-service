package domainmodel

import (
	"online-tests/delivery/commands"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Result struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	FinishedAt time.Time
	DeletedAt  gorm.DeletedAt `sql:"index"`
	FirstName  string
	LastName   string
	Email      string
	Result     float32
	TestID     uuid.UUID
}

func (Result) TableName() string {
	return "onlinetests.results"
}

func (result *Result) BeforeCreate(tx *gorm.DB) (err error) {
	nullUUID := uuid.UUID{}
	if result.ID == nullUUID {
		result.ID = uuid.New()
	}
	return
}

func NewResultModel(cmd *commands.StartTestCmd, id *uuid.UUID, testTime *uint) *Result {
	return &Result{
		FirstName:  cmd.Result.FirstName,
		LastName:   cmd.Result.LastName,
		Email:      cmd.Result.Email,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		FinishedAt: time.Now().Add(time.Duration(*testTime) * time.Minute),
		TestID:     *id,
	}
}
