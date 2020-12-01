package domainmodel

import (
	"online-tests/delivery/commands"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Result struct {
	ID         uint64         `json:"id" gorm:"primary_key"`
	ResultUUID string         `json:"resultUUID" qorm:"unique"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `json:"-" sql:"index"`
	FirstName  string         `json:"firstName"`
	LastName   string         `json:"lastName"`
	Email      string         `json:"email"`
	Result     float32        `json:"result"`
	TestID     uint64         `json:"-"`
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
