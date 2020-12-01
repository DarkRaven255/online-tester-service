package domainmodel

import (
	"online-tests/delivery/commands"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Result struct {
	ID         uint64 `gorm:"primary_key"`
	ResultUUID string `qorm:"unique"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `sql:"index"`
	FirstName  string
	LastName   string
	Email      string
	Result     float32
	TestID     uint64
}

func (Result) TableName() string {
	return "onlinetests.results"
}

func NewResultModel(cmd *commands.StartTestCmd, id uint64) *Result {
	resultUUID := uuid.New().String()
	return &Result{
		ResultUUID: resultUUID,
		FirstName:  cmd.Result.FirstName,
		LastName:   cmd.Result.LastName,
		Email:      cmd.Result.Email,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		TestID:     id,
	}
}
